/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {
	// make主要是用来初始化内置的数据结构 主要是map 切片 和channel
	// new是获取指向某个类型的指针
	slice := make([]int, 10, 10)
	mp := make(map[string]int, 10)
	ch := make(chan int, 5)
	fmt.Println(slice, mp, ch)

	i := new(int)

	var v int
	j := &v
	fmt.Printf("%T %T\n", i, j)
	// 上面关于int的操作是等价的 都是创建一个指向int零值的指针

	// make关键字的原理,go是将make关键字的OMAKE节点 根据参数的不同分别转换成为了
	// OMAKESLICE OMAKEMAP OMAKECHAN 三种不同类型的节点 这些节点
	// 最终会调用不同的运行时函数类初始化数据结构

	// new主要是为类型申请一片内存空间，并返回指向这片内存空间的指针
}
