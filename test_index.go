package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "AAABBB"
	fmt.Print(strings.Index(str, "A"), strings.Index(str, "C"))
}
