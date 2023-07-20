package main

import "log"

func main() {
	// 常规定义
	// var myStr string
	// var myInt int

	// myStr = "ppx"
	// myInt = 666

	// myStr2 := "shy"

	// log.Println(myInt, myStr, myStr2) // 666 ppx shy

	// map 怎么写
	myMap := make(map[string]string) // 索引-值 类似js 数组

	myMap["cat"] = "miaomiaomao"

	log.Println(myMap["cat"], myMap["dog"]) // miaomiaomao

	// 尝试修改值 可以被修改
	myMap["cat"] = "ppx"

	log.Println(myMap["cat"], myMap["dog"]) // ppx
	log.Println(myMap) // map[cat:ppx]
}
