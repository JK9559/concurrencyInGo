/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {
	// 变量初始化标准格式
	var a int = 10
	fmt.Println(a)

	// 编译期推导类型格式
	var b = 20.0
	fmt.Println(b)

	// 短变量声明并初始化
	c := 30
	fmt.Println(c)
}
