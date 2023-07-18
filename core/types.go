package main

import "log"

var s = "777"

func main() {
	var s2 = "666"

	log.Println("s  is", s)   // s  is 777
	log.Println("s2  is", s2) // s2  is 666

	saySomething("xxx")
}

func saySomething(s3 string) string {
	log.Println("s3 is", s3) // s is xxx  传递的参数
	log.Println("s is", s)   // s is 777  函数外层的全局变量 函数内都可访问到
	return s
}
