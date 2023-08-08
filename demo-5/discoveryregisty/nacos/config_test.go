package nacos

import (
	"encoding/json"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"grpc-demo/demo-5/config"
	"testing"
)

func TestNacosRegister_RegisterConf2Registry(t *testing.T) {
	m := map[string]interface{}{"name": "test", "age": 20}
	bytes, _ := json.Marshal(m)
	type fields struct {
		client      config_client.IConfigClient
		userName    string
		password    string
		nameSpaceID string
		nacosAddr   string
		timeout     uint64
	}
	type args struct {
		key  string
		conf []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			fields: fields{},
			args: args{
				key:  "jsonData444",
				conf: bytes,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n, _ := NewClient(config.Config.Nacos.NamespaceID, config.Config.Nacos.Address, config.Config.Nacos.Schema, WithGroupName("grpc-demo5"))
			// 监听config变化
			// n.ListenConfig(tt.args.key, func(namespace, group, dataId, data string) {
			// 	t.Log("namespace:", namespace, "group:", group, "dataId:", dataId, "data:", data)
			// })

			// 注册配置
			// if err := n.RegisterConf2Registry(tt.args.key, tt.args.conf); (err != nil) != tt.wantErr {
			// 	t.Errorf("RegisterConf2Registry() error = %v, wantErr %v", err, tt.wantErr)
			// }

			// 获取配置
			// bytes, err := n.GetConfFromRegistry(tt.args.key)
			// if err != nil {
			// 	t.Error(err)
			// }
			// var mapResult map[string]interface{}
			// err = json.Unmarshal(bytes, &mapResult)
			// if err != nil {
			// 	t.Error(err)
			// }
			// t.Log(m)

			// 服务注册
			// n.Register("helloServer", "127.0.0.1", 7777)

			// 服务注销
			n.namingClient.DeregisterInstance(vo.DeregisterInstanceParam{
				Ip:          "127.0.0.1",
				Port:        7777,
				ServiceName: "helloServer",
				GroupName:   n.groupName,
				Cluster:     n.clusterName,
			})
		})
	}
}
