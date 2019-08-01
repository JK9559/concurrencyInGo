/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream is exiting")
			defer close(randStream)
			for {
				select {
				case randStream <- rand.Int():
				case <-done:
					return
				}
			}
		}()
		return randStream
	}

	done := make(chan interface{})
	randStream := newRandStream(done)

	fmt.Println("3 random ints:")
	for i := 0; i < 3; i++ {
		fmt.Println(<-randStream)
	}
	close(done)
	// 想想这里是为什么，如果不睡眠 就会像之前的例子一样 不会对channel进行close 从而泄露
	time.Sleep(1 * time.Second)
}
