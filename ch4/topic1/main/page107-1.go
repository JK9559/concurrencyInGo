/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	newRandStream := func() <-chan int {
		newStream := make(chan int)
		go func() {
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
