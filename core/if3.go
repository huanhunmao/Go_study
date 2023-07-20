package main

import "log"

func main() {
	myNum := 666
	isTrue := false

	if myNum > 666 && !isTrue {
		log.Println("bad input, to big num")
	}else if myNum == 666 && !isTrue {
		log.Println("good input")
	}else if myNum == 666 && isTrue {
		log.Println("better input")
	}else{
		log.Println("bad input")
	}
}