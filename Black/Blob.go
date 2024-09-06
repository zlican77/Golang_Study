package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubbleSort(list []int) {
	n := len(list)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	arr1 := make([]int, 100)
	for i := range arr1 {
		arr1[i] = rand.Intn(100)

	}

	bubbleSort(arr1)
	fmt.Println(arr1) // 输出: [4 5 6 7 10 18 31 42]

	fmt.Printf("the slice type is %T\n", arr1)
}
