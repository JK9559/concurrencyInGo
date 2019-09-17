/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {
	// Go中没有枚举 可以配合const使用iota来实现模拟枚举
	// 基础版
	type TYPE_A int
	const (
		a TYPE_A = iota
		b
		c
	)
	var d TYPE_A = c
	fmt.Println(a, b, c, d)

	// 拓展版
	type TYPE_B int
	const (
		aa TYPE_B = 1 << iota
		bb
		cc
	)
	var dd TYPE_B = cc
	fmt.Println(aa, bb, cc, dd)
	// %b 表示二进制输出
	fmt.Printf("%b %b %b %b\n", aa, bb, cc, dd)
}
