/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	si := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range si {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	// 大多数情况下 for循环块里的代码是在同一个goroutine里执行的，
	// 为了避免空间的浪费和GC的压力 ，复用了range迭代的临时变量i
	// go run -race 文件名.go
}
