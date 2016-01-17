package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	person := &Person{
		Name:   "toto",
		Id:     1,
		Email:  "toto@toto.toto",
		Phones: []*Person_PhoneNumber{{"0606060606", Person_MOBILE}},
	}
	fmt.Println(person)
	data, err := proto.Marshal(person)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	person.Reset()
	fmt.Println("Data", data)
	if err = proto.Unmarshal(data, person); err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	fmt.Println(person)
}
