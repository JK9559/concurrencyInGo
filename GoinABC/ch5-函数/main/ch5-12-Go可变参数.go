/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

// 可变参数类型 ...type 被称为语法糖 其本质是一个切片 也就是[]type
func myfunc(args ...int) {
	for _, v := range args {
		fmt.Println(v)
	}
	// 观察调用方法 也是切片
	myfunc2(args...)
	myfunc2(args[1:]...)
}

func myfunc2(args ...int) {
	fmt.Println(args)
}

func main() {

}
