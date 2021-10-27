package core

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	//"github.com/cetic/kubeedge-controller/internal/api"
	api "github.com/cetic/kubeedge-controller/internal/api"
	"github.com/cetic/kubeedge-controller/internal/config"
	"github.com/emicklei/go-restful"
	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/rest"

	//"time"
	devices "github.com/kubeedge/kubeedge/cloud/pkg/apis/devices/v1alpha2"
)

type Controller struct {
	CRDClient *rest.RESTClient
	Devices   map[string]*Device
	Namespace string
}

func (c *Controller) Init() {
	conf := config.Parse()
	log.Debugf("kubeconfig File: %s", conf.KubeConfig)
	log.Debugf("kubeconfig File: %s", conf.MasterAdress)
	c.Devices = make(map[string]*Device)
	c.CRDClient, _ = NewCRDClient(conf.MasterAdress, conf.KubeConfig)
	c.Namespace = conf.Namespace
	log.Debugf("%+v", c)
	for _, device := range conf.Devices {
		c.Devices[device] = NewDevice()
		c.Devices[device].InitDevice(device, c.Namespace, c.CRDClient)
		log.Debugf("%+v", c.Devices)
	}
}

func (c Controller) GetDevicebyId(req *restful.Request, rsp *restful.Response) {
	//ctx := context.Background()
	deviceid := req.PathParameter("deviceid")
	raw, _ := c.CRDClient.Get().Namespace(c.Namespace).Resource(ResourceTypeDevices).Name(deviceid).DoRaw(context.TODO())
	log.Debugf("raw: %s", raw)
	result := devices.Device{}
	_ = json.Unmarshal(raw, &result)
	twins := make(map[string]interface{})
	for _, twin := range result.Status.Twins {
		twins[twin.PropertyName] = twin.Reported.Value
	}
	rsp.WriteEntity(twins)
}

func (c Controller) UpdateDevicebyId(req *restful.Request, rsp *restful.Response) {
	input := new(api.Update)
	deviceid := req.PathParameter("deviceid")
	if err := req.ReadEntity(&input); err != nil {
		log.Debugf(err.Error())
		rsp.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Debugf("Req : %s", input)
	if c.Devices[deviceid].FSM.Current() == "run" {
		c.Devices[deviceid].AddDesiredJob("Stop")
		c.Devices[deviceid].PatchStatus()
		for c.Devices[deviceid].FSM.Current() != "ready" {
			time.Sleep(500 * time.Millisecond)
		}
	}
	filename := input.Filename
	url := input.URL
	log.Debugf("filename : %s , url : %s", filename, url)
	go c.Devices[deviceid].Launch(filename, url)
}

func (c Controller) StopDevicebyId(req *restful.Request, rsp *restful.Response) {
	deviceid := req.PathParameter("deviceid")
	if c.Devices[deviceid].FSM.Current() == "run" {
		c.Devices[deviceid].AddDesiredJob("Stop")
		c.Devices[deviceid].PatchStatus()
		for c.Devices[deviceid].FSM.Current() != "ready" {
			time.Sleep(500 * time.Millisecond)
		}
	}
	rsp.WriteEntity("Stop Request Sended")
}
