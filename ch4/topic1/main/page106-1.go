/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"time"
)

/**
设置了几个channel 首先一个原则 为了避免channel的泄漏 必须被close(谁用谁关)
1. dowork内的 completed 需要被关闭，关闭函数在go func内，必须执行go func才行
main goroutine对于 completed channel的读取 就促使必须执行go func
2. 又因为传入的strings channel为空 在select语句中 始终不会被执行，所以引入了done来保证
dowork中的go func可以执行完毕(最终目标是close(completed))，而为了done不被阻塞，最外层
main goroutine中的go func必须执行也就是通过close(done)来向done channel发出信号
3. 在done的控制下，completed channel最终被关闭了
*/

func main() {
	// 这里传入一个done的只读channel
	dowork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("dowork exited")
			defer close(completed)
			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return completed
	}

	done := make(chan interface{})
	completed := dowork(done, nil)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling dowork function")
		close(done)
	}()

	<-completed
	fmt.Println("Done")

}
