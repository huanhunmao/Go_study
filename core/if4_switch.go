package main

import "log"

func main() {
	myNum := 666
	isTrue := true

	// if myNum > 666 && !isTrue {
	// 	log.Println("bad input, to big num")
	// }else if myNum == 666 && !isTrue {
	// 	log.Println("good input")
	// }else if myNum == 666 && isTrue {
	// 	log.Println("better input")
	// }else{
	// 	log.Println("bad input")
	// }

	// switch  简化一下这个流程 
	switch myNum {
		case 666:
			log.Println("good input")
		case 66:
			log.Println("bad input")
		default:
			log.Println("bad input")
	}

	switch isTrue {
	case true:
		log.Println("good input")
	case false:
		log.Println("bad input")
	default:
		log.Println("bad input")
}
}