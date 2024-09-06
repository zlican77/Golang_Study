package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3, 4, 5, 6, 7, 8}

	slice2 := slice[:4]

	fmt.Printf("slice2, len = %d, cap = %d\n", len(slice2), cap(slice2))
	for i := range slice2 {
		fmt.Println(slice2[i])
	}
}
