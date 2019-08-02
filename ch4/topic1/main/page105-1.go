/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {
	dowork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			fmt.Println("yes")
			defer fmt.Println("dowork exited")
			defer close(completed)
			for s := range strings {
				fmt.Println(s)
			}
		}()
		return completed
	}

	// 即使传入的channel不为空 运行时也会发生deadlock 因为不论如何dowork内的go func都不会被执行
	// strings := make(chan string)
	// strings <- "abccd"
	// dowork(strings)

	// 因为dowork内的go func不论如何都不会执行 函数内部的completed channel创建了就不会被释放 造成泄漏
	dowork(nil)
	fmt.Println("Done")
}
