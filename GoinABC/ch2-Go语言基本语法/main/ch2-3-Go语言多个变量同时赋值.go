/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {
	// 多变量赋值
	var a int = 100
	var b int = 200

	fmt.Println(a, b)

	a, b = b, a

	fmt.Println(a, b)
}
