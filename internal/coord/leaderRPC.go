package coord

type ArgExit struct {
	ID int
}

func (me *Node) Exit(args *ArgExit, results *Empty) error {
	// Copy the remote node info first and the make the RPC
	me.Mu.RLock()

	nodesToCall := make([]*RemoteNode, 0, len(me.Endpoints))
	for id, remoteNode := range me.Endpoints {
		if id != args.ID {
			nodesToCall = append(nodesToCall, remoteNode)
		}
	}

	me.Mu.RUnlock()

	leaveArgs := ArgRemoveParticipant{
		ID: args.ID,
	}

	for _, remoteNode := range nodesToCall {
		me.callRemoteAndCheck(remoteNode, "Node.RemoveParticipant", leaveArgs, Empty{})
	}

	_ = me.RemoveParticipant(&leaveArgs, &Empty{})

	return nil
}
