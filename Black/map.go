package main

import "fmt"

func main() {
	map1 := make(map[int]string)

	map2 := map[int]string{
		110: "报警电话",
	}

	fmt.Println(map1)
	fmt.Println(map2)
}
