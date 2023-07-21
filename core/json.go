package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	HasChildren bool   `json:"has_children"`
}

func main() {
	// json 每一项 末尾项 不能加 逗号
	myJson := `
	[
		{
		"first_name": "John",
		"last_name": "Vc",
		"has_children": true
	},
	{
		"first_name": "Pika",
		"last_name": "Pc",
		"has_children": false
	}
	]
`

	var value []Person

	err := json.Unmarshal([]byte(myJson), &value)

	if err != nil {
		log.Println("error unmarshalling json: ", err)
	}

	log.Println("unmarshalling: %v", value) // unmarshalling: %v [{John Vc true} {Pika Pc false}]
}
