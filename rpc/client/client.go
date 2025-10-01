package client

import (
    "net/rpc"
    "rpc-example/types"
)

type KVClient struct {
    c *rpc.Client
}

func Connect() (*KVClient, error) {
    c, err := rpc.Dial("tcp", "localhost:8000")
    if err != nil {
        return nil, err
    }
    return &KVClient{c: c}, nil
}

func (kv *KVClient) Close() {
    kv.c.Close()
}

func (kv *KVClient) Get(key string) (string, error) {
    args := &types.GetArgs{Key: key}
    var reply types.GetReply
    err := kv.c.Call("Kv.Get", args, &reply)
    if err != nil {
        return "", err
    }
    return reply.Value, nil
}

func (kv *KVClient) Put(key, value string) error {
    args := &types.PutArgs{Key: key, Value: value}
    var reply types.PutReply
    err := kv.c.Call("Kv.Put", args, &reply)
    return err
}

