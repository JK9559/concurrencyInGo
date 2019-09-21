/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {
	// 切片默认指向一段连续的内存区域 可以是数组 可以是切片本身

	var array = [3]int{1, 2, 3}
	fmt.Println(array[1:2])
	fmt.Printf("%T\n", array[1:2])
	fmt.Println(array[2:len(array)])
	// 取出元素的个数是 结束位置减去开始位置

	// 取指定位置的切片
	var initArr [30]int
	for i := 0; i < len(initArr); i++ {
		initArr[i] = i + 1
	}
	fmt.Println(initArr[10:15])
	fmt.Println(initArr[20:])
	fmt.Println(initArr[:6])

	sliceA := []int{1, 2, 3}
	fmt.Println(sliceA[:])
	fmt.Println(sliceA[0:0])

	// 切片声明
	var sli1 []int
	// 不仅声明了 还分配了空间 所以不是nil了
	var sli2 = []int{}
	fmt.Println(sli1 == nil)
	fmt.Println(sli1)
	fmt.Println(sli2 == nil)
	fmt.Println(sli2)
}
