/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {
	// 在panic发生后 会执行defer语句
	defer fmt.Println("close the file")
	defer fmt.Println("close the goroutine")
	panic("crash")
}
