/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {
	// 创建整型切片的切片
	sli := [][]int{{10, 20, 30, 40, 50}, {10, 20}}
	fmt.Println(sli)
	// 可见slice的首地址和sli的地址不同 且切片的切片的容量初始化为元素的长度
	fmt.Printf("%p %d %p %d\n", sli, cap(sli), sli[0], cap(sli[0]))
	// append
	sli[0] = append(sli[0], 100)
	fmt.Println(sli)
}
