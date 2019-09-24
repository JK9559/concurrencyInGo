/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

type AA struct {
	a int
	b int
}

func main() {
	// Go map声明
	// var map1 map[key]value
	var map1 map[string]int
	fmt.Println(map1 == nil)
	// 1. 声明时无需知道map的长度
	// 2. key是任意可以用==或=!操作符比较的类型，string int float array不可用 slice struct
	var map2 = make(map[[3]int]int)
	map2[[3]int{1, 2, 3}] = 5
	fmt.Println(map2)

	var map3 = make(map[AA]int)
	map3[AA{
		a: 1,
		b: 3,
	}] = 6
	map3[AA{
		a: 1,
		b: 4,
	}] = 7
	fmt.Println(map3)
	// map是引用类型的 需要使用make方法来分配内存

	map4 := make(map[string]string)
	map5 := map[string]string{}
	fmt.Println(map4 == nil)
	fmt.Println(map5 == nil)
	// 上面两种map的创建方式是相同的

	// map也是有容量的
	map6 := make(map[string]string, 100)
	fmt.Println(len(map6))
	// 切片作为map的值
	// 在有一个key对应多个value的情况下
	mp1 := make(map[int][]int)
	mp1[0] = []int{1, 2, 3}
	mp1[1] = []int{}
	mp1[2] = make([]int, 2, 3)
	fmt.Println(mp1[0][1])
	fmt.Println(mp1[1] == nil)

}
