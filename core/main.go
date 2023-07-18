package main

import "fmt"

func main() {
	fmt.Println("Hello world") // Hello world

	var testStr string
	var i int

	testStr = "Good string"
	i = 666

	fmt.Println(testStr)             // Good string
	fmt.Println("i is steed at ", i) // i is steed at  666

	// var ppx string
	ppx := returnSomething()
	fmt.Println(ppx) // something

	ppx, otherPPX := returnSomething2()
	fmt.Println(ppx, otherPPX) // something else
}

// 自定义函数
func returnSomething() string {
	return "something"
}

func returnSomething2() (string, string) {
	return "something", "else"
}
