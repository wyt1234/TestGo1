package main

import "fmt"

type IGreeting interface {
	say()
}

// 别搞混。。
func sayHello(i IGreeting) {
	i.say()
}

type Go struct{}

func (g Go) say() {
	fmt.Println("Hi, I am GO!")
}

type PHP struct{}

func (p PHP) say() {
	fmt.Println("Hi, I am PHP!")
}

func main() {
	golang := Go{}
	php := PHP{}

	sayHello(golang)
	sayHello(php)
}
