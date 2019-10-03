/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func modifyVal(s []int) {
	s[1] = 3
	fmt.Printf("%p\n", &s)
	fmt.Println("in func:", s)
}

func modifyVal1(s []int) {
	s = append(s, 3)
	fmt.Printf("%p\n", &s)
	fmt.Println("in func:", s)
}

func modifyVal2(s *[]int) {
	*s = append(*s, 3)
	fmt.Println("in func:", *s)
}

func main() {
	// 默认采取值传递的方式来传递参数，也就是传递参数的副本。函数内部职能对副本进行操作，无法改变原值
	// 可以将参数地址传递给函数，就是按引用传递
	// 本质上引用传递传递指针 是将指针的值复制 但是不会复制指针指向的值。按引用传递也是按值传递

	// 传递指针的消耗比传递副本要小很多

	// slice map interface channel 默认是传递引用的

	// 总结 Go的全部都是值传递的，都是副本。
	// 只是有些是非引用类型 int struct等
	// 有些是引用类型 slice map interface channel

	// 尽管有时对于引用类型可以直接在函数修改原值，但是有些情况还是需要传递指针
	v1 := []int{1, 2}
	modifyVal(v1)
	fmt.Println(v1)

	// 因为这里操作的只是切片的副本 不是切片本身
	v2 := []int{1, 2}
	modifyVal1(v2)
	fmt.Println(v2)

	// 这里就需要传递指针
	v3 := []int{1, 2}
	modifyVal2(&v3)
	fmt.Println(v3)
}
