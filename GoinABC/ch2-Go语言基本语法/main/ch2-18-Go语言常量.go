/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Go中的常量使用关键字const定义,用于存储不会被改变的数据。
	// 常量是在编译时被创建,即使被定义成为函数的局部变量
	// 只能是布尔型，整数型，浮点型，复数和字符串类型
	const (
		e   = 1.223
		str = "abc"
		bo  = true
		a   = 1
		b
		c
		d = 2
		f = 10
		g
	)

	// 或者是调用一下函数 都是返回常量结果 len cap real imag complex unsafe.Sizeof
	// 常量间的所有算数运算、逻辑运算和比较运算的结果都是常量，对常量的类型转换操作
	arr := [g]int{1, 2, 3, 4, 5}
	fmt.Println(cap(arr))

	fmt.Println(len(str))

	var x complex128 = complex(1, 2)
	fmt.Println(real(x))
	fmt.Println(imag(x))

	// 对于字符串返回的始终是16 与字符串的长度无关 因为组成字符串的结构是
	// 指向字符串地址的指针8字节 另一部分是字符串的长度8字节
	fmt.Println(unsafe.Sizeof(str))

	var s1 = []string{}
	var s2 = []int{1, 2, 3}

	// 对于切片常为24 因为它是由三部分组成 1.指向数组的指针8 2.长度8 3.容量8
	// 对于切片来说 长度是切片引用的元素数 容量则是底层数组中元素的数量
	fmt.Println(unsafe.Sizeof(s1), unsafe.Sizeof(s2))

	type Weekday int
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}
