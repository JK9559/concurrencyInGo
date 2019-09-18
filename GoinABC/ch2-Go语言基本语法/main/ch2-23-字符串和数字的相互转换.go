/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 整数转换为字符串，有两种选择
	// 1. fmt.Sprintf
	x := 123
	fmt.Printf("%T\n", x)
	y := fmt.Sprintf("%d", x)
	fmt.Printf("%T %s\n", y, y)

	// 2. strconv.itoa
	z := strconv.Itoa(x)
	fmt.Printf("%T %s\n", z, z)

	// 按照不同的进制格式化数字
	h := 32
	fmt.Printf("%b %o %d %x\n", h, h, h, h)

	// 字符串转整数
	// strconv.Atoi ParseInt
	i, _ := strconv.Atoi("123")
	fmt.Printf("%T %d\n", i, i)

	j, _ := strconv.ParseInt("1000", 2, 64)
	fmt.Printf("%T %d\n", j, j)

	k, _ := strconv.ParseInt("32", 10, 16)
	fmt.Printf("%T %d\n", k, k)

	l, _ := strconv.ParseInt("32", 10, 0)
	fmt.Printf("%T %d\n", l, l)

}
