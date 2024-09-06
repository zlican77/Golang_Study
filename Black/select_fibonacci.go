package main

import "fmt"

func fibonacci(ch chan<- int, quit <-chan bool) {
	x, y := 1, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case flag := <-quit:
			fmt.Println(flag)
			return
		}
	}
}

func main() {
	ch := make(chan int)    //斐波那契数字通道
	quit := make(chan bool) //使用通信分享结束消息

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		quit <- true //main结束任务协程后通过通道传递结束信息
	}()

	fibonacci(ch, quit)
}
