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
		defer fmt.Println("go func is Done")
		time.Sleep(5 * time.Second)
		close(c)
	}()

	fmt.Println("Blocking on read ...")

	select {
	// 这里的channel c一直为阻塞状态 直到close关闭了c，才触发了case语句 结束了对main goroutine的阻塞
	case <-c:
		fmt.Printf("UnBlocked %v later\n", time.Since(start))
	}
}
