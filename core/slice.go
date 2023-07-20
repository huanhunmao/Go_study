package main

import (
	"log"
	"sort"
)

func main() {
	var mySlice  []string

	mySlice = append(mySlice, "ppx")
	mySlice = append(mySlice, "shy")
	mySlice = append(mySlice, "fisher")

	log.Println(mySlice) // [ppx shy fisher]  类似js 数组

	sort.Int(mySlice) // 不是 Int类型报错
}