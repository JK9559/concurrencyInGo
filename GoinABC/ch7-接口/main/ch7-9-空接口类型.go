/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {
	// 空接口没任何方法，空接口可以保存任何值，也可以从空接口中取原值
	// 将值保存到空接口
	var any interface{}

	any = 1
	fmt.Printf("%T ", any)
	fmt.Println(any)

	any = "hello"
	fmt.Printf("%T ", any)
	fmt.Println(any)

	any = false
	fmt.Printf("%T ", any)
	fmt.Println(any)

	// 从空接口中获取值
	var a int = 1
	var i interface{} = a
	// 这句话将会编译报错 无法将i视为int类型赋值给b
	// var b int = i
	var b int = i.(int)
	fmt.Println(b)

	// 空接口的值比较
	// 空接口在保存了不同的值之后 可以和其他变量值一样使用==来进行比较操作
	// 有以下几种特性
	// 1. 类型不同的空接口间的比较结果不同
	var x interface{} = 100
	var y interface{} = "hi"
	fmt.Println(x == y)

	var e interface{} = [10]int{1, 2, 3}
	var f interface{} = [10]int{1, 2, 3}
	fmt.Println(e == f)

	var c interface{} = []int{10}
	var d interface{} = []int{20}
	fmt.Println(c == d)

	// 以下
	// map 		panic不可比较
	// []T切片 	panic不可比较
	// channel	可比较 只有同一个通道才为true(同一个make生成的) 否则为false
	// array,struct,function 可比较
}
