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
	Key string
}

type PutArgs struct {
	Key string
	Value string
}

type PutReply struct{}

type GetReply struct {
	Value string
}

type Server_t interface {
	Get(key string, reply *GetReply) error
	Put(key, value string, reply *GetReply) error
}

func Server() {
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
	defer KV.mu.Unlock()

	reply.Value = KV.data[args.Key]
	return nil
}

func (KV *Kv) Put(args *PutArgs, reply *PutReply) error {
	KV.mu.Lock()
	defer KV.mu.Unlock()

	KV.data[args.Key] = args.Value
	return nil
}
