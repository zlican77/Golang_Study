package main

import "fmt"

type Animal interface {
	Humaner //继承接口
}

type Humaner interface {
	eat()
}

type Student struct {
	name string
	age  int
}

func (s *Student) eat() {
	fmt.Printf("%v eating\n", *s)
}

type Teacher struct {
	name string
	age  int
	sex  byte
}

func (t *Teacher) eat() {
	fmt.Printf("%v eating\n", *t)
}

type Animals struct {
	types string
}

func (a *Animals) eat() {
	fmt.Printf("%v eating\n", *a)
}

func main() {
	s := &Student{"zl", 20}
	t := &Teacher{"hf", 30, 'm'}
	a := &Animals{"cat"}

	var h Animal
	array := [3]Animal{s, t, a}

	for _, data := range array {
		h = data
		h.eat()
	}
}
