package main

import "fmt"

// 一个指针 指向 num 指针对应值改变 原值改变

func main() {
	var num int = 42
	var ptr *int // 定义一个指向整数的指针变量
	ptr = &num   // 将 num 的内存地址赋值给指针变量 ptr

	fmt.Println("num:", num)
	fmt.Println("ptr:", ptr)
	fmt.Println("value at ptr:", *ptr) // 通过指针访问 num 的值

	*ptr = 666               // 通过指针修改 num 的值
	fmt.Println("num:", num) // num 的值已被修改
}

// num: 42
// ptr: 0x140000140d0
// value at ptr: 42
// num: 666
