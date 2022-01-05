package main

import (
	"log"
	"net/rpc"
)

type Args struct {
}

func main() {
	// Arguments to send
	var reply int64
	args := Args{}
	// Function name
	name := "TimeServer.GiveServerTime"
	// Dialing to server
	client, err := rpc.DialHTTP("tcp", "localhost"+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Calling to server
	err = client.Call(name, args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	// Getting result
	log.Printf("%d", reply)
}
