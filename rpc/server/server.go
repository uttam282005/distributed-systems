package server

import (
	"net"
	"net/rpc"
	"sync"
)

type Kv struct {
	mu sync.Mutex
	data map[string]string
}

type GetArgs struct {
	key string
}

type PutArgs struct {
	key string
	value string
}

type PutReply struct{}

type GetReply struct {
	Value string
}

type Server interface {
	Get(key string, reply *GetReply) error
	Put(key, value string, reply *GetReply) error
}

func server() {
	KV := new(Kv)
	KV.data = map[string]string{}

	rpcs := rpc.NewServer()
	rpcs.Register(KV)

	l, e := net.Listen("tcp", ":8000")
	if e != nil {
		l.Close()
	}

	go func() {
		for {
			conn, e := l.Accept()
			if e != nil {
				conn.Close()
				break
			}
			go rpcs.ServeConn(conn)
		}
		l.Close()
	}()
}

func (KV *Kv) Get(args *GetArgs, reply *GetReply) error {
	KV.mu.Lock()
	reply.Value = KV.data[args.key]
	KV.mu.Unlock()
	return nil
}

func (KV *Kv) Put(args *PutArgs, reply *PutReply) error {
	KV.mu.Lock()
	KV.data[args.key] = args.value
	KV.mu.Unlock()
	return nil
}
