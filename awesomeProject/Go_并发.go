package main

import (
	"fmt"
	"time"
)

// Go语言支持并发,只需要用go关键字来开启 goroutine即可
// goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")

	// 测试通道
	channel_test()

	// 测试缓冲通道
	channel_test_buffer()

	// Go遍历通道里的值
	channel_test_slice()
}


// ------------------ channel -----------------------
// 通道（channel）是用来传递数据的一个数据结构。
// 通道可用于两个 goroutine 之间传递一个指定类型的值来同步运行和通讯. 操作符 <- 用于指定通道的方向,发送或接收,如果未指定方向,则为双向通道.


func channel_test() {
	// 声明一个通道
	var channel1 = make(chan int) // 默认情况下通道是不带缓冲区的,发送端发送数据,同时必须又接受端相应的接收数据
	go sum(1,channel1)
	go sum(2,channel1)
	go sum(3,channel1)
	
	// 从通道中取值
	x :=  <- channel1
	y := <- channel1
	z := <- channel1
	fmt.Println("x:",x,"y:",y,"z:",z)
	fmt.Println(channel1,channel1,channel1,channel1,channel1)
}

func sum(s int,c chan int) {
	var result = 0
	result += s
	c <- result // 将值发送到通道里面
}

// 通道缓冲区
func channel_test_buffer() {
	// 定义一个可以存储整数类型的带缓冲通道
	channel1 := make(chan int,2)
	channel1 <- 1
	channel1 <- 2

	// 获取这两个数据
	fmt.Println(<-channel1)
	fmt.Println(<-channel1)
}

// Go遍历通道
func channel_test_slice() {
	channel1 := make(chan int, 2)
	fibonacci(2, channel1)
	for i := range channel1 {
		fmt.Println("for i: ",i)
	}
}

func fibonacci(x int ,c chan int) {
	for i := 0; i < x; i++ {
		c <- i + x
	}
	close(c)
}
