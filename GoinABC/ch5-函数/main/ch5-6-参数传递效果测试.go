/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

type InnerData struct {
	a int
}

type Data struct {
	complex  []int
	instance InnerData
	ptr      *InnerData
}

func passByValue(inFunc Data) Data {
	fmt.Printf("inFunc value: %+v\n", inFunc)

	fmt.Printf("inFunc ptr: %p\n", &inFunc)

	return inFunc
}

func main() {
	// 准备传入函数的结构
	in := Data{
		complex: []int{1, 2, 3},
		instance: InnerData{
			5,
		},
		ptr: &InnerData{1},
	}
	// 输入结构的成员情况
	fmt.Printf("in value: %+v\n", in)
	// 输入结构的指针地址
	fmt.Printf("in ptr: %p\n", &in)
	// 传入结构体，返回同类型的结构体
	out := passByValue(in)
	// 输入结构的成员情况
	fmt.Printf("in value: %+v\n", in)
	// 输入结构的指针地址
	fmt.Printf("in ptr: %p\n", &in)
	// 输出结构的成员情况
	fmt.Printf("out value: %+v\n", out)
	// 输出结构的指针地址
	fmt.Printf("out ptr: %p\n", &out)
}

// 结论 值复制发生在参数进入函数的时候 以及返回值离开函数的时候
// 指针在传递的时候 只是传递指针的值，不会复制指针指向的部分
