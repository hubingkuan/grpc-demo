package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

var typeSet = map[string]struct{}{
	"json":       {},
	"yaml":       {},
	"text":       {},
	"xml":        {},
	"html":       {},
	"properties": {},
}

// 这里需要注意如果方法的json数据 GetConf的时候 string无法反序列化成对应的数据结构  需要给param加Type
func (n *NacosRegister) RegisterConf2Registry(key string, conf []byte) error {
	param := n.generateConfigParam(key)
	param.Content = string(conf)
	_, err := n.configClient.PublishConfig(param)
	return err
}

func (n *NacosRegister) GetConfFromRegistry(key string) ([]byte, error) {
	param := n.generateConfigParam(key)
	content, err := n.configClient.GetConfig(param)
	return []byte(content), err
}

// 监听配置文件的变化
func (n *NacosRegister) ListenConfig(key string, functions func(namespace, group, dataId, data string)) error {
	param := n.generateConfigParam(key)
	param.OnChange = functions
	return n.configClient.ListenConfig(param)
}

// 取消监听配置文件
func (n *NacosRegister) CancelListenConfig(key string) error {
	param := n.generateConfigParam(key)
	return n.configClient.CancelListenConfig(param)
}

// 根据key生成configParam
func (n *NacosRegister) generateConfigParam(key string) vo.ConfigParam {
	param := vo.ConfigParam{}
	if n.groupName == "" {
		param.Group = constant.DEFAULT_GROUP
	} else {
		param.Group = n.groupName
	}
	param.DataId = key
	param.AppName = n.groupName
	return param
}
