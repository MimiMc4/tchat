package coord

import (
	"sync"
	"time"

	"github.com/mimimc4/tchat/internal/config"
	"github.com/mimimc4/tchat/internal/network"
	"github.com/mimimc4/tchat/pkg/utils"
)

type ChatMessage struct {
	Sender  string
	Message string
	Time    time.Time // use Clock() to get hours, mins and sec
}

func Equal(a, b ChatMessage) bool {
	return a.Sender == b.Sender && a.Time.Equal(b.Time)
}

type LeaderInfo struct{}

type RemoteNode struct {
	Name     string
	Endpoint network.RPCEndpoint
}

type Timeout struct {
	HeartbeatTimeout chan bool
	ElectionTimeout  chan bool
}

type Node struct {
	Mu sync.RWMutex

	ID        int
	Name      string
	IsLeader  bool
	LeaderID  int
	MsgBuffer *utils.CircularBuffer[ChatMessage]
	Endpoints map[int]RemoteNode
	Timeout   Timeout
	EventChan chan Event
}

func NewNode(id int, name string) *Node {
	return &Node{
		ID:        id,
		Name:      name,
		IsLeader:  false,
		LeaderID:  -1,
		MsgBuffer: utils.NewCircularBuffer[ChatMessage](config.BuffSize, Equal),
		Endpoints: make(map[int]RemoteNode),
		Timeout: Timeout{
			HeartbeatTimeout: make(chan bool),
			ElectionTimeout:  make(chan bool),
		},
		EventChan: make(chan Event, 64),
	}
}
