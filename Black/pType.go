package main

import "fmt"

func main() {
	a := 10

	p := &a

	fmt.Printf("p type = %T\n", p)

	var p2 *int
	p2 = new(int)
	data := *p2
	fmt.Println(data)
}
