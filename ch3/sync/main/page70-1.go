/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"sync"
)

/**
once只计算调用Do方法的次数，而不是多少次唯一调用Do方法
*/
func main() {
	var count int

	increment := func() {
		fmt.Println("This is Increment func")
		count++
	}

	decrement := func() {
		fmt.Println("This is Decrement func")
		count--
	}

	var once sync.Once
	once.Do(decrement)
	once.Do(increment)

	fmt.Printf("Count is %d\n", count)
}
