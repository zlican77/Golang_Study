package main

import (
	"fmt"
	"time"
)

var c = make(chan bool)

func PrintBot(s string) {
	for _, data := range s {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
}

func person1() {
	PrintBot("Hello")
	c <- true
}

func person2() {
	<-c
	PrintBot("World")
}

func main() {
	go person1()
	go person2()

	for {

	}
}
