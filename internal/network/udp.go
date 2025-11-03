package network

import (
	"fmt"
	"net"
	"strconv"

	"github.com/mimimc4/tchat/pkg/config"
)

type UDPEndpoint struct {
	listener      *net.UDPConn
	broadcastConn *net.UDPConn
}

func (me *UDPEndpoint) Init() error {
	// Init listener
	addr := ":" + config.UDPPort

	UDPAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}

	me.listener, err = net.ListenUDP("udp", UDPAddr)
	if err != nil {
		return err
	}

	// Init broadcastConn
	UDPPortInt, _ := strconv.Atoi(config.UDPPort)
	bcastAddr := &net.UDPAddr{IP: net.IPv4bcast, Port: UDPPortInt}
	me.broadcastConn, err = net.DialUDP("udp4", nil, bcastAddr)
	if err != nil {
		return fmt.Errorf("error creando conexión de broadcast: %w", err)
	}

	return nil
}

func (me *UDPEndpoint) BroadcastMsg(message []byte) error {
	_, err := me.broadcastConn.Write(message)
	return err
}

// Listens for udp messages and calls callback on message received
func (me *UDPEndpoint) Listen(callback func([]byte, *net.UDPAddr)) {
	buf := make([]byte, 1024)

	for {
		n, remoteAddr, err := me.listener.ReadFromUDP(buf)
		if err != nil {
			continue
		}

		callback(buf[:n], remoteAddr)
	}
}

func (me *UDPEndpoint) Close() {
	if me.listener != nil {
		me.listener.Close()
	}
	if me.broadcastConn != nil {
		me.broadcastConn.Close()
	}
}
