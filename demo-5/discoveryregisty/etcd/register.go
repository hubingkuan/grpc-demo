package etcd

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"net"
	"strconv"
	"time"
)

const default_lease_ttl = 5

func (r *EtcdRegister) Register(serviceName, host string, port int) error {
	ctx := context.Background()
	args := fmt.Sprintf("RegisterEtcd args: schema:%s,serviceName:%s,host:%s,port:%d", r.schema, serviceName, host, port)
	fmt.Println(args)

	serviceValue := net.JoinHostPort(host, strconv.Itoa(port))
	serviceKey := GetPrefix(r.schema, serviceName) + serviceValue
	r.key = serviceKey
	// 授权租约
	resp, err := r.cli.Grant(ctx, default_lease_ttl)
	if err != nil {
		fmt.Println("Grant failed ", err.Error())
		return err
	}
	fmt.Println("Grant ok,resp ID ", resp.ID)
	// key:   schema:///serviceName/ip:port
	// value:  ip:port
	// 服务注册
	if _, err := r.cli.Put(ctx, serviceKey, serviceValue, clientv3.WithLease(resp.ID)); err != nil {
		fmt.Println("cli.Put failed ", err.Error())
		return fmt.Errorf("put failed, errmsg:%v， key:%s, value:%s", err, serviceKey, serviceValue)
	}
	// 租约续租
	r.keepAliveCh, err = r.cli.KeepAlive(ctx, resp.ID)
	if err != nil {
		fmt.Println("KeepAlive failed ", err.Error(), args, resp.ID)
		return fmt.Errorf("keepalive failed, errmsg:%v, lease id:%d", err, resp.ID)
	}
	fmt.Println("RegisterEtcd ok ", args)
	r.closeCh = make(chan struct{})

	go func() {
		ticker := time.NewTicker(time.Duration(default_lease_ttl+1) * time.Second)
		for {
			select {
			// 收到注销通知后 取消授权租约
			case <-r.closeCh:
				fmt.Println("unregister")
				if _, err = r.cli.Revoke(context.Background(), resp.ID); err != nil {
					fmt.Println("Revoke failed error", resp.ID, err)
				}
				goto END
			// 收到续租通知
			case res := <-r.keepAliveCh:
				if res != nil {
					fmt.Println("续租:", res.ID)
				}
			case <-ticker.C:
				if r.keepAliveCh == nil {
					// 定时器时间到了但是续租通知没有收到 此时重新进行续租
					ctx, cancel := context.WithCancel(context.Background())
					defer cancel()
					resp, err := r.cli.Grant(ctx, int64(default_lease_ttl))
					if err != nil {
						fmt.Println("Grant failed ", err.Error(), args)
						continue
					}
					if _, err := r.cli.Put(ctx, serviceKey, serviceValue, clientv3.WithLease(resp.ID)); err != nil {
						fmt.Println("etcd Put failed ", err.Error(), args, " resp ID: ", resp.ID)
						continue
					} else {
						fmt.Println("etcd Put ok ", args, " resp ID: ", resp.ID)
					}
				}
			}
		}
	END:
		ticker.Stop()
	}()
	return err
}

func (r *EtcdRegister) UnRegister() {
	r.closeCh <- struct{}{}
}
