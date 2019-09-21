/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {
	// 数组写法如下
	// var 数组名 [数量]类型
	// 元素数量必须是在编译期确定的整型数值，不能是运行时才能确认大小的值

	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2])

	// 表示数组长度是根据初始化值的个数来计算的
	q = [...]int{1, 2, 3}
	fmt.Printf("%T\n", q)

	type Currency int

	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)

	symbol := [...]string{EUR: "€", RMB: "￥", GBP: "￡", USD: "$"}

	fmt.Println(RMB, symbol[RMB])

	x := [...]int{99: -1}
	// 这里直接标记数组x的第99个元素(从0开始编号)的元素为-1 所以 该数组长度为100,且第99个元素为-1 剩余元素为0
	fmt.Printf("%d %d %d\n", x[0], x[99], x[98])

	// 两个数组的比较
	// 只有当两个数组中的所有元素都是相等的时候 才是相等的 == 或 != 不能用 < >
	arrA := [2]int{1, 2}
	arrB := [...]int{1, 2}
	arrC := [2]int{1, 3}
	fmt.Println(arrA == arrB, arrA == arrC, arrB == arrC)
	// 以下报错
	// fmt.Println(arrA < arrC)
	// arrD := [3]int{1, 2}
	// 下面直接报错 因为arrD和arrA的类型不同
	// fmt.Println(arrA == arrD)
}
