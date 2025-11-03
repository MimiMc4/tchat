package network

import (
	"fmt"
	"net"
	"strconv"

	"github.com/mimimc4/tchat/pkg/config"
)

type UDPEndpoint struct {
	addr          *net.UDPAddr
	listener      *net.UDPConn
	broadcastConn *net.UDPConn
}

// Initializes udp endpoint and sets up message listener goroutine
func (me *UDPEndpoint) Init() error {
	// Init listener
	addr := ":" + config.UDPPort

	UDPAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}
	me.addr = UDPAddr

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

func (me *UDPEndpoint) Listen() ([]byte, *net.UDPAddr, error) {
	buf := make([]byte, 1024)

	n, remoteAddr, err := me.listener.ReadFromUDP(buf)
	if err != nil {
		return nil, nil, err
	}

	return buf[:n], remoteAddr, nil
}

func (me *UDPEndpoint) Close() {
	if me.listener != nil {
		me.listener.Close()
	}
	if me.broadcastConn != nil {
		me.broadcastConn.Close()
	}
}
