package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName   string
	LastName    string
	HasChildren bool
}

func main() {
	var mySlice []Person
	var person1 Person

	person1.FirstName = "PPX"
	person1.LastName = "Vc"
	person1.HasChildren = true

	mySlice = append(mySlice, person1)

	var person2 Person

	person2.FirstName = "Shy"
	person2.LastName = "PP"
	person2.HasChildren = false


	mySlice = append(mySlice, person2)

	newJson, err := json.MarshalIndent(mySlice, "", "    ") // 数据 前缀 缩进

	if err != nil {
		log.Println("error marshalling", err)
	}

	fmt.Println(string(newJson))
}
