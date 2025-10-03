package main

import (
    "fmt"
    "log"
    "rpc-example/client"
    "rpc-example/server"
    "time"
)

func main() {
    server.Start()

    // Give server a moment to start
    time.Sleep(500 * time.Millisecond)

    kv, err := client.Connect()
    if err != nil {
        log.Fatal("Connect error:", err)
    }
    defer kv.Close()

    err = kv.Put("Hello", "world")
    if err != nil {
        log.Fatal("Put error:", err)
    }

    value, err := kv.Get("Hello")
    if err != nil {
        log.Fatal("Get error:", err)
    }

    fmt.Println("Get(Hello) =", value) 
}

