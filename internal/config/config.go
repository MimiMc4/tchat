// Package config includes different constant for all the codebase
package config

const (
	UDPPort string = "30000"
	TCPPort string = "30001"

	BuffSize = 5

	RPCDuration = 300

	HeartbeatDuration = 200
	ElectionDuration  = 600
	RandomDuration    = 300

	MaxMissedCalls = 5
)
