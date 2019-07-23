/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"time"
)

/**
和switch不同 select中的case 是没有执行顺序的。
*/

func main() {
	start := time.Now()
	c := make(chan interface{})

	go func() {
		time.Sleep(5 * time.Second)
		close(c)
	}()

	fmt.Println("Blocking on read ...")

	select {
	case <-c:
		fmt.Printf("UnBlocked %v later\n", time.Since(start))
	}
}
