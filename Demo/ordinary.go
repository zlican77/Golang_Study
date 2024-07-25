package main

import "fmt"

var str1 = "abcabcbb"
var str2 = "aavcdefbjggggkjhgfdsalm"
var str3 = "pwwkew"

func oper(str string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0

	for i, ch := range []byte(str) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(oper(str1))
	fmt.Println(oper(str2))
	fmt.Println(oper(str3))
}
