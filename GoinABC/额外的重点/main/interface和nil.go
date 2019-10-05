/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// https://studygolang.com/articles/1908
	// 先说interface 在底层 interface作为了两个成员来实现 一个类型 一个是值
	// 接下来用到反射
	var val interface{} = int64(58)
	// 打印出的是 int64 是因为已经显示的将类型int64的数据58 赋值给了interface变量val
	// val底层是 int64,58
	fmt.Println(reflect.TypeOf(val))
	val = 50
	// 打印出的是 int 因为字面量的整数在golang中默认的类型是int
	// 所以这时的底层结构就变成了 int,50
	fmt.Println(reflect.TypeOf(val))

	fmt.Println(">=====<")
	// type用来存储变量的动态类型 data用来存储变量的具体数据
	// 该值被称为接口的动态值，是一个任意的具体值，而该接口的类型则是该值的类型。
	// 只有当内部值和类型都未设置(nil,nil)时，一个接口的值才是nil。
	var p interface{}
	// 这里的接口类型就是为空
	fmt.Println(reflect.TypeOf(p))
	// 因此这里的p为nil
	fmt.Println(p == nil)

	var p1 *int
	fmt.Println(p1 == nil)
	// 这里将一个*int类型的指针赋值给了接口类型p
	p = p1
	// 可以看到这时 p的类型为*int 值为空
	fmt.Println(reflect.TypeOf(p))
	// 但是p不为nil
	fmt.Println(p == nil)

	fmt.Println(">=====<")
	// 关于(*interface{})(nil)
	var pval interface{} = (*interface{})(nil)
	// b := 3
	// 这里编译报错
	// pval = &b
	// 这里 panic: runtime error: invalid memory address or nil pointer dereference
	// *pval = 3
	// 原因在于 pval指向的是无效的内存地址
	fmt.Println(reflect.TypeOf(pval))
	fmt.Println(pval == nil)

	// 在自定义错误类型的时候 可能会用到
	// 见 ch3-14

}
