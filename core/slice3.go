package main

import "log"

func main() {
	// 字符串类型
	names := []string{"ppx", "shy", "faker"}

	// 数字类型
	numbers := []int{1, 2, 3, 4, 5}

	log.Println("names", names) // names [ppx shy faker]
	log.Println("numbers", numbers) //  numbers [1 2 3 4 5]

	log.Println(numbers[0:2]) // [1 2]
	log.Println(numbers[3:5]) // [4 5]

	log.Println(numbers[0]) // 1
}
