package coord

import "github.com/mimimc4/tchat/internal/network"

type Empty struct{}

// ADD PARTICIPANT ====================================================================
type ArgAddParticipant struct {
	id      int
	name    string
	enpoint network.RPCEndpoint
}

func (me *Node) AddParticipant(args *ArgAddParticipant, results *Empty) {
	me.mu.Lock()
	defer me.mu.Unlock()

	newNode := RemoteNode{
		name:     args.name,
		endpoint: args.enpoint,
	}
	me.endpoints[args.id] = newNode

	*results = Empty{}

	me.eventChan <- NodeJoinEvent{name: args.name}
}

// REMOVE PARTICIPANT ====================================================================
type ArgRemoveParticipant struct {
	id int
}

func (me *Node) RemoveParticipant(args *ArgRemoveParticipant, results *Empty) {
	me.mu.Lock()
	defer me.mu.Unlock()

	leaveEvent := NodeLeaveEvent{
		name: me.endpoints[args.id].name,
	}

	delete(me.endpoints, args.id)

	*results = Empty{}

	me.eventChan <- leaveEvent
}
