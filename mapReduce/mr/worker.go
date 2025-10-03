package mr

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"log"
	"net/rpc"
	"os"
)

// Map functions return a slice of KeyValue.
type KeyValue struct {
	Key   string
	Value string
}

type GetJobArgs struct{}

type GetJobReply struct {
	FileName string
	Content  string
}

// use ihash(key) % NReduce to choose the reduce
// task number for each KeyValue emitted by Map.
func ihash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32() & 0x7fffffff)
}

func check(err any) {
	if err != nil {
		log.Fatal(err)
	}
}

func handleMapJob(fileName string, content string, mapf func(string, string) []KeyValue, mapTaskNumber int) bool {
	mapgoResult := mapf(fileName, content)
	reduceTaskNumber := ihash(fileName)
	f, err := os.Create(fmt.Sprintf("mr-%d-%d", mapTaskNumber, reduceTaskNumber))
	if err != nil {
		log.Fatal("cannot create map output file")
		return false
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	for _, kv := range mapgoResult {
		err := enc.Encode(&kv)
		if err != nil {
			log.Fatal("cannot encode kv pair")
			return false
		}
	}
	return true
}

func Worker(mapf func(string, string) []KeyValue,
	reducef func(string, []string) string,
) {
	mapTaskNumber := 0
	args := GetJobArgs{}
	reply := GetJobReply{}

	for {
		// ask for a map job
		ok := call("Coordinator.GetJob", &args, &reply)
		if ok {
			mapTaskNumber++
			go handleMapJob(reply.FileName, reply.Content, mapf, mapTaskNumber)
		} else {
			break
		}
	}
}

// send an RPC request to the coordinator, wait for the response.
// usually returns true.
// returns false if something goes wrong.
func call(rpcname string, args interface{}, reply interface{}) bool {
	// c, err := rpc.DialHTTP("tcp", "127.0.0.1"+":1234")
	sockname := coordinatorSock()
	c, err := rpc.DialHTTP("unix", sockname)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer c.Close()

	err = c.Call(rpcname, args, reply)
	if err == nil {
		return true
	}

	fmt.Println(err)
	return false
}
