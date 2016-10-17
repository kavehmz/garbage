package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/kavehmz/garbage/012_protobuf/tutorial"
)

func main() {

	p := &tutorial.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*tutorial.Person_PhoneNumber{
			{Number: "555-4321", Type: tutorial.Person_HOME},
		},
	}

	out, err := ioutil.ReadFile("test.out")
	if err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

	fmt.Println(p)

	p2 := &tutorial.Person{}
	if err := proto.Unmarshal(out, p2); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
	p2.Pin = "test"
	fmt.Println(p2)
}
