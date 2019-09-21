/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {
	sli1 := []int{1, 2, 3, 4, 5}
	sli2 := []int{6, 7, 8}
	// 将sli1的前3个元素复制到sli2
	ret1 := copy(sli2, sli1)
	fmt.Println(ret1)
	fmt.Println(sli1)
	fmt.Println(sli2)

	sli1 = []int{1, 2, 3, 4, 5}
	sli2 = []int{6, 7, 8}
	ret2 := copy(sli1, sli2)
	fmt.Println(ret2)
	// 将sli2的全部元素复制到sli1的前3个位置
	fmt.Println(sli1)
	fmt.Println(sli2)
	// 两个slice会共享同一个底层数组
}
