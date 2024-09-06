package main

import (
	"fmt"
	"math/rand"
	"time"
)

//1. 设计创建随机4位数 数字
//2. 将该数字转化为切片
//3. 设计游戏

func createNum() int {
	rand.Seed(time.Now().UnixNano())
	var num int
	for {
		num = rand.Intn(10000)
		if num > 999 && num < 10000 {
			break
		}
	}
	return num
}

func main() {
	var num = createNum()
	fmt.Println(num)
}
