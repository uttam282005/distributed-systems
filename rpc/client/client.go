package client

import (
	"net/rpc"
	types "rpc-example/types"
)

type Client struct {
	client *rpc.Client
}

func Connect() (*Client, error) {
	client, e := rpc.Dial("tcp", "8000")
	if e != nil {
		return nil, nil
	}
	return  &Client{
		client,
	} , nil
}

func (c *Client) Get(key string) string {
	client := c.client
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

func (c *Client) Put(key string, val string) error {
	client := c.client
	defer client.Close()

	args := types.PutArgs{ Key: key, Value: val }
	reply := types.PutReply{}
	err := client.Call("KV.Put", &args, &reply)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Close() {
	c.client.Close()
}
