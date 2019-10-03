/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	test()
	end := time.Now()
	res := end.Sub(start)
	fmt.Printf("%s\n", res)
}

func test() {
	sum := 0
	for i := 0; i < 100000000; i++ {
		sum += i
	}
}
