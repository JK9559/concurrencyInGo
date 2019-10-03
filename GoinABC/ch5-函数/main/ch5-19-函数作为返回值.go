/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {
	var f = Adder()
	// 开始x为0 当调用f(5) 返回0+5 =5
	fmt.Println(f(5))
	// 这时x为5
	fmt.Println(f(20))
	// 这时x为25
	fmt.Println(f(300))
}

func Adder() func(int) int {
	var x int
	return func(i int) int {
		x += i
		return x
	}
}
