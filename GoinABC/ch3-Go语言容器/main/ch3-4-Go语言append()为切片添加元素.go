/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {
	// Go 切片的append函数使用
	var sli []int
	// 追加一个元素
	sli = append(sli, 1)
	// 追加多个元素 手写解包方式
	sli = append(sli, 1, 2, 3)
	// 追加多个元素 切片需要解包
	sli = append(sli, []int{1, 2, 3}...)
	fmt.Println(sli)

	// 扩容
	var sli1 []int
	for i := 0; i < 10; i++ {
		sli1 = append(sli1, i)
		fmt.Printf("len: %d, cap: %d, pointer: %p\n", len(sli1), cap(sli1), sli1)
	}

	// 利用append插入元素
	var sli2 = []int{1, 2, 3}
	// 插入头 头插法的性能差 1.内存会重新分配 2.所有已有的元素会复制一遍
	sli2 = append([]int{0}, sli2...)
	fmt.Println(sli2)
	sli2 = append([]int{-3, -2, -1}, sli2...)
	fmt.Println(sli2)

	// 指定位置i插入
	var sli3 = []int{1, 2, 3}
	i := 1
	sli3 = append(sli3[:i], append([]int{4}, sli3[i:]...)...)
	fmt.Println(sli3)

	j := 2
	sli3 = append(sli3[:j], append([]int{5, 6, 7}, sli3[j:]...)...)
	fmt.Println(sli3)
}
