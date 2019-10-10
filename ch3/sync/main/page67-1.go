/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"sync"
)

/**
Cond的Broadcast
*/
func main() {
	type Button struct {
		Clicked *sync.Cond
	}
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	// 问：为什么注释掉goroutineRunning之后 会报deadlock的错？
	// 答：因为观察subscribe结构可以看到内部有一个函数fn() 传入的函数里都有一个clickRegistered.Done()
	// 如果不针对subscribe内部的func进行同步的话，就没有机会执行clickRegistered.Done() 自然会deadlock
	subscribe := func(c *sync.Cond, fn func()) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)

		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			// 这里等待一个信号来激活
			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}

	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)
	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing window.")
		clickRegistered.Done()
	})

	subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialog box!")
		clickRegistered.Done()
	})

	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked")
		clickRegistered.Done()
	})

	// 这里给了一个激活的信号
	button.Clicked.Broadcast()

	clickRegistered.Wait()

	/*
		Mouse clicked
		Maximizing window.
		Displaying annoying dialog box!
	*/
}
