package main

import (
	"fmt"

	pb "carrenolg.io/books/gorest/src/ch6/01-proto/protofiles"
	"google.golang.org/protobuf/proto"
)

func main() {
	p := &pb.Person{
		Id:    1234,
		Name:  "Roger F",
		Email: "rf@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}

	p1 := &pb.Person{}
	body, _ := proto.Marshal(p)
	_ = proto.Unmarshal(body, p1)
	fmt.Println("Original struct loaded from proto file:", p)
	fmt.Println("########")
	fmt.Println("Marshaled proto data:", body)
	fmt.Println("########")
	fmt.Println("Unmarshaled struct:", p1)
}
