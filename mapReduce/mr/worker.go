package mr

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"io"
	"log"
	"net/rpc"
	"os"
	"sort"
	"time"
)

// Map functions return a slice of KeyValue.
type KeyValue struct {
	Key   string
	Value string
}

type GetJobArgs struct{}

type GetJobReply struct {
	FileName string
	JobId    int
	NReduce  int
	Type     string
}

type JobDoneArgs struct{}

type JobDoneReply struct{}

// use ihash(key) % NReduce to choose the reduce
// task number for each KeyValue emitted by Map.
func ihash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32() & 0x7fffffff)
}

func handleMapJob(
	fileName string,
	mapf func(string, string) []KeyValue,
	mapTaskId int,
	nReduce int,
) bool {
	inputFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("cannot read %v", fileName)
	}
	defer inputFile.Close()

	content, err := io.ReadAll(inputFile)
	if err != nil {
		log.Fatalf("cannot read %v", fileName)
	}

	mapgoResult := mapf(fileName, string(content))

	enc := make([]*json.Encoder, nReduce)

	for r := range nReduce {
		intermediateFile := fmt.Sprintf("mr-%d-%d", mapTaskId, r)
		f, err := os.Create(intermediateFile)
		if err != nil {
			log.Fatalf("cannot create intermediate file %v", intermediateFile)
		}
		defer f.Close()
		enc[r] = json.NewEncoder(f)
	}

	for _, kv := range mapgoResult {
		reduceTaskId := ihash(kv.Key) % nReduce
		err := enc[reduceTaskId].Encode(&kv)
		if err != nil {
			log.Fatal("cannot encode map result")
		}
	}

	return true
}

func notifyCoordinatorDone() {
	for {
		args := JobDoneArgs{}
		reply := JobDoneReply{}
		ok := call("Coordinator.Done", &args, &reply)
		if !ok {
			time.Sleep(2 * time.Second)
			continue
		}
		break
	}
}

func handleReduceJob(
	reduceTaskId int,
	reducef func(string, []string) string,
	nMap int,
) bool {
	outputFileName := fmt.Sprintf("mr-out-%d", reduceTaskId)
	outputFile, err := os.Create(outputFileName)
	defer outputFile.Close()

	if err != nil {
		log.Fatalf("cannot create outputfile %v", outputFileName)
	}

	var files []*os.File
	for m := 0; m < nMap; m++ {
		fileName := fmt.Sprintf("mr-%d-%d", m, reduceTaskId)
		if f, err := os.Open(fileName); err == nil {
			files = append(files, f)
		}  
	}

	for _, file := range files {
		kva := []KeyValue{}
		var kv KeyValue
		dec := json.NewDecoder(file)
		for {
			if err := dec.Decode(&kv); err != nil {
				break
			}
			kva = append(kva, kv)
		}

		sort.Slice(kva, func(i, j int) bool {
			return kva[i].Key < kva[j].Key
		})

		kvmap := make(map[string][]string)
		for _, kv := range kva {
			kvmap[kv.Key] = append(kvmap[kv.Key], kv.Value)
		}

		for k, v := range kvmap {
			output := reducef(k, v)
			fmt.Fprintf(outputFile, "%v %v\n", k, output)
		}

		file.Close()
	}
	return true
}

func Worker(
	mapf func(string, string) []KeyValue,
	reducef func(string, []string) string,
) {
	for {
		args := GetJobArgs{}
		reply := GetJobReply{}

		if !call("Coordinator.GetJob", &args, &reply) {
			log.Println("Worker: coordinator unavailable, retrying...")
			time.Sleep(2 * time.Second)
			continue
		}

		// Check if there’s actually a job
		if reply.FileName == "" {
			log.Println("Worker: no more jobs, exiting")
			break
		}

		switch reply.Type {
		case "map":
			handleMapJob(reply.FileName, mapf, reply.JobId, reply.NReduce)
		case "reduce":
			handleReduceJob()
		}

		notifyCoordinatorDone()

		time.Sleep(500 * time.Millisecond) // avoid hot-looping
	}
}

// send an RPC request to the coordinator, wait for the response.
// usually returns true.
// returns false if something goes wrong.
func call(rpcname string, args any, reply any) bool {
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
