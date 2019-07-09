package main

import "fmt"

func main() {
	// 1. 直接声明(声明)
	var slice1 []int
	// 以这种方式创建的slice是一个nil的slice。长度 容量均为0，并且slice内部的也没有分配内置数组
	// 与nil比较为true
	fmt.Println(slice1 == nil)

	// 2. new方式(声明)
	var slice2 = *new([]int)
	// 创建了一个nil切片 与nil比较为true
	fmt.Println(slice2 == nil)

	// 1. 直接声明(创建) 长度 容量为0，但是slice指向的数组有分配地址
	var slice11 = []int{}
	// 与nil比较为false
	fmt.Println(slice11 == nil)

	// 2. new方式(创建) 长度 容量为0，但是slice指向的数组有分配地址
	var slice12 = make([]int, 0)
	// 与nil比较为false
	fmt.Println(slice12 == nil)

	// 直接将第9个位置的元素置为100，其他未注明的元素默认0值
	s1 := []int{1, 2, 3, 4, 9: 100}
	fmt.Println(s1)

	// make函数需要传入三个参数：切片类型、长度、容量。容量可以不传，默认和长度相同
	s2 := make([]int, 5, 10)
	s2[2] = 2
	fmt.Println(s2)

}
