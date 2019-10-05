/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

type class struct {
}

func (c *class) Do(v int) {
	fmt.Println("call method Do:", v)
}

func funcDo(v int) {
	fmt.Println("call function funcDo:", v)
}

func main() {
	// Go中将类型的方法和普通函数视为同一概念
	var delegate func(int)

	c := new(class)

	delegate = c.Do

	delegate(100)

	delegate = funcDo

	delegate(100)
	// 结论：无论是普通函数还是结构体的方法，只要它们的签名一致，
	// 与它们签名一致的函数变量就可以保存普通函数或是结构体方法。
}
