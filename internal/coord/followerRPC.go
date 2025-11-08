package coord

import (
	"github.com/mimimc4/tchat/internal/network"
)

type Empty struct{}

// ADD PARTICIPANT ====================================================================
type ArgAddParticipant struct {
	ID      int
	Name    string
	Enpoint network.RPCEndpoint
}

func (me *Node) AddParticipant(args *ArgAddParticipant, results *Empty) error {
	me.Mu.Lock()
	defer me.Mu.Unlock()

	newNode := RemoteNode{
		ID:       args.ID,
		Name:     args.Name,
		Endpoint: args.Enpoint,
	}
	me.Endpoints[args.ID] = &newNode

	*results = Empty{}

	me.EventChan <- NodeJoinEvent{name: args.Name}

	return nil
}

// REMOVE PARTICIPANT ====================================================================
type ArgRemoveParticipant struct {
	ID int
}

func (me *Node) RemoveParticipant(args *ArgRemoveParticipant, results *Empty) error {
	me.Mu.Lock()
	defer me.Mu.Unlock()

	leaveEvent := NodeLeaveEvent{
		name: me.Endpoints[args.ID].Name,
	}

	delete(me.Endpoints, args.ID)

	*results = Empty{}

	me.EventChan <- leaveEvent

	return nil
}

// RECEIVE MESSAGE ====================================================================
type ArgReceiveMessage struct {
	Message ChatMessage
}

func (me *Node) ReceiveMessage(args *ArgReceiveMessage, results *Empty) error {
	me.Mu.Lock()
	defer me.Mu.Unlock()

	me.MsgBuffer.Add(args.Message)

	*results = Empty{}

	me.EventChan <- NewMessageEvent{message: args.Message}

	return nil
}

// RECEIVE HEARTBEAT ====================================================================
func (me *Node) ReceiveHeartbeat(args *Empty, results *Empty) error {
	me.Timeout.HeartbeatTimeout <- true

	*results = Empty{}

	return nil
}

// ELECTION ====================================================================
type ArgElection struct {
	ID int
}

type ResultElection struct {
	Vote bool
}

func (me *Node) Election(args *ArgElection, results *ResultElection) error {
	me.Mu.Lock()
	defer me.Mu.Unlock()

	var vote ResultElection
	if me.ID < args.ID {
		// I have priority -> try to become leader
		vote.Vote = false

		// StartElection()
	} else {
		vote.Vote = true
	}

	// An election means the leader died -> remove from endpoints
	if me.LeaderID != -1 {
		me.EventChan <- NodeLeaveEvent{name: me.Endpoints[me.LeaderID].Name}
		delete(me.Endpoints, me.LeaderID)
		me.LeaderID = -1
	}

	*results = vote

	return nil
}

// SYNC ====================================================================
type ArgSync struct {
	messages []ChatMessage
}

func (me *Node) Sync(args *ArgSync, results *Empty) error {
	me.Mu.Lock()
	defer me.Mu.Unlock()

	newMessages := args.messages

	for _, message := range newMessages {
		if !me.MsgBuffer.Contains(message) {
			me.MsgBuffer.Add(message)
			me.EventChan <- NewMessageEvent{message: message}
		}
	}

	return nil
}
