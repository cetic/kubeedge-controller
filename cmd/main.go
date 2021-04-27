package main

import (
	//"flag"
	//"fmt"
	//log "github.com/sirupsen/logrus"
	//"io/ioutil"
	//"net/http"
	//"fmt"
	"github.com/cetic/kubeedge-controller/internal/core"
	log "github.com/sirupsen/logrus"
	//"github.com/cetic/kubeedge-controller/internal/webserver"
	//"os"
	////"reflect"
	//"time"
)

//func Facetrack(w http.ResponseWriter, r *http.Request) {
//  log.Println(d.GetStatus())
//  if d.FSM.Current() == "run" {
//    d.AddDesiredJob("Stop")
//    d.PatchStatus()
//    for d.FSM.Current() != "ready" {
//      time.Sleep(500*time.Millisecond)
//    }
//  }
//
//  filename := "/app/dpu_face_tracking.py"
//  url := ""
//  go d.Launch(filename,url)
//}
//
//func Passthrough(w http.ResponseWriter, r *http.Request) {
//  if d.FSM.Current() == "run" {
//    d.AddDesiredJob("Stop")
//    d.PatchStatus()
//    for d.FSM.Current() != "ready" {
//      time.Sleep(500*time.Millisecond)
//    }
//  }
//
//  filename := "/app/passthrough.py"
//  url := ""
//  go d.Launch(filename,url)
//}
//
//
//func Stop(w http.ResponseWriter, r *http.Request) {
//  if d.FSM.Current() == "run" {
//    d.AddDesiredJob("Stop")
//    d.PatchStatus()
//    for d.FSM.Current() != "ready" {
//      time.Sleep(500*time.Millisecond)
//    }
//  }
//}

var d = core.Device{}

func main() {
	log.Info(d)
	//Config := dto.Config{}
	//configFile := flag.String("c", "config/config.yaml", "config file")
	//flag.Parse()
	//args := flag.Args()
	//myself := os.Args[0]
	//if len(args) != 0 {
	//  fmt.Printf("Wrong number of argument : %s [-c configfile] \n", myself)
	//  os.Exit(1)
	//}
	//yamlFile, err := ioutil.ReadFile(*configFile)
	//if err != nil {
	//  log.Errorf("yamlFile.Get err   #%v ", err)
	//}
	//err = yaml.Unmarshal(yamlFile, &Config)
	//if err != nil {
	//  log.Errorf("Unmarshall err   #%v ", err)
	//}
	//crdClient, _ := core.NewCRDClient(os.Args[1],os.Args[2])
	//d.InitDevice(os.Args[3],"default",crdClient)
	//s := new(web.Site)
	//s.Init()
	//s.AddPage("Home","gotpl/welcome.gohtml","/passthrough","home", Passthrough)
	//s.AddPage("Home","gotpl/welcome.gohtml","/facetrack","home", Facetrack)
	//s.AddPage("Home","gotpl/welcome.gohtml","/stop","home", Stop)
	//http.Handle("/", s.Mux)
	//http.ListenAndServe(":9090", nil)
}
