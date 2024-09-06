package main

import (
	"fmt"
	"time"
)

func routine() {
	for {
		fmt.Println("goroutine")
		time.Sleep(time.Second)
	}
}

func main() {
	go routine()

	for {
		fmt.Println("main")
		time.Sleep(time.Second)
	}
}
