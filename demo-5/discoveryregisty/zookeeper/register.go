package zookeeper

import (
	"github.com/go-zookeeper/zk"
	"google.golang.org/grpc"
)

func (s *ZkClient) Register(rpcRegisterName, host string, port int) error {
	if err := s.ensureName(rpcRegisterName); err != nil {
		return err
	}
	addr := s.getAddr(host, port)
	node, err := s.conn.CreateProtectedEphemeralSequential(s.getPath(rpcRegisterName)+"/"+addr+"_", []byte(addr), zk.WorldACL(zk.PermAll))
	if err != nil {
		return err
	}
	s.node = node
	return nil
}

func (s *ZkClient) UnRegister() error {
	s.lock.Lock()
	defer s.lock.Unlock()
	err := s.conn.Delete(s.node, -1)
	if err != nil && err != zk.ErrNoNode {
		return err
	}
	s.node = ""
	s.localConns = make(map[string][]grpc.ClientConnInterface)
	s.resolvers = make(map[string]*Resolver)
	return nil
}
