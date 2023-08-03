package nacos

import (
	"errors"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"google.golang.org/grpc"
	"strconv"
	"strings"
	"sync"
)

const (
	timeout      = 5
	beatInterval = 5 * 1000
)

type NacosRegister struct {
	// 动态配置客户端
	configClient config_client.IConfigClient
	// 服务发现客户端
	namingClient naming_client.INamingClient

	nacosAddr   []string
	userName    string
	password    string
	namespaceID string
	clusterName string
	groupName   string
	schema      string
	// 连接超时时间
	timeout      uint64
	beatInterval int64
	// 服务注册名
	key string

	// 服务发现专用
	resolvers    map[string]*Resolver
	localConns   map[string][]grpc.ClientConnInterface
	options      []grpc.DialOption
	balancerName string
	// 全局锁操作resolvers 和localConns
	lock sync.Locker
}

func (n NacosRegister) Scheme() string {
	return strings.ToLower(n.schema)
}

func (n NacosRegister) AddOption(opts ...grpc.DialOption) {
	// TODO implement me
	panic("implement me")
}

type NacosOption func(*NacosRegister)

func WithUserNameAndPassword(userName, password string) NacosOption {
	return func(n *NacosRegister) {
		n.userName = userName
		n.password = password
	}
}

func WithTimeout(timeout uint64) NacosOption {
	return func(client *NacosRegister) {
		client.timeout = timeout
	}
}

func WithBeatInterval(beatInterval int64) NacosOption {
	return func(client *NacosRegister) {
		client.beatInterval = beatInterval
	}
}

func WithGroupName(groupName string) NacosOption {
	return func(client *NacosRegister) {
		client.groupName = groupName
	}
}

func WithClusterName(clusterName string) NacosOption {
	return func(client *NacosRegister) {
		client.clusterName = clusterName
	}
}

// namespaceID 需要提前设置好
func NewClient(namespaceId string, nacosAddr []string, opts ...NacosOption) (*NacosRegister, error) {
	register := &NacosRegister{
		namespaceID:  namespaceId,
		nacosAddr:    nacosAddr,
		timeout:      timeout,
		beatInterval: beatInterval,
	}
	for _, opt := range opts {
		opt(register)
	}

	// create ServerConfig
	sc := make([]constant.ServerConfig, 0, len(nacosAddr))
	for i := range nacosAddr {
		addrInfo := strings.Split(nacosAddr[i], ":")
		if len(addrInfo) != 2 {
			return nil, errors.New("nacosAddr format error")
		}
		port, err := strconv.Atoi(addrInfo[1])
		if err != nil {
			return nil, err
		}
		sc = append(sc, *constant.NewServerConfig(addrInfo[0], uint64(port)))
	}

	// create ClientConfig
	cc := *constant.NewClientConfig(
		// 命名空间ID
		constant.WithNamespaceId(register.namespaceID),
		// 请求Nacos服务端的超时时间
		constant.WithTimeoutMs(register.timeout),
		// 向服务器发送心跳的时间间隔，默认值为5000ms
		constant.WithBeatInterval(register.beatInterval),
		// 启动的时候不使用本地缓存加载服务信息数据
		constant.WithNotLoadCacheAtStart(true),
		// 缓存service信息的目录，默认是当前运行目录
		constant.WithCacheDir("/tmp/nacos/cache"),
		// 日志存储路径
		constant.WithLogDir("/tmp/nacos/log"),
		// 日志默认级别 默认INFO
		constant.WithLogLevel("debug"),
		// 服务端API鉴权的用户名密码
		constant.WithUsername(register.userName),
		constant.WithPassword(register.password),
		// 如果在config 中开启了鉴权的话
		// constant.WithAccessKey(),
		// constant.WithSecretKey()
	)

	// 创建动态配置客户端
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		return nil, err
	}

	// 创建服务发现客户端
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		return nil, err
	}
	register.configClient = configClient
	register.namingClient = namingClient
	// resolver.Register(register)
	return register, nil
}

func (n NacosRegister) GetClientLocalConns() map[string][]grpc.ClientConnInterface {
	// TODO implement me
	panic("implement me")
}
