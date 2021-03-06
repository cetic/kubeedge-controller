package core

import (
	"context"
	"encoding/json"
	"time"

	ke "github.com/cetic/kubeedge-controller/internal/kubeedge"
	"github.com/looplab/fsm"
	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/rest"
)

type Device struct {
	Status    ke.DeviceStatus `json:"status"`
	DeviceID  string          `json:"-"`
	Namespace string          `json:"-"`
	crdClient *rest.RESTClient
	FSM       *fsm.FSM `json:"-"`
	Filename  *string
	Url       *string
}

func NewDevice() *Device {
	var out Device
	return &out
}

func (s *Device) InitDevice(id, ns string, crdClient *rest.RESTClient) error {
	metadata := map[string]string{"timestamp": GetTimeStamp(), "type": "string"}
	twins := []ke.Twin{
		{PropertyName: "job",
			Desired:  ke.TwinProperty{Value: "unknown", Metadata: metadata},
			Reported: ke.TwinProperty{Value: "unknown", Metadata: metadata},
		},
		{PropertyName: "arg",
			Desired:  ke.TwinProperty{Value: "unknown", Metadata: metadata},
			Reported: ke.TwinProperty{Value: "unknown", Metadata: metadata},
		},
		{PropertyName: "status",
			Desired: ke.TwinProperty{Value: "unknown", Metadata: metadata},
		},
	}
	s.Status = ke.DeviceStatus{}
	s.Status.Twins = twins
	s.DeviceID = id
	s.Namespace = ns
	s.crdClient = crdClient
	f := "defaultFilename"
	s.Filename = &f
	u := "defaultURL"
	s.Url = &u
	s.SyncStatus()
	if s.GetStatus() != "Waiting" {
		log.Debug("Device is not in Waiting Mode...")
		s.AddDesiredJob("Wait")
		s.AddDesiredArg("init")
		_, err := s.PatchStatus()
		if err != nil {
			return err
		}
		log.Info("Waiting request sended")
		for s.GetStatus() != "Waiting" {
			s.SyncStatus()
			if err != nil {
				return err
			}
			time.Sleep(500 * time.Millisecond)
		}
	}
	log.Info("Device Connected and Ready")
	s.FSM = fsm.NewFSM(
		"ready",
		fsm.Events{
			{Name: "LaunchTask", Src: []string{"ready", "download"}, Dst: "run"},
			{Name: "TaskCompleted", Src: []string{"run"}, Dst: "done"},
			{Name: "FileNotFound", Src: []string{"run"}, Dst: "download"},
			{Name: "DownloadError", Src: []string{"download"}, Dst: "error"},
			{Name: "TaskError", Src: []string{"run"}, Dst: "error"},
			{Name: "Waiting", Src: []string{"ready", "done"}, Dst: "ready"},
		},
		fsm.Callbacks{
			"LaunchTask": func(e *fsm.Event) {
				log.Infof("Launch app %s request", *s.Filename)
				s.AddDesiredJob("Launch")
				s.AddDesiredArg(*s.Filename)
				s.PatchStatus()
			},
			"TaskCompleted": func(e *fsm.Event) {
				s.AddDesiredJob("Wait")
				s.AddDesiredArg("Finished")
				s.PatchStatus()
			},
			"FileNotFound": func(e *fsm.Event) {
				log.Printf("File %s not found on Device %s", s.Filename, s.DeviceID)
				log.Printf("Download from ", s.Url)
				s.AddDesiredJob("Download")
				s.AddDesiredArg(*s.Url)
				s.PatchStatus()
			},
		},
	)
	go func(d *Device) {
		log.Debugf("Current Status: %s", d.FSM.Current())
		time.Sleep(time.Second)
	}(s)
	return nil
}

func (s *Device) Launch(filename, url string) {
	log.Debugf("filename : %s , url : %s", filename, url)
	*s.Filename = filename
	*s.Url = url
	log.Debugf("filename : %s , url : %s", s.Filename, s.Url)
	s.FSM.Event("LaunchTask")

	for s.FSM.Current() != "ready" {
		s.SyncStatus()
		s.FSM.Event(s.GetStatus())
	}
}

func (s *Device) AddDesiredJob(job string) {
	for id, property := range s.Status.Twins {
		if property.PropertyName == "job" {
			s.Status.Twins[id].Desired.Value = job
		}
	}
}

func (s *Device) AddDesiredArg(arg string) {
	for id, property := range s.Status.Twins {
		if property.PropertyName == "arg" {
			s.Status.Twins[id].Desired.Value = arg
		}
	}
}

func (s *Device) PatchStatus() ([]byte, error) {
	//ctx := context.Background()
	body, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return s.crdClient.Patch(MergePatchType).Namespace(s.Namespace).Resource(ResourceTypeDevices).Name(s.DeviceID).Body(body).DoRaw(context.TODO())
}

func (s *Device) SyncStatus() error {
	//ctx := context.Background()
	raw, err := s.crdClient.Get().Namespace(s.Namespace).Resource(ResourceTypeDevices).Name(s.DeviceID).DoRaw(context.TODO())
	_ = json.Unmarshal(raw, &s)
	return err
}

func (s *Device) GetStatus() string {
	for id, property := range s.Status.Twins {
		if property.PropertyName == "status" {
			return s.Status.Twins[id].Reported.Value
		}
	}
	return "unknown"
}
