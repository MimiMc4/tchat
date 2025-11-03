package network

import (
	"fmt"
	"net"
	"net/rpc"
	"time"

	"github.com/mimimc4/tchat/pkg/config"
)

type RPCEndpoint string

func InitRPCServer(service any) (net.Listener, error) {
	addr := ":" + config.TCPPort

	if err := rpc.Register(service); err != nil {
		return nil, err
	}

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	go rpc.Accept(listener)

	return listener, nil
}

func (me RPCEndpoint) CallRemote(serviceMethod string, args interface{},
	reply interface{}, timeout time.Duration,
) error {
	client, err := rpc.Dial("tcp", string(me))
	if err != nil {
		return err
	}
	defer client.Close()

	call := client.Go(serviceMethod, args, reply, make(chan *rpc.Call, 1))

	select {
	case result := <-call.Done:
		return result.Error
	case <-time.After(timeout):
		return fmt.Errorf("timeout on %s call to %s", serviceMethod, string(me))
	}
}
