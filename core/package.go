package main

import (
	"github.com/tsawler/myniceprogram/helpers"
	"log"
)

func main() {
	var myVar helpers.SomeType

	myVar.TypeName = "ppx"
	myVar.TypeNumber = 666

	log.Println(myVar.TypeName, myVar.TypeNumber) // ppx 666
}