package main

import (
	"fmt"
	"sync"
)

/**
WaitGroup
*/
func main() {
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
