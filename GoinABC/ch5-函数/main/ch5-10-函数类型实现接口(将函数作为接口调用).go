/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

// 其他类型能实现接口 函数也可以 以下代码对比了函数 和 结构体实现的接口
import "fmt"

// 定义了一个接口
type Invoker interface {
	// 实现该接口需要实现一个Call方法
	Call(interface{})
}

// 定义了一个结构体类型
type Struct struct {
}

// 在结构体内定义了一个函数Call
func (s *Struct) Call(p interface{}) {
	fmt.Println("from struct", p)
}

// 定义了一个函数类型--本节重点 函数类型实现接口
type FuncCaller func(interface{})

// 实现了Invoker的Call
func (f FuncCaller) Call(p interface{}) {
	// 调用f函数
	f(p)
}

func main() {
	// 声明了一个接口变量
	var invoker Invoker

	// 创建了一个Struct
	s := new(Struct)

	// 因为Struct实现了接口Invoker 所以可以赋值
	invoker = s

	// 调用结构体的Call函数
	invoker.Call("Hello")

	// 首先将匿名函数转换为FuncCaller类型 然后赋值给接口
	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("from function", v)
	})

	// 由接口调用函数的Call函数
	invoker.Call("Hello")
}
