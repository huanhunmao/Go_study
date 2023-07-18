package main

import "log"

// var UserName string
// var Age string
// var Email string
// var Password string
// var birthday string
// var jobName string

type User struct {
	UserName string
	Age      string
	Email    string
}

func main() {
	user := User{
		UserName: "ppx",
		Age:      "18",
		Email:    "11qq.com",
	}

	log.Println("user:", user)          // user: {ppx 18 11qq.com}
	log.Println("user:", user.UserName) // user: ppx
	log.Println("user:", user.Age)      // user: 18

	user.Age = "28"
	log.Println("user:", user.Age) // user: 28
}
