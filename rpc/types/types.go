package types

import "sync"

type Kv struct {
	mu   sync.Mutex
	data map[string]string
}

type GetArgs struct {
	Key string
}

type PutArgs struct {
	Key   string
	Value string
}

type PutReply struct{}

type GetReply struct {
	Value string
}
