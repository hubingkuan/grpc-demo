package zookeeper

import (
	"fmt"
	"github.com/go-zookeeper/zk"
	"grpc-demo/demo-5/config"
	"testing"
)

func TestZKRegister(t *testing.T) {
	client, err := NewClient(config.Config.Zookeeper.Address, config.Config.Zookeeper.Schema, WithUserNameAndPassword(config.Config.Zookeeper.UserName, config.Config.Zookeeper.Password))
	if err != nil {
		t.Fatal(err)
	}

	path := "/gozk-test"

	// 这里的-1 代表的是期望删除的节点的期望版本号(乐观锁)  当version为-1的时候 表示忽略版本检查 直接删除节点
	if err := client.conn.Delete(path, -1); err != nil && err != zk.ErrNoNode {
		t.Fatalf("Delete returned error: %+v", err)
	}

	if p, err := client.conn.Create(path, []byte("test"), 0, zk.WorldACL(zk.PermAll)); err != nil && err != zk.ErrNodeExists {
		t.Fatalf("Create returned error: %+v", err)
	} else {
		// p返回的是节点路径  如果是已创建则p为""
		fmt.Println(p)
	}

	// data代表节点数据
	// stata代表节点的原数据信息(version Ctime DataLength)
	if data, stat, err := client.conn.Get(path); err != nil {
		t.Fatalf("Get returned error: %+v", err)
	} else if stat == nil {
		t.Fatal("Get returned nil stat")
	} else {
		fmt.Println(string(data))
	}

	// w获取所有的子节点名字
	// stat获取节点的原数据信息
	// events
	w, stat, events, err := client.conn.ChildrenW("/openIM/helloServer")
	if err != nil {
		t.Fatal(err)
	} else {
		fmt.Println(w)
		fmt.Println(stat)
		for event := range events {
			fmt.Println("event:", event)
		}
	}

	// parent: /openIM/hello
	// children: [_c_969474d6afef1bd41c41173fde95629d-127.0.0.1:8888_0000000001]

	// client.Register("hello", "127.0.0.1", 8888)
}
