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
	root = filepath.Join(filepath.Dir(b), "../")
)

var Config config

type config struct {
	Etcd struct {
		Schema   string   `yaml:"schema"`
		Address  []string `yaml:"address"`
		UserName string   `yaml:"username"`
		Password string   `yaml:"password"`
	} `yaml:"etcd"`

	Zookeeper struct {
		Schema   string   `yaml:"schema"`
		Address  []string `yaml:"address"`
		UserName string   `yaml:"username"`
		Password string   `yaml:"password"`
	} `yaml:"zookeeper"`

	Nacos struct {
		Schema      string   `yaml:"schema"`
		Address     []string `yaml:"address"`
		NamespaceID string   `yaml:"namespace"`
		UserName    string   `yaml:"username"`
		Password    string   `yaml:"password"`
	} `yaml:"nacos"`
}

func init() {
	configPath := root + "/config/config.yaml"
	bytes, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(bytes, &Config)
	if err != nil {
		panic(err)
	}
}
