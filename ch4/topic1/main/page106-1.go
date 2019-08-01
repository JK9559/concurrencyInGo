/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	dowork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("dowork exited")
			defer close(completed)
			for {
				fmt.Println("hah")
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
	//
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling dowork function")
		close(done)
	}()

	fmt.Println(<-completed)
	fmt.Println("Done")
}
