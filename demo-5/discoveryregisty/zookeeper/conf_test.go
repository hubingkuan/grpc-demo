package zookeeper

import (
	"encoding/json"
	"grpc-demo/demo-5/config"
	"testing"
)

func TestZKConfig(t *testing.T) {
	client, err := NewClient(config.Config.Zookeeper.Address, config.Config.Zookeeper.Schema, WithUserNameAndPassword(config.Config.Zookeeper.UserName, config.Config.Zookeeper.Password))
	if err != nil {
		t.Fatal(err)
	}
	marshal, _ := json.Marshal(config.Config)
	err = client.RegisterConf2Registry("test", marshal)
	if err != nil {
		t.Fatal(err)
	}

	bytes, err := client.GetConfFromRegistry("test")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))
}
