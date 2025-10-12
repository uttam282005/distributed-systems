// Package mr This file contains the implementation of the Coordinator for a MapReduce system.
package mr

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"sync"
)

type Job struct {
	JobID    int
	Type     string
	FileName string
}

type MetaData struct {
	Nreduce int
	Nmap    int
}

type Coordinator struct {
	mu sync.Mutex

	Jobs  []Job
	phase string

	NReduce int
	Nmap    int
}

// GetJob RPC handlers for the worker to call.
func(c *Coordinator) GetJob(args *GetJobArgs, reply *GetJobReply) error {

	return nil
}

// start a thread that listens for RPCs from worker.go
func (c *Coordinator) server() {
	rpc.Register(c)
	rpc.HandleHTTP()
	// l, e := net.Listen("tcp", ":1234")
	sockname := coordinatorSock()
	os.Remove(sockname)
	l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

// Done is called periodically by the workers to find out
// if the entire job has finished.
func (c *Coordinator) Done() bool {
	ret := false

	// Your code here.

	return ret
}

// MakeCoordinator creates a Coordinator.
func MakeCoordinator(files []string, nReduce int) *Coordinator {
	c := Coordinator{}

	c.server()
	return &c
}
