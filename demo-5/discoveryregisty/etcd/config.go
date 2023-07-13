package etcd

import (
	"context"
	"errors"
)

func (r *EtcdRegister) RegisterConf2Registry(key string, conf []byte) error {
	_, err := r.cli.Put(context.Background(), key, string(conf))
	if err != nil {
		return err
	}

	return nil
}

func (r *EtcdRegister) GetConfFromRegistry(key string) ([]byte, error) {
	resp, err := r.cli.Get(context.Background(), key)
	if err != nil {
		return nil, err
	}

	if len(resp.Kvs) == 0 {
		return nil, errors.New("get config by key no data")
	}

	return resp.Kvs[0].Value, nil
}
