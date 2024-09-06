package main

import "fmt"

func printFunc(text string) bool {
	fmt.Println(text)
	return true
}

type onlyFunc func(string) bool

func usePrint(f onlyFunc) {
	f("we will win")
}

func main() {
	usePrint(printFunc)
}
