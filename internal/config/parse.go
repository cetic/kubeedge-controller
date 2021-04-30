package config

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func Parse() Config {
	conf := Config{}
	configFile := flag.String("c", "../configs/config.yaml", "config file")
	flag.Parse()
	args := flag.Args()
	myself := os.Args[0]
	if len(args) != 0 {
		log.Errorf("Wrong number of argument : %s [-c configfile] \n", myself)
		os.Exit(1)
	}
	yamlFile, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Errorf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Errorf("Unmarshall err   #%v ", err)
	}
	return conf
}
