// package main

// import (
// 	"log"

// 	"github.com/tsawler/myniceprogram/helpers"
// )

// const numPool = 100 // 数字池 多大范围 可控的

// func CalculateValue(inChan chan int) { // 输入 int 类型的通道 inChan chan int
// 	randomNumber := helpers.RandomNumber(numPool)

// 	inChan <- randomNumber // 将生成的随机数  给通道
// }

// func main() {
// 	intChan := make(chan int) // 通道定义
// 	defer close(intChan)      // 使用完成后关闭通道  good  practice

// 	go CalculateValue(intChan) // 将通道结果 给一个变量 并打印出变量
// 	num := <-intChan

// 	log.Println("num", num)
// }

package main

import (
	"github.com/tsawler/myniceprogram/helpers"
	"log"
)

const numberPool = 100

func CalculateValue(intChan chan int){
	value  := helpers.RandomNumber(numberPool)

	intChan <- value
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	number := <- intChan

	log.Println("number: ", number)
}