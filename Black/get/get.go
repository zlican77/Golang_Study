package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person //继承Person 结构，没有名字匿名字段，包含Person的所有成员
	id     int
	addr   string
}

func (tmp Person) PrintInfo() {
	fmt.Println(tmp.name, tmp.sex, tmp.age)
}

func main() {
	//顺序初始化
	var s1 Student = Student{Person{"mike", 'm', 20}, 1, "合肥"} //属性继承
	s1.PrintInfo()                                             //方法继承
	//指定成员初始化, 其他默认自动为0
	s3 := Student{id: 1}
	fmt.Println(s3)

}
