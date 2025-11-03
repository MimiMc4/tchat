package coord

type Event interface {
	Type() string
}

// NODE JOIN EVENT
type NodeJoinEvent struct {
	name string
}

func (e NodeJoinEvent) Type() string {
	return "NodeJoinEvent"
}

// NODE LEAVE EVENT
type NodeLeaveEvent struct {
	name string
}

func (e NodeLeaveEvent) Type() string {
	return "NodeLeaveEvent"
}

// NEW MESSAGE EVENT
type NewMessageEvent struct {
	message ChatMessage
}

func (e NewMessageEvent) Type() string {
	return "NewMessageEvent"
}
