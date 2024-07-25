package main

import "fmt"

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unsupported operation: " + op)
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func sum(args ...int) int {
	s := 0
	for i := range args {
		s += args[i]
	}
	return s
}

func main() {
	fmt.Println(eval(3, 4, "*"))
	q, r := div(12, 5)
	fmt.Println(q, r)
	fmt.Println(sum(1, 3, 4, 5, 6, 8))
}
