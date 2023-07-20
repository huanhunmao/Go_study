package main

import "fmt"

func sender(ch chan <- string){
	ch <- "Hello from sender!" // 发送数据到 channel  ch
}

func receiver(ch <-  chan string){
	message := <- ch  // 从 channel 获取数据 给 message
	fmt.Println("Receiver", message)
}

func main() {
	// 创建一个 通道 
	ch := make(chan string)

	// 启动 sender
	go sender(ch)

	// 启动 receiver
	go receiver(ch)

	// 主协程等待一段时间，以允许其他协程完成
	// 这里只是简单的等待，实际应用中可以使用 sync.WaitGroup 或其他同步方法
	// 来确保所有协程都完成了任务
	// 这段时间可以不用等待，因为 sender 和 receiver 都是异步进行的
	// 但为了演示同步过程，我们在这里等待了一段时间
	// 请注意，在实际代码中，您可能不需要这样的等待操作
	fmt.Println("Waiting for other goroutines to finish...")
	fmt.Scanln() // 等待用户输入

	fmt.Println("Main goroutine exited.")

// 	Waiting for other goroutines to finish...
// Receiver Hello from sender!
}