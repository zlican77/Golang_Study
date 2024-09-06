package main

import "fmt"

type Student struct {
	name string
	id   int
}

func main() {
	i := make([]interface{}, 3)
	i[0] = 1
	i[1] = "hello golang"
	i[2] = Student{"mike", 8}
	//if 断言
	for index, data := range i {
		if value, ok := data.(int); ok {
			fmt.Printf("x[%d] 类型为int， 内容为%d\n", index, value)
		} else if value, ok := data.(string); ok {
			fmt.Printf("x[%d] 类型为string， 内容为%s\n", index, value)
		} else if value, ok := data.(Student); ok {
			fmt.Printf("x[%d] 类型为Student， 内容为%+v\n", index, value)
		}
	}

	//switch 断言
	for index, data := range i {
		switch data.(type) {
		case int:
			fmt.Printf("x[%d] 类型为int， 内容为%d\n", index, data.(int))
		case string:
			fmt.Printf("x[%d] 类型为string， 内容为%s\n", index, data.(string))
		case Student:
			fmt.Printf("x[%d] 类型为Student， 内容为%+v\n", index, data.(Student))
		}
	}
}
