package nacos

import (
	"errors"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
	"os"
	"path/filepath"
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
	timeout uint64
	// 服务注册名
	key     string
	closeCh chan struct{}

	// 服务发现专用
	resolvers    map[string]*Resolver
	localConns   map[string][]grpc.ClientConnInterface
	options      []grpc.DialOption
	balancerName string
	// 全局锁操作resolvers 和localConns
	lock sync.Locker
}

func (n *NacosRegister) Scheme() string {
	return strings.ToLower(n.schema)
}

func (n *NacosRegister) AddOption(opts ...grpc.DialOption) {
	n.options = append(n.options, opts...)
}

type NacosOption func(*NacosRegister)

func WithUserNameAndPassword(userName, password string) NacosOption {
	return func(n *NacosRegister) {
		n.userName = userName
		n.password = password
	}
}

func WithOptions(opts ...grpc.DialOption) NacosOption {
	return func(client *NacosRegister) {
		client.options = opts
	}
}

func WithRoundRobin() NacosOption {
	return func(client *NacosRegister) {
		client.balancerName = roundrobin.Name
	}
}

func WithTimeout(timeout uint64) NacosOption {
	return func(client *NacosRegister) {
		client.timeout = timeout
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
func NewClient(namespaceId string, nacosAddr []string, schema string, opts ...NacosOption) (*NacosRegister, error) {
	register := &NacosRegister{
		namespaceID: namespaceId,
		nacosAddr:   nacosAddr,
		timeout:     timeout,
		clusterName: "DEFAULT",
		groupName:   "DEFAULT_GROUP",
		schema:      schema,
		lock:        &sync.Mutex{},
		localConns:  make(map[string][]grpc.ClientConnInterface),
		resolvers:   make(map[string]*Resolver),
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
	currentProcessPath, _ := os.Executable()
	cacheDir := os.TempDir() + string(os.PathSeparator) + filepath.Base(currentProcessPath) + string(os.PathSeparator) + "cache" + string(os.PathSeparator) + namespaceId
	logDir := os.TempDir() + string(os.PathSeparator) + filepath.Base(currentProcessPath) + string(os.PathSeparator) + "logger" + string(os.PathSeparator) + namespaceId
	cc := *constant.NewClientConfig(
		// 命名空间ID
		constant.WithNamespaceId(register.namespaceID),
		// 请求Nacos服务端的超时时间
		constant.WithTimeoutMs(register.timeout),
		// 启动的时候不使用本地缓存加载服务信息数据
		constant.WithNotLoadCacheAtStart(true),
		// 缓存service信息的目录，默认是当前运行目录
		constant.WithCacheDir(cacheDir),
		// 日志存储路径
		constant.WithLogDir(logDir),
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
	resolver.Register(register)
	return register, nil
}

func (n *NacosRegister) GetClientLocalConns() map[string][]grpc.ClientConnInterface {
	return n.localConns
}

func (n *NacosRegister) flushResolverAndDeleteLocal(serviceName string) {
	fmt.Println("start flush ", serviceName)
	n.flushResolver(serviceName)
	delete(n.localConns, serviceName)
}

func (n *NacosRegister) flushResolver(serviceName string) {
	r, ok := n.resolvers[serviceName]
	if ok {
		r.ResolveNowNacos(resolver.ResolveNowOptions{})
	}
}
