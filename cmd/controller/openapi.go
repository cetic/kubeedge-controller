package main

import (
	"context"
	"github.com/cetic/kubeedge-controller/internal/api"
	"github.com/cetic/kubeedge-controller/internal/config"
	"github.com/cetic/kubeedge-controller/internal/core"
	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"
	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/rest"
	"net/http"
	"time"
)

type Ressource struct {
	BasePath string
	CRDClient *rest.RESTClient
	Device core.Device
}

func (r Ressource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path(r.BasePath).
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML) // you can specify this per route as well

	tags := []string{"Hello"}
	ws.Route(ws.GET("/hello").
		To(r.sayHello).
		Doc("Just Say Hello").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	tags = []string{"Controller"}
	ws.Route(ws.GET("/device").
		To(r.get).
		Doc("Get").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.POST("/device").
		To(r.post).
		Doc("Post").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(api.Post{}))
	return ws
}

func (r Ressource) sayHello(req *restful.Request, rsp *restful.Response) {
	rsp.WriteEntity("Hello World")
}

func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "ServiceName",
			Description: "Description of the service",
			Contact: &spec.ContactInfo{
				Name:  "john",
				Email: "john@doe.rp",
				URL:   "http://johndoe.org",
			},
			License: &spec.License{
				Name: "MIT",
				URL:  "http://mit.org",
			},
			Version: "1.0.0",
		},
	}
}


func (r *Ressource) init() {
	conf := config.Parse()
	log.Debugf("kubeconfig File: %s", conf.KubeConfig)
	log.Debugf("kubeconfig File: %s", conf.MasterAdress)
	r.CRDClient, _ = core.NewCRDClient(conf.MasterAdress, conf.KubeConfig)
	r.Device.InitDevice(conf.Device, "default", r.CRDClient)
}

func (r Ressource) get(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()
	raw, _ := r.CRDClient.Get().Namespace("default").Resource(core.ResourceTypeDevices).Name(r.Device.DeviceID).DoRaw(ctx)
	log.Debugf("raw: %s",raw)
	rsp.WriteEntity(string(raw))
}

func (r Ressource) post(req *restful.Request, rsp *restful.Response) {
	input := new(api.Post)
	if err := req.ReadEntity(&input); err != nil {
		log.Debugf(err.Error())
		rsp.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Debugf("Req : %s",input)
	if r.Device.FSM.Current() == "run" {
		r.Device.AddDesiredJob("Stop")
		r.Device.PatchStatus()
		for r.Device.FSM.Current() != "ready" {
			time.Sleep(500 * time.Millisecond)
		}
	}
	filename := input.Filename
	//filename := "/Users/tse/CETIC/viaduct/POCEdge/kubeedge-edge-worker/mapper/hello-loop.py"
	url := input.URL
	log.Debugf("filename : %s , url : %s",filename,url)
	// *r.Device.Filename = filename
	//r.Device.Url = url
	go r.Device.Launch(filename, url)
}


func main() {
	log.SetLevel(log.DebugLevel)
	wc := restful.NewContainer()
	u := Ressource{BasePath: "/api/v1"}
	u.init()
	wc.Add(u.WebService())

	// OpenAPI output documentation
	config := restfulspec.Config{
		WebServices: wc.RegisteredWebServices(),
		APIPath:     "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}
	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(config))
	wc.Add(restfulspec.NewOpenAPIService(config))


	// Optionally, you may need to enable CORS for the UI to work.
	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		CookiesAllowed: false,
		Container:      restful.DefaultContainer}
	wc.Filter(cors.Filter)

	log.Fatal(http.ListenAndServe(":8080", wc))
}

