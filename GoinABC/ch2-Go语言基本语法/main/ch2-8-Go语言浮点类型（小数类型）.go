/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	// Go 的浮点型 提供了两种精度 float32 float64
	// float32 取值范围 最小为1.4e-45 最大为3.4e38
	fmt.Println(math.MaxFloat32)
	// float64 取值范围 最小为4.9e-324 最大为1.8e308
	fmt.Println(math.MaxFloat64)

	// 通常应该优先使用float64 因为float32的累计计算误差容易扩散，且能精确表示的正整数不大
	var f float32 = 1 << 24
	// 注意 此时下面输出为true 因为float32的有效bit位只有23个
	fmt.Println(f == f+1)

	// print浮点类型可以使用 %g %f %e
	for i := 0; i < 8; i++ {
		f := math.Exp(float64(i))
		fmt.Printf("i = %d e^x = %8.3f e^x = %8.3g e^x = %8.3e\n", i, f, f, f)
	}

	var z float64
	// 0 -0 +Inf -Inf NaN
	fmt.Println(z, -z, 1/z, -1/z, z/z)

	fmt.Println((1 / z) == (1 / z))

	// 函数math.IsNaN 用来判断一个数字是否为非数NaN
	// 浮点数中 Nan 正/负无穷大都不是唯一的
	nan := math.NaN()
	fmt.Println(math.IsNaN(nan), nan == nan, nan < nan, nan > nan)
}
