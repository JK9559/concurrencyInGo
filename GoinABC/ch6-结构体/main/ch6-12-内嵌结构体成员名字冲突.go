/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

type A struct {
	a int
}

type B struct {
	a int
}

type C struct {
	A
	B
}

// 内嵌结构体成员名字冲突
func main() {
	c := &C{}
	c.A.a = 1
	// 当赋值时 会编译错误
	// c.a = 2
}
