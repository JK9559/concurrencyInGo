/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"bytes"
	"fmt"
	"os"
)

/**
带缓冲区的channel
*/
func main() {
	// 创建一个内存缓存区，来帮助减少输出的不确定性。比直接写stdout要快
	var stdoutBuff bytes.Buffer
	// 确保在进程退出之前缓冲区内容都写到了stdout
	defer stdoutBuff.WriteTo(os.Stdout)

	// 1. intStream := make(chan int)
	intStream := make(chan int, 4)
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done")
		for i := 0; i < 5; i++ {
			_, _ = fmt.Fprintf(&stdoutBuff, "Sending: %v.\n", i)
			// 2. time.Sleep(1 * time.Second)
			intStream <- i
		}
	}()

	for integer := range intStream {
		_, _ = fmt.Fprintf(&stdoutBuff, "Received: %v.\n", integer)
	}

	/**
	代码的输出为：
	Sending: 0.
	Sending: 1.
	Sending: 2.
	Sending: 3.
	Sending: 4.
	Producer Done
	Received: 0.
	Received: 1.
	Received: 2.
	Received: 3.
	Received: 4.
	顺序的。
	当使用了 无缓存的chan的时候 是乱序的 见注释1.
	当在循环中使用了sleep 也就是减慢写入的速度时，顺序也没法保证 见注释2.
	*/
}
