package config

type Config struct {
	MasterAdress string `yaml:"master"`
	KubeConfig   string `yaml:"kubeconfig"`
	Device       string `yaml:"device"`
}
