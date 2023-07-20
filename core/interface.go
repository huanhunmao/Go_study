package main

import "fmt"

// 这个接口包含两个方法 
type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name string
	Age int
}

func main() {
	dog := Dog{
		Name: "ppx",
		Age:2,
	}

	

	PrintInfo(dog)
}

func PrintInfo(a Animal){
	fmt.Println("This animal says", a.Says(),"and has", a.NumberOfLegs(),"legs") // This animal says wangwangwang and has 4 legs
}

func(d Dog) Says() string {
	return "wangwangwang"
}

func (d Dog) NumberOfLegs() int {
	return 4
}