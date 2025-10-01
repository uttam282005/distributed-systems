package server

import (
	"log"
	"net"
	"net/rpc"
	"rpc-example/types"
	"sync"
)

type Kv struct {
	mu   sync.Mutex
	Data map[string]string
}

func Start() {
	KV := &Kv{Data: map[string]string{}}

	rpcs := rpc.NewServer()
	err := rpcs.Register(KV)
	if err != nil {
		log.Fatal("Register error:", err)
	}

	l, e := net.Listen("tcp", ":8000")
	if e != nil {
		log.Fatal("Listen error:", e)
	}

	go func() {
		for {
			conn, e := l.Accept()
			if e != nil {
				break
			}
			go rpcs.ServeConn(conn)
		}
		l.Close()
	}()
}

func (KV *Kv) Get(args *types.GetArgs, reply *types.GetReply) error {
	KV.mu.Lock()
	defer KV.mu.Unlock()
	reply.Value = KV.Data[args.Key]
	return nil
}

func (KV *Kv) Put(args *types.PutArgs, reply *types.PutReply) error {
	KV.mu.Lock()
	defer KV.mu.Unlock()
	KV.Data[args.Key] = args.Value
	return nil
}
