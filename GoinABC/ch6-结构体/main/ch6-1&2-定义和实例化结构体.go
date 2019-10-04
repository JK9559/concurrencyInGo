/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	// 1. 基本的实例化形式
	var p1 Point
	p1.X = 1
	p1.Y = 2
	fmt.Printf("%T %v\n", p1, p1)

	// 2. 创建指针类型的结构体
	var p2 = new(Point)
	// 这里的p2为指针类型 但是可以继续使用. 来操作成员 因为使用了语法糖
	// p2.X 相当于 (*p2).X
	p2.X = 2
	p2.Y = 3
	fmt.Printf("%T %v\n", p2, p2)

	// 3. 取结构体的地址实例化
	// 对结构体进行&取地址操作的时候 视为对该类型进行了一次new的实例化操作
	// 最常用
	var p3 = &Point{}
	p3.X = 3
	p3.Y = 4
	fmt.Printf("%T %v\n", p3, p3)

	// 可以使用函数来封装
	var p4 = newPoint(4, 5)
	fmt.Printf("%T %v\n", p4, p4)

}

func newPoint(x, y int) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}
