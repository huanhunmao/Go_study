package main

import "log"

func main() {
	myMap := make(map[string]int)

	myMap["first"] = 1 * 1
	myMap["last"] = 9 * 9

	log.Println("myMap:", myMap) // myMap: map[first:1 last:81]
	log.Println("myMap['first']:", myMap["first"]) // myMap['first']: 1
	log.Println("myMap['last']:", myMap["last"])  //  myMap['last']: 81
} 
