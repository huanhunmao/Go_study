package main

import "log"

func main() {
	var myString string
	myString = "ppx"

	log.Println("myString is set to", myString) // 2023/07/18 17:16:59 myString is set to ppx

	usingPointer(&myString)
	log.Println("myString is set to", myString) // 2023/07/18 17:19:18 myString is set to the shy
}

func usingPointer(s *string) {
	log.Println("s is set to", s) // 2023/07/18 17:19:18 s is set to 0x1400010a230
	// var newValue string
	// newValue = "the shy"
	// or
	newValue := "the shy"
	*s = newValue
}
