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
}
