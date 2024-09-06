package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for {
			fmt.Println("go")
		}
	}()

	go func() {
		for {
			fmt.Println("routine")
		}
	}()

	for i := 0; i < 3; i++ {
		runtime.Gosched()
		fmt.Println("hello")
	}
}
