package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type args struct {
	A, B int
}

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply int
	done := make(chan *rpc.Call, 1)
	client.Go("Arith.Multiply", &args{4, 2}, &reply, done)

	<-done

	fmt.Println("The reply pointer value has been changed to: ", reply)
}
