package main

import (
	"fmt"
	"vis"
)

func main() {
	s := new(vis.Student)
	fmt.Println(*s)
	s.Id = "A"
	fmt.Println(s.Id)
	fmt.Println(*s)
}
