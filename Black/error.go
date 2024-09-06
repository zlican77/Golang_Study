package main

import (
	"errors"
	"fmt"
)

func myDiv(a, b int) (result int, err error) {
	if b == 0 {
		err = errors.New("b connot be zero")
	} else {
		result = a / b
	}
	return
}

func main() {
	result, err := myDiv(10, 0)
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Print(result)
	}
}
