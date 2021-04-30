//package main
//
//import (
//	"context"
//	"time"
//
//	//"flag"
//	//"fmt"
//	//log "github.com/sirupsen/logrus"
//	//"io/ioutil"
//	//"net/http"
//	//"fmt"
//	"github.com/cetic/kubeedge-controller/internal/config"
//	"github.com/cetic/kubeedge-controller/internal/core"
//	log "github.com/sirupsen/logrus"
//	//"github.com/cetic/kubeedge-controller/internal/webserver"
//	//"os"
//	////"reflect"
//	//"time"
//)
//
////func Facetrack(w http.ResponseWriter, r *http.Request) {
////  log.Println(d.GetStatus())
////  if d.FSM.Current() == "run" {
////    d.AddDesiredJob("Stop")
////    d.PatchStatus()
////    for d.FSM.Current() != "ready" {
////      time.Sleep(500*time.Millisecond)
////    }
////  }
////
////  filename := "/app/dpu_face_tracking.py"
////  url := ""
////  go d.Launch(filename,url)
////}
////
//func Passthrough() {
//	if d.FSM.Current() == "run" {
//		d.AddDesiredJob("Stop")
//		d.PatchStatus()
//		for d.FSM.Current() != "ready" {
//			time.Sleep(500 * time.Millisecond)
//		}
//	}
//
//	filename := "/Users/tse/CETIC/viaduct/POCEdge/kubeedge-edge-worker/mapper/me-loop.py"
//	url := ""
//	go d.Launch(filename, url)
//}
//
//func Passthrough2() {
//	if d.FSM.Current() == "run" {
//		d.AddDesiredJob("Stop")
//		d.PatchStatus()
//		for d.FSM.Current() != "ready" {
//			time.Sleep(500 * time.Millisecond)
//		}
//	}
//
//	filename := "/Users/tse/CETIC/viaduct/POCEdge/kubeedge-edge-worker/mapper/hello-loop.py"
//	url := ""
//	go d.Launch(filename, url)
//}
//
////
////
////func Stop(w http.ResponseWriter, r *http.Request) {
////  if d.FSM.Current() == "run" {
////    d.AddDesiredJob("Stop")
////    d.PatchStatus()
////    for d.FSM.Current() != "ready" {
////      time.Sleep(500*time.Millisecond)
////    }
////  }
////}
//
//
//
//
//
//
//var d = core.Device{}
//
//func main() {
//	log.SetLevel(log.DebugLevel)
//	conf := config.Parse()
//	log.Debugf("kubeconfig File: %s", conf.KubeConfig)
//	log.Debugf("kubeconfig File: %s", conf.MasterAdress)
//	crdClient, _ := core.NewCRDClient(conf.MasterAdress, conf.KubeConfig)
//	d.InitDevice(conf.Device, "default", crdClient)
//	ctx := context.Background()
//	for {
//		raw, _ := crdClient.Get().Namespace("default").Resource(core.ResourceTypeDevices).Name(conf.Device).DoRaw(ctx)
//		log.Debugf("result: %s", raw)
//		time.Sleep(1 * time.Second)
//	}
//
//	//time.Sleep(10 * time.Second)
//	//Passthrough()
//	//time.Sleep(10 * time.Second)
//	//Passthrough2()
//	//time.Sleep(10 * time.Second)
//	//s := new(web.Site)
//	//s.Init()
//	//s.AddPage("Home","gotpl/welcome.gohtml","/passthrough","home", Passthrough)
//	//s.AddPage("Home","gotpl/welcome.gohtml","/facetrack","home", Facetrack)
//	//s.AddPage("Home","gotpl/welcome.gohtml","/stop","home", Stop)
//	//http.Handle("/", s.Mux)
//	//http.ListenAndServe(":9090", nil)
//}
