package main

import "log"

func main() {
	// 结构体
	type User struct {
		FirstName string
		LastName string
		Email string
		Age int
	}

	var users []User
	users = append(users,User{"Ppx","shy","Ppx@shy.com",18})
	users = append(users,User{"Edg","omg","Edg@omg.com",29})
	users = append(users,User{"Wbg","plc","Ppx@shy.com",30})


	// or 
// 创建多个 User 结构体的实例(单个创建)
// user1 := User{
// 	FirstName: "John",
// 	LastName:  "Doe",
// 	Email:     "john.doe@example.com",
// 	Age:       30,
// }

// user2 := User{
// 	FirstName: "Jane",
// 	LastName:  "Smith",
// 	Email:     "jane.smith@example.com",
// 	Age:       25,
// }


	log.Println("users", users) //  users [{Ppx shy Ppx@shy.com 18} {Edg omg Edg@omg.com 29} {Wbg plc Ppx@shy.com 30}]
	for _ , l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
// 		2023/07/20 14:13:30 Ppx shy Ppx@shy.com 18
// 2023/07/20 14:13:30 Edg omg Edg@omg.com 29
// 2023/07/20 14:13:30 Wbg plc Ppx@shy.com 30
	}
}