package main

import "log"

func main() {
	var isTrue bool

	isTrue = false

	if isTrue { // 条件可以省略括号 
		log.Println("is True", isTrue)
	} else {
		log.Println("is True", isTrue) // is True false
	}
}
