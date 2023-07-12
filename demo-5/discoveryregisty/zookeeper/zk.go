package zookeeper

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"sync"
	"time"

	"github.com/go-zookeeper/zk"
)

const (
	defaultFreq = time.Minute * 30
	timeout     = 5
)

type Logger interface {
	Printf(string, ...interface{})
}

type ZkClient struct {
	// zk server地址
	zkServers []string
	zkRoot    string
	userName  string
	password  string

	scheme string

	timeout   int
	conn      *zk.Conn
	eventChan <-chan zk.Event
	node      string
	ticker    *time.Ticker

	lock    sync.Locker
	options []grpc.DialOption

	resolvers    map[string]*Resolver
	localConns   map[string][]resolver.Address
	balancerName string

	logger Logger
}
