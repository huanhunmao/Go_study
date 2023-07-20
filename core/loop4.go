package main

import (
	"fmt"
	"log"
	"unicode/utf8"
)

func main() {
	myStr := "I have a apple "

	for _, line := range myStr{
		log.Println(line) // 73 32 104 ... 一堆数字  go 中字符串是 字节形式 
		// char, size := utf8.DecodeRuneInString(string(line))
		// fmt.Println(char, size)
	}
}