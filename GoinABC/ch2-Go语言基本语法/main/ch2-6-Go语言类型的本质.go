/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {
	// 1. 内置类型：数值类型 字符串类型 布尔类型
	var a int = 10
	var b string = "abc"
	var c bool = true
	changeFun := func(a int, b string, c bool) (int, string, bool) {
		a = 6
		b = "hah"
		c = false
		return a, b, c
	}
	d, e, f := changeFun(a, b, c)
	fmt.Println(a, b, c, d, e, f)
	// 说明对于内置类型 函数是值传递 传递的是一个对应值的副本

	// 2. 引用类型：切片 映射 接口 函数类型
	var slic = make([]int, 4)
	slic = []int{1, 2, 3, 4}
	changeFunAgain := func(s []int) []int {
		s[0] = 99
		return s
	}
	s := changeFunAgain(slic)
	fmt.Println(s)
	fmt.Println(slic)

	var mp = make(map[string]string, 5)
	mp["a"] = "a"
	mp["b"] = "b"
	mp["c"] = "c"
	fmt.Println(mp)
	mp1 := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println(mp1)
	changeFunMap := func(m map[string]string) map[string]string {
		m["a"] = "bbb"
		m["ss"] = "cc"
		return m
	}
	m := changeFunMap(mp)
	fmt.Println(m)
	fmt.Println(mp)
}
