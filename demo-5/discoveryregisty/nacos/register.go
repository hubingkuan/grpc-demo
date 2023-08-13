package nacos

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

func (n NacosRegister) Register(serviceName, host string, port int) error {
	addr := net.JoinHostPort(host, strconv.Itoa(port))
	_, err := grpc.Dial(addr)
	if err != nil {
		return err
	}
	_, err = n.namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          host,
		Port:        uint64(port),
		ServiceName: serviceName,
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		// 临时节点
		Ephemeral: true,
		// 元数据信息
		// Metadata:    map[string]string{"idc": "shanghai"},
		GroupName:   n.groupName,
		ClusterName: n.clusterName,
	})
	if err != nil {
		return fmt.Errorf("nacos register instance error: %w", err)
	}
	n.closeCh = make(chan struct{})
	go func() {
		for {
			select {
			case <-n.closeCh:
				fmt.Println("unregister")
				if _, err = n.namingClient.DeregisterInstance(vo.DeregisterInstanceParam{
					Ip:          "127.0.0.1",
					Port:        7777,
					ServiceName: "helloServer",
					Ephemeral:   true,
					GroupName:   n.groupName,
					Cluster:     n.clusterName,
				}); err != nil {
					fmt.Println("nacos deregister instance error: %w", err)
				}
				return
			default:
			}
		}
	}()
	return nil
}

func (n NacosRegister) UnRegister() error {
	n.closeCh <- struct{}{}
	return nil
}
