package coord

import (
	"sync"
	"time"

	"github.com/mimimc4/tchat/internal/config"
	"github.com/mimimc4/tchat/internal/network"
	"github.com/mimimc4/tchat/pkg/utils"
)

type ChatMessage struct {
	sender  string
	message string
	time    time.Time // use Clock() to get hours, mins and sec
}

type LeaderInfo struct{}

type RemoteNode struct {
	name     string
	endpoint network.RPCEndpoint
}

type Node struct {
	mu sync.RWMutex

	id        int
	name      string
	msgBuffer *utils.CircularBuffer[ChatMessage]
	endpoints map[int]RemoteNode
	eventChan chan Event
}

func NewNode(id int, name string) *Node {
	return &Node{
		id:        id,
		name:      name,
		msgBuffer: utils.NewCircularBuffer[ChatMessage](config.BuffSize),
		endpoints: make(map[int]RemoteNode),
		eventChan: make(chan Event, 64),
	}
}
