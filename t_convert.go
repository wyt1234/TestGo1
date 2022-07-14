package main

import "fmt"

// 一、隐式类型转换和强转
func test1() {
	var a float32 = 5.6
	var b int = 10
	//fmt.Println (a * b) // -- 隐式类型转换，编译报错，不支持
	fmt.Println(a * float32(b)) // -- 强转
}

type IBase interface {
	hello()
}

type IParent interface {
	hi()
}
type Son1 struct {
}

func (son1 *Son1) hi() {
	fmt.Println("son1: hi")
}

func (son1 *Son1) hello() {
	fmt.Println("son1: hello")
}

type Son2 struct {
}

func (son2 Son2) hi() {
	fmt.Println("son2: hi")
}

type grandson struct {
	Son1
}

func (grandson *grandson) hi() {
	fmt.Println("grandson: hi")
}

// 二、类型断言。注意由指针和非指针实现的方法，断言时的写法不同
// TODO .(T) 用来类型断言，返回参数1为断言之后的类型值，如果失败则是nil，参数2为是否断言成功
// 		如果类型本身就是断言的类型，则断言成功，会转换成这个类型并返回
// 可以断言的情况：
// 1.由接口断言到结构体
// 2.由父接口断言到子接口
func test2() {
	// 1.由接口断言到结构体
	var p1 IParent = &Son1{} // 指针实现的方法hi
	p1.(*Son1).hi()

	var p2 IParent = Son2{} // 非指针实现的方法hi
	p2.(Son2).hi()

	// 2.由父接口断言到子接口
	var iBase IBase = &Son1{}
	iBase.(IParent).hi()
}

// 三、“向上造型”(java中这么叫，即转为父类)
func test3() {
	grandson := grandson{}
	grandson.hi()
	// 因为golang中继承的语义是将父结构嵌入(即匿名字段)到子结构，所以只需要调用嵌入的父结构体即可
	grandson.Son1.hi()
}
func main() {
	//t1()
	test2()
	test3()
}
