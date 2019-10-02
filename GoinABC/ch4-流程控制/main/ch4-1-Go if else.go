/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {
	// if语句的特殊写法
	var va = 5
	// 可以再if表达式前加一个执行语句 再根据变量值进行判断
	if i := va; i == 4 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
