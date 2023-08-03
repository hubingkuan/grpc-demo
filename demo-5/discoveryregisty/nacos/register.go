package nacos

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func (n NacosRegister) Register(serviceName, host string, port int) error {
	success, err := n.namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          host,
		Port:        uint64(port),
		ServiceName: serviceName,
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   false,
		// 元数据信息
		// Metadata:    map[string]string{"idc": "shanghai"},
		GroupName:   n.groupName,
		ClusterName: n.clusterName,
	})
	if err != nil {
		return err
	}
	fmt.Printf("Register ServiceName:%s,state:%t\n", serviceName, success)
	n.key = serviceName

	services, err := n.namingClient.GetService(vo.GetServiceParam{
		ServiceName: serviceName,
		GroupName:   n.groupName,
	})
	fmt.Println("services:", services)
	return nil
}

func (n NacosRegister) UnRegister() error {
	// todo 弄明白 Ephemeral的作用
	success, err := n.namingClient.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          "127.0.0.1",
		Port:        7777,
		ServiceName: "helloServer",
		GroupName:   n.groupName,
		Cluster:     n.clusterName,
	})
	fmt.Printf("UnRegister ServiceName:%s,state:%t\n", n.key, success)
	return err
}
