/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan interface{})
	// 因为for循环内部的select有default所以select一定不会被阻塞
	// 那么为什么下面的go func还能执行，因为sleep阻塞了main goroutine使go func有机会执行
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	workCount := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}
		workCount++
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("Achieved %d loops before signalled to stop\n", workCount)

	/*
		Achieved 5 loops before signalled to stop
	*/
}
