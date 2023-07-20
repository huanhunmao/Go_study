package main

import "log"

type myStruct struct {
	FirstName string
}

func (s *myStruct) printFirstName()  string {
	return s.FirstName
}

func main() {
	var myVar myStruct

	myVar.FirstName = "ppx"

	myVar2 := myStruct{
		FirstName: "shy",
	}

	log.Println("myVar is set to", myVar.printFirstName())   // myVar is set to ppx
	log.Println("myVar2 is set to", myVar2.printFirstName()) // myVar2 is set to shy
}
