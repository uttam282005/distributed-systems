package client

import (
	"net/rpc"
	types "rpc-example/types"
)

func connect() *rpc.Client {
	client, e := rpc.Dial("tcp", "8000")
	if e != nil {
		client.Close()
	}
	return client
}

func Get(key string) string {
	client := connect()
	defer client.Close()

	args := types.GetArgs {
		Key: key,
	} 
	reply := types.GetReply{}
	e := client.Call("KV.Get", &args, &reply)
	if e != nil {
		return "" 
	}
	return reply.Value
}

func Put(key string, val string) error {
	client := connect()
	defer client.Close()

	args := types.PutArgs{ Key: key, Value: val }
	reply := types.PutReply{}
	err := client.Call("KV.Put", &args, &reply)
	if err != nil {
		return err
	}
	return nil
}
