/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

func main() {
	// map在并发情况下 只读是线程安全的 但是同时读写线程就不安全
	m := make(map[int]int)

	go func() {
		for {
			m[1] = 1
		}
	}()

	go func() {
		for {
			_ = m[1]
		}
	}()

	for {

	}

	// 报错 fatal error: concurrent map read and map write
}
