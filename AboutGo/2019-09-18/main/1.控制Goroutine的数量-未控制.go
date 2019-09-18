/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// 未控制goroutine的数量
	jobCount := 10
	group := sync.WaitGroup{}
	for i := 0; i < jobCount; i++ {
		group.Add(1)
		go func(i int) {
			fmt.Printf("hello %d\n", i)
			time.Sleep(1 * time.Second)
			group.Done()
		}(i)
		fmt.Printf("index: %d,goroutine Num: %d\n", i, runtime.NumGoroutine())
	}
	group.Wait()
	fmt.Println("done")
	// 上述代码可以看到 10个任务的协程 达到了11个 并且协程的执行顺序不固定

	// 通过channel来传递信息 实现减少goroutine的操作

}
