/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {
	// 1. 所有内存在Go中都是经过初始化的 int 0 float 0.0 bool false string 空字符串 指针 nil
	var a int
	fmt.Println(a)
	var b float64
	fmt.Println(b)
	var c bool
	fmt.Println(c)
	var d string
	fmt.Println(d)
	var e *int
	fmt.Println(e)

	// 2. byte基本类型为uint8的别名 rune基本类型为int32的别名
	// 3. 批量声明变量
	var (
		f int
		g float64
	)
	fmt.Println(f, g)
}
