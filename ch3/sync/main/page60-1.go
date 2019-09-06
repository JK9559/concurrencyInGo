package main

import (
	"fmt"
	"sync"
)

/**
WaitGroup
*/
func main() {
	// 这里对于WaitGroup的操作需要传引用 否则 函数体内的 wg.Done() 不生效报错
	hello := func(wg *sync.WaitGroup, i int) {
		defer wg.Done()
		fmt.Printf("Hello from %v!\n", i)
	}

	var wg sync.WaitGroup
	const numGreeters = 5
	// 在循环前只调用一次Add来追踪一组goroutine
	wg.Add(numGreeters)
	for i := 1; i <= numGreeters; i++ {
		go hello(&wg, i)
	}
	wg.Wait()
}
