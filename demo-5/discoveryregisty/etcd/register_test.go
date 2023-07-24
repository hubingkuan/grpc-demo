package etcd

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"grpc-demo/demo-5/config"
	"sync"
	"testing"
	"time"
)

func TestEtcdClient_Register(t *testing.T) {
	type fields struct {
		cli          *clientv3.Client
		schema       string
		key          string
		closeCh      chan struct{}
		keepAliveCh  <-chan *clientv3.LeaseKeepAliveResponse
		etcdAddr     []string
		lock         sync.Locker
		options      []grpc.DialOption
		resolvers    map[string]*Resolver
		localConns   map[string][]resolver.Address
		balancerName string
	}
	type args struct {
		serviceName string
		host        string
		port        int
		opts        []grpc.DialOption
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			args: args{
				serviceName: "maguahu",
				host:        "127.0.0.1",
				port:        333,
				opts:        nil,
			},
		},
	}
	r, _ := NewClient(config.Config.Etcd.Address, config.Config.Etcd.Schema)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := r.Register(tt.args.serviceName, tt.args.host, tt.args.port); (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	time.Sleep(time.Second * 10)
	r.UnRegister()
}
