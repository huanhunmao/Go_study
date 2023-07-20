package main

import "log"

func main() {
	animals := []string{"cat","dog","fish","pig"}

	for i, animal := range animals {
		log.Println("i: ", i, "animal: ", animal) // i:  0 animal:  cat
	}

	// 不关心索引 i 可以 用 _ 占位 
	for _, animal := range animals {
		log.Println("animal: ", animal) //animal:  cat
	}

	log.Println(animals) // [cat dog fish pig]
}