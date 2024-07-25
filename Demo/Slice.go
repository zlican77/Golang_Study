package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s := arr[2:7]
	fmt.Println("arr[2:7] = ", s)

	fmt.Println("Traversing slice")

	for i, v := range arr {
		fmt.Println(i, v)
	}
}
