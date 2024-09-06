package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case <-ch:
				fmt.Print("success")
			case <-time.After(3 * time.Second):
				fmt.Println("3秒超时")
				quit <- true
			}
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			ch <- 1
		}
	}()

	<-quit
	fmt.Println("程序结束")
}
