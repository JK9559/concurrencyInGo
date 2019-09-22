/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {
	// 可以利用range来对slice进行迭代
	sli := []int{1, 2, 3, 4, 5}
	for k, v := range sli {
		fmt.Printf("%d %d\n", k, v)
	}

	// 重点在于 range迭代slice的时候 是对slice进行了值引用(也就是复制了每个元素的副本)
	sli1 := []int{1, 2, 3, 4, 5}
	for k, v := range sli1 {
		fmt.Printf("val: %d, val`s p: %p, sli`s p: %p\n", v, &v, &sli1[k])
	}
	// 可以看到v的地址和实际slice元素的地址是不同的
	// 想要获取每个元素的实际地址 可以使用切片变量和索引值
}
