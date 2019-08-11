/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"math/rand"
)

/**
运行结果
*/
func main() {
	newRandStream := func() <-chan int {
		newStream := make(chan int)
		go func() {
			// 以下这个print说明了 是进入到了这个go func 但是没有defer执行 所以可以判断
			// 这个go func没有结束 所以 newStream 没有关闭 造成泄漏
			fmt.Println("in the go func")
			defer fmt.Println("newRandStream is exiting")
			defer close(newStream)
			for {
				newStream <- rand.Int()
			}
		}()
		return newStream
	}

	randStream := newRandStream()
	fmt.Println("3 random ints:")
	for i := 0; i < 3; i++ {
		fmt.Println(<-randStream)
	}
}
