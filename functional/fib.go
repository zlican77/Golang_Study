package main

import "fmt"

func fibonacci() func() int {
	a, b := 1, 0
	base := 0
	return func() int {
		base = a + b
		a, b = b, a
		b = base
		return base
	}
}

func main() {
	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
