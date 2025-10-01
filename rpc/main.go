package main

import (
	// client "rpc-example/client"
	server "rpc-example/server"
	// "fmt"
)

func main() {
	server.Start()
	// kv, err := client.Connect()
	// if err != nil {
	// 	kv.Close()
	// }
	//
	// kv.Put("Hello", "world")
	//
	// fmt.Println(kv.Get("Hello"))
}
