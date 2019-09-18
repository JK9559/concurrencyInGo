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
	// 通过channel减少goroutine的数量
	jobCount := 10
	poolCount := 3
	group := sync.WaitGroup{}
	jobsChan := make(chan int)
	defer close(jobsChan)
	// 开启固定数量的goroutine来消费传入的值
	for i := 0; i < poolCount; i++ {
		go func() {
			for val := range jobsChan {
				fmt.Printf("hello %d\n", val)
				time.Sleep(1 * time.Second)
				group.Done()
			}
		}()
	}

	// 将消费者用的值i放入channel 利用channel的线程安全传递信息
	for i := 0; i < jobCount; i++ {
		jobsChan <- i
		group.Add(1)
		fmt.Printf("index: %d, goroutine Num: %d\n", i, runtime.NumGoroutine())
	}

	group.Wait()
	fmt.Println("done")

}
