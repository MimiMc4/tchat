package model

import "sync"

type Node struct {
	mtx sync.Mutex

	leaderID  int
	endpoints map[int]Endpoint
}
