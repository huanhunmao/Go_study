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

type Tiger struct{
	Name string
	Color string
	Abilities string
}

func main() {

	tiger := Tiger{
		Name: "xiaohu",
		Color: "yellow",
		Abilities: "eat meat",
	}

	PrintInfo(&tiger)
}

func PrintInfo(a Animal){
	fmt.Println("This animal says", a.Says(),"and has", a.NumberOfLegs(),"legs") // This animal says AOaoao and has 4 legs
}

func (t *Tiger) Says() string {
	return "AOaoao"
}

func (t *Tiger) NumberOfLegs() int{
	return 4
}