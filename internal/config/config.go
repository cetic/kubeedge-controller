package config

type Config struct {
	MasterAdress string   `yaml:"master"`
	KubeConfig   string   `yaml:"kubeconfig"`
	Devices      []string `yaml:"devices"`
	Namespace    string   `yaml:"namespace"`
}
