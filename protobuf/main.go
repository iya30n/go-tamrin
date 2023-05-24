package main

import (
	"fmt"
	"go-tamrin/grpc/api/person"
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/proto"
)

func writeMessage(p *person.Person) {
	data, err := proto.Marshal(p)
	if err != nil {
		log.Fatalln(err)
	}

	if err := ioutil.WriteFile("person", data, 0644); err != nil {
		log.Fatalln("Failed to write person:", err)
	}
}

func readMessage(fileName string) *person.Person {
	input, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}

	person := &person.Person{}
	if err := proto.Unmarshal(input, person); err != nil {
		log.Fatalln("Failed to parse person:", err)
	}

	return person
}

func main() {
	p := person.Person{
		Id: 1,
		Name: "john person",
		Email: "person@gmail.com",
		Phones: []*person.Person_PhoneNumber{
			{Number: "021-2121221", Type: person.Person_HOME},
		},
	}

	writeMessage(&p)

	readPerson := readMessage("person")

	fmt.Println(readPerson.Name)
}