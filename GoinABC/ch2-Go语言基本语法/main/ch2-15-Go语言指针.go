/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {
	// 指针概念在Go中被拆分成为两个核心概念：
	// 1. 类型指针 允许对这个指针类型的数据进行修改，使用指针来传递数据，不需要拷贝数据。类型指针不能偏移和运算
	// 2. 切片，由指向起始元素的原始指针、元素数量和容量组成

	// 关于指针的几个概念：
	// 一个指针变量可以指向任何一个值的内存地址(该内存地址在32位机器上占4字节，64位机器上占8字节，与所指向的值大小无关)。
	// 尽管如此，也可以声明指针指向的是什么类型的值，来表明他的结构性。加上前缀*可以获取指针指向的内容

	// 指针被定义后未被分配给任何变量时，值为nil，一个指针变量缩写为ptr
	var a int = 1
	var b string = "abc"
	fmt.Printf("%p %p\n", &a, &b)
	// 上面代码输出了两个变量的地址 (十六进制格式)
	// 在32位平台上是32位地址，在64位平台上是64位地址
	// 每个变量都有地址 地址就是指针的值

	var bala = "Yes, it is a string"
	ptr := &bala

	// 打印一个变量的类型 用%T
	fmt.Printf("ptr type: %T\n", ptr)
	// 打印指针地址
	fmt.Printf("address: %p\n", ptr)

	value := *ptr

	fmt.Printf("value type: %T\n", value)
	fmt.Printf("value: %s\n", value)

	swap1 := func(x, y int) {
		a := x
		x = y
		y = a
	}

	swap2 := func(x, y *int) {
		// 取x指向的变量的值
		p := *x
		// 取y指向的变量的值 赋值给x指向的变量
		*x = *y
		*y = p
	}

	var x, y int = 1, 2

	fmt.Println(x, y)
	swap1(x, y)
	fmt.Println(x, y)
	swap2(&x, &y)
	fmt.Println(x, y)
	// 总结
	// 当*操作符为右值的时候 表示取指针的值(取指向变量的值)
	// 当*操作符为左值的时候 表示指针指向的变量(将值设置给指向的变量)

	// 创建指针 也可以通过new(type)来创建
	str := new(string)
	fmt.Printf("%T\n", str)

	*str = "haha"
	fmt.Printf("%T %s\n", *str, *str)
}
