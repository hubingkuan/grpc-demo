package zookeeper

import (
	"grpc-demo/demo-5/config"
	"testing"
)

func TestZKRegister(t *testing.T) {
	client, err := NewClient(config.Config.Zookeeper.Address, config.Config.Zookeeper.Schema, WithUserNameAndPassword(config.Config.Zookeeper.UserName, config.Config.Zookeeper.Password))
	if err != nil {
		t.Fatal(err)
	}
	// parent: /openIM/hello
	// children: [_c_969474d6afef1bd41c41173fde95629d-127.0.0.1:8888_0000000001]
	client.Register("hello", "127.0.0.1", 8888)
}
