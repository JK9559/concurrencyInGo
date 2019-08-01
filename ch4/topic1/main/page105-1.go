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

	dowork(nil)
	fmt.Println("Done")
}
