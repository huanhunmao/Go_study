package main

import "log"

func main() {
	animals := make(map[string]string)

	animals["cat"] = "xiaomao"
	animals["pig"] = "pgone"
	animals["tigger"] = "xiaohu"

	for key, val := range animals {
		log.Println(key,val) // cat xiaomao    pig pgone            tigger xiaohu
	} 
}