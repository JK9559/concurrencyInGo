/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func report1(a, b int, s string) (int, int, string) {
	return a, b, s
}

// 如果一个函数有命名的返回值，可以省略return语句的操作数，称为裸返回
// 代码逻辑容易不清晰 慎用
func report2(a, b int, s string) (a1, b1 int, s1 string) {
	a1 = a
	b1 = b
	s1 = s
	return
}

func main() {
	a, b := 0, 1
	s := "abc"

	a1, b1, s1 := report1(a, b, s)
	fmt.Println(a1, b1, s1)

	a2, b2, s2 := report2(a, b, s)
	fmt.Println(a2, b2, s2)
}
