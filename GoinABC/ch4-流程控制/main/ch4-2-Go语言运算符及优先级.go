/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {
	// ^ 表示按位异或
	var a = (int8)(5)
	var b = (int8)(11)
	fmt.Printf("%b %b %b\n", a, b, a^b)

	// ^ 单目运算时表示取反
	var c = uint8(0)

	fmt.Printf("%x %x %b %b %b\n", c, ^c, c&(^c), c|(^c), c^(^c))

	var d = uint8(0xb)
	fmt.Printf("%b %b\n", d, ^d)

	var e = int8(5)
	fmt.Printf("%08b %08b %d\n", e, ^e, ^e)
	// 5的二进制为 0000 0101
	// 取反为 1111 1010
	// 因为5默认为int 是有符号数 所以符号位取反为负数
	// 所以 1111 1010 取反再加1 为 0000 0110 所以为-0000 0110 也就是-6

	// &^ 按位消除
	fmt.Println(1&^1, 1&^0, 0&^1, 0&^0)
	// 规则 向异的保留左边 相同的清零
	var i = uint8(5)
	var j = uint8(11)
	fmt.Printf("%b %b %b\n", i, j, i&^j)
}
