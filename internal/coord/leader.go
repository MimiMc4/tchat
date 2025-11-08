package coord

import (
	"time"

	"github.com/mimimc4/tchat/internal/config"
)

func (me *Node) callRemoteAndCheck(node *RemoteNode, funcName string, args interface{}, results interface{}) {
	err := node.Endpoint.CallRemote(funcName, args, results, time.Duration(config.RPCDuration)*time.Millisecond)
	if err != nil {
		me.Mu.Lock()
		me.Leader.MissedCalls[node.ID]++
		if me.Leader.MissedCalls[node.ID] > config.MaxMissedCalls {
			argExit := ArgExit{ID: node.ID}
			_ = me.Exit(&argExit, &Empty{})
		}
		me.Mu.Unlock()
	}
}
