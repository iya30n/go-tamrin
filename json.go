package main

import (
	"encoding/json"
	"fmt"
)

type MyJson struct {
	User `json:"user"`
}

type User struct {
	Name string `json:"name"`
	Age uint `json:"age"`
}

func main() {
	myJson := []byte(`{"user": {"name": "sina", "age": 12}}`)

	u := MyJson{}
	err := json.Unmarshal(myJson, &u)
	if err != nil {
		panic(err)
	}

	// fmt.Println(u.User.Name)

	fmt.Println(u)
}
