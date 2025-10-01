package server

import (
	"net"
	"net/rpc"
)

type Kv struct {
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
			}

		}
	}()
}

func (KV *Kv) Get(args *GetArgs, reply *GetReply) error {
	reply.Value = KV.data[args.key]
	return nil
}

func (KV *Kv) Put(args *PutArgs, reply *PutReply) error {
	KV.data[args.key] = args.value
	return nil
}
