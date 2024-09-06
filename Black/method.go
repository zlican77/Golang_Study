package main

import "fmt"

type long int

//自定义类型封装方法
func (l long) Add(x long) long { //为某个类型绑定函数，即叫做方法。
	//总共两个参数，l为自身参数，x为变化参数
	return l + x
}

//结构体封装方法
type Student struct {
	name string
	sex  byte
}

func (tmp Student) PrintInfo() {
	fmt.Println(tmp)
}
func (tmp *Student) SetInfo(name string, sex byte) {
	tmp.name = name
	tmp.sex = sex
}

func main() {
	var a long
	a = 5

	fmt.Println(a.Add(4))

	s := Student{name: "zz", sex: 'm'}
	s.PrintInfo()

	var s3 = new(Student)
	s3.SetInfo("s3", 'w')
	fmt.Println(*s3)

	var s4 Student
	s4.SetInfo("s4", 'm')
	fmt.Println(s4)
}
