package zookeeper

import (
	"github.com/go-zookeeper/zk"
	"google.golang.org/grpc"
)

func (s *ZkClient) Register(rpcRegisterName, host string, port int) error {
	// 先创建 /openIM/rpcRegisterName  父节点
	if err := s.ensureName(rpcRegisterName); err != nil {
		return err
	}
	addr := s.getAddr(host, port)
	_, err := grpc.Dial(addr)
	if err != nil {
		return err
	}
	// 再创建顺序临时节点 同时还具有保护性(确保断开重连后临时节点可以和客户端状态对接上)
	node, err := s.CreateTempNode(rpcRegisterName, addr)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	s.node = node
	s.rpcRegisterName = rpcRegisterName
	s.rpcRegisterAddr = addr
	s.isRegistered = true
	return nil
}

func (s *ZkClient) UnRegister() error {
	err := s.conn.Delete(s.node, -1)
	if err != nil && err != zk.ErrNoNode {
		return err
	}
	s.node = ""
	s.rpcRegisterName = ""
	s.rpcRegisterAddr = ""
	s.isRegistered = false
	s.localConns = make(map[string][]grpc.ClientConnInterface)
	s.resolvers = make(map[string]*Resolver)
	return nil
}

func (s *ZkClient) CreateTempNode(rpcRegisterName, addr string) (node string, err error) {
	return s.conn.CreateProtectedEphemeralSequential(
		s.getPath(rpcRegisterName)+"/"+addr+"_",
		[]byte(addr),
		zk.WorldACL(zk.PermAll),
	)
}
