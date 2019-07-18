/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"sync"
)

/**
对channel的close也是一种同时给多个goroutine发信号的方法。
如果有n个goroutine在一个channel上等待，而不是在channel上写n次来打开每个goroutine，
就可以简单的关闭channel
*/
func main() {
	begin := make(chan interface{})
	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("now it is %v\n", i)
		}(i)
	}
	// 可以看到当使用begin channel的读取作为阻塞时，在close(begin)发生之前，不会运行for内的语句
	// 结果就是下面的print始终先被触发，然后才是for内的print
	// 如果去掉了读取begin channel的阻塞，那么这些都是无序的print
	fmt.Println("before close begin_channel")
	close(begin)
	wg.Wait()
}
