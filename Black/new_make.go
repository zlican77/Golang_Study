package main

import "fmt"

func main() {
	type structing struct {
		s string
	}

	var s []int
	var m map[int]string
	var c chan int
	var a *[2]string
	var b *structing

	s = make([]int, 10)
	m = make(map[int]string)
	c = make(chan int)
	a = new([2]string)
	b = new(structing)
	fmt.Println(s, m, c, *a, *b)

	s = append(s, 1)
	m[0] = "ling"
	a[0] = "zi"
	b.s = "chen"
	fmt.Println(s, m, c, *a, *b)

}
