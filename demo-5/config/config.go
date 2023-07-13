package config

import (
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	// Root folder of this project
	Root = filepath.Join(filepath.Dir(b), "../")
)

var Config config

type config struct {
	Etcd struct {
		EtcdSchema string   `yaml:"etcdSchema"`
		EtcdAddr   []string `yaml:"etcdAddr"`
		UserName   string   `yaml:"userName"`
		Password   string   `yaml:"password"`
	} `yaml:"etcd"`
}

func init() {
	configPath := Root + "/config/config.yaml"
	bytes, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(bytes, &Config)
	if err != nil {
		panic(err)
	}
}
