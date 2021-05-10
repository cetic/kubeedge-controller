package main

import (
	"github.com/cetic/kubeedge-controller/internal/api"
	"github.com/cetic/kubeedge-controller/internal/core"
	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func sayHello(req *restful.Request, rsp *restful.Response) {
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

func main() {
	log.SetLevel(log.DebugLevel)
	wc := restful.NewContainer()
	const BasePath = "/api/v1"
	u := core.Controller{}
	u.Init()
	ws := new(restful.WebService)
	ws.Path(BasePath).
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML) // you can specify this per route as well

	tags := []string{"Hello"}
	ws.Route(ws.GET("/hello").
		To(sayHello).
		Doc("Just Say Hello").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	tags = []string{"Controller"}
	ws.Route(ws.GET("/device/{deviceid}").
		To(u.GetDevicebyId).
		Doc("Get Device Status by ID").
		Param(ws.PathParameter("deviceid", "Device ID").Required(true).DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.POST("/device/{deviceid}").
		To(u.UpdateDevicebyId).
		Doc("Get Device Status by ID").
		Param(ws.PathParameter("deviceid", "Device ID").Required(true).DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(api.Update{}).
		Returns(200, "OK", ""))

	ws.Route(ws.GET("/stop/{deviceid}").
		To(u.StopDevicebyId).
		Doc("Stop Device with ID").
		Param(ws.PathParameter("deviceid", "Device ID").Required(true).DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags))
	wc.Add(ws)

	// OpenAPI output documentation
	config := restfulspec.Config{
		WebServices:                   wc.RegisteredWebServices(),
		APIPath:                       "/apidocs.json",
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

	log.Fatal(http.ListenAndServe(":8090", wc))
}
