/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {
	// 多维数组
	// 从左往右看 是数组长度为4 每个元素为[2]int 类型
	var array [4][2]int
	array = [4][2]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	fmt.Println(array)
	array = [4][2]int{1: {9, 10}, 3: {11, 12}}
	fmt.Println(array)
	array = [4][2]int{1: {0: 20}, 0: {1: 21}}
	fmt.Println(array)

}
