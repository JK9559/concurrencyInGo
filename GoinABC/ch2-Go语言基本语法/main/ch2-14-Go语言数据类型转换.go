/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	// Go语言不存在隐式的类型转换 所以所有的转换都必须是显性的
	// B的值 = B的类型(A的值)
	a := 5.0
	var b int
	b = int(a)
	fmt.Println(b)

	// 当取值范围较小的类型往取值范围较大的类型转换的时候 没问题
	// 反之 则会出现精度丢失的情况
	fmt.Println("int8 range:", math.MinInt8, math.MaxInt8)
	fmt.Println("int16 range:", math.MinInt16, math.MaxInt16)
	fmt.Println("int32 range:", math.MinInt32, math.MaxInt32)
	fmt.Println("int64 range:", math.MinInt64, math.MaxInt64)

	var ai int32 = 1047483647
	fmt.Printf("int32: 0x%x %d\n", ai, ai)

	bi := int16(ai)
	fmt.Printf("int16: 0x%x %d\n", bi, bi)

	var ci float32 = math.Pi
	fmt.Println(int(ci))
	// 观察上面的运行结果 可以看到当高精度向低精度转换的时候
	// 高精度数字被截断，结果可以看到int32转换为int16时截断了16bit对应十六进制为4位
	// 浮点数转换为整型的时候 小数部分被截断
}
