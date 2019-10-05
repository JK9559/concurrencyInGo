/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {
	// 一个类型断言表达式的语法为 i.(T) i为一个接口值 T为一个类型名或者类型字面表示。
	// 类型T 可以是任意一个非接口类型 或者任意接口类型

	// i是断言值 T为断言类型 一个断言可能失败 可能成功

	// 对于T是一个非接口类型：
	// 如果断言值i的动态类型存在 并且与T为同一类型 断言成功 否则失败

	// 如果是一个接口类型：
	// 当i的动态类型实现了接口类型T 则成功 否则失败
	var x interface{} = 123

	n, ok := x.(int)
	fmt.Println(n, ok)
	n = x.(int)
	fmt.Println(n)

	a, ok := x.(float32)
	fmt.Println(a, ok)
	// panic
	a = x.(float32)

}
