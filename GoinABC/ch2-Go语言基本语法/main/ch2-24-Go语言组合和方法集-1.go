/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

type X struct {
	a int
}

type Y struct {
	// 这里的值和类型都是X 可以省略成为一个X
	// 另外还需要注意的是大写开头是public 小写开头的是private的
	X
	b int
}

type Z struct {
	Y
	c int
}

func (x X) Print() {
	fmt.Printf("In X, a ＝ %d\n", x.a)
}
func (x X) XPrint() {
	fmt.Printf("In X, a ＝ %d\n", x.a)
}
func (y Y) Print() {
	fmt.Printf("In Y, b ＝ %d\n", y.b)
}
func (z Z) Print() {
	fmt.Printf("In Z, c ＝ %d\n", z.c)
	//显式的完全路径调用内嵌字段的方法
	z.Y.Print()
	z.Y.X.Print()
}

func main() {
	// struct 可以嵌入任何其他类型的字段 也可以嵌套自身的指针类型的字段
	// 内存的分配时按照字段的顺序依次开辟连续的存储空间，除了字段对齐之外，不会插入任何额外的东西

	// Go没有继承的语义，没有父子的概念
	x := X{a: 1}
	y := Y{
		X: x,
		b: 2,
	}
	z := Z{
		Y: y,
		c: 3,
	}

	fmt.Println(z.a, z.Y.X.a, z.X.a, z.Y.a)

	// 接下来是函数的调用
	// 原则就是 优先从外向内查找方法，同名方法中外层方法能够覆盖内层的方法
	z.Print()

	z.XPrint()

	z.Y.XPrint()
}
