package main

import "log"

type User struct {
	firstName string
	lastName string
}

func main() {
	myMap := make(map[string]User)

	// 存储任意类型  类似于 js any 不推荐 
	myMap1 := make(map[string]interface{})

	mySelf := User{
		firstName: "Ppx",
		lastName: "Shy",
	}

	myMap["mySelf"] = mySelf
	// myMap.mySelf = mySelf   // 这样写不可以

	log.Printf("mySelf",mySelf) // main.User={Ppx Shy}
	log.Println("mySelf.firstName",mySelf.firstName) // Ppx
	log.Println("mySelf.lastName",mySelf.lastName) // Shy
}