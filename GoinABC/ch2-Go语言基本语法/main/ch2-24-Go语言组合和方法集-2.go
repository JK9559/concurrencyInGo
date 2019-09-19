/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

// 若类型S包含匿名字段T 则S的方法集包含T的方法集
// 若类型S包含匿名字段*T 则S的方法集包含T和*T的方法集
// 不管S中嵌入的匿名字段是T还是*T *S方法集总包含了T和*T的方法集
type A struct {
	a int
}

type B struct {
	A
}

type C struct {
	*A
}

func (a A) Get() int {
	return a.a
}

func (a *A) Set(i int) {
	a.a = i
}

func main() {
	a := A{a: 1}
	b := B{A: a}
	fmt.Println(b.Get())

	// b.Set(2) 也可以用b是因为编译器做了自动转化
	(&b).Set(2)
	fmt.Println(b.Get())
	fmt.Println(B.Get(b))

	(*B).Set(&b, 3)
	fmt.Println(b.Get())

	c := C{
		A: &a,
	}
	C.Set(c, 4)
	fmt.Println(c.Get())

	(*C).Set(&c, 5)
	fmt.Println(c.Get())
}
