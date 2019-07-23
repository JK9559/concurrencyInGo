/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"time"
)

/**
当没用可用的channel时，select做什么。答案是有默认值
*/

func main() {
	start := time.Now()
	var c1, c2 chan interface{}

	select {
	case <-c1:
		fmt.Println("this is c1 channel.")
	case <-c2:
		fmt.Println("this is c2 channel.")
	default:
		fmt.Printf("no channel can be used. after %v\n", time.Since(start))
	}
}
