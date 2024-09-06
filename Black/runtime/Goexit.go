package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("ccc")

	runtime.Goexit()

	fmt.Println("ddd")
}

func main() {
	go func() {
		fmt.Println("aaa")

		test()

		fmt.Println("bbb")
	}()

	for {

	}

}
