/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func Accumulate(val int) func() int {
	return func() int {
		// 这里引用的val是被这个匿名函数引用 所以形成了闭包
		val++
		return val
	}
}

func playGen(name string) func() (string, int) {
	hp := 150
	return func() (s string, i int) {
		s = name
		i = hp
		return
	}
}

// 闭包 -- 引用了外部变量的匿名函数
func main() {
	// Go的匿名函数就是一个闭包，是可以包含自由变量的代码块，这些变量不在这个代码块内
	// 或者任何全局上下文中定义，而是在定义代码块的环境中定义。

	// 1. 在闭包内部修改引用的变量
	str := "Hello World"
	foo := func() {
		str = "Hey"
	}
	foo()
	fmt.Println(str)

	// 2. 闭包的记忆效应
	accumulator := Accumulate(1)
	fmt.Println(accumulator())
	fmt.Println(accumulator())
	fmt.Printf("%p\n", &accumulator)

	accumulator2 := Accumulate(10)
	fmt.Println(accumulator2())
	fmt.Println(accumulator2())
	fmt.Printf("%p\n", &accumulator2)

	fmt.Println(accumulator())

	acc1 := Accumulate(1)
	acc2 := Accumulate(1)
	acc3 := Accumulate(1)
	fmt.Printf("%p %p %p\n", &acc1, &acc2, &acc3)

	// 应用：闭包的记忆效应可被用于设计模式中工厂模式的生成器
	generator := playGen("hi")

	fmt.Println(generator())
}
