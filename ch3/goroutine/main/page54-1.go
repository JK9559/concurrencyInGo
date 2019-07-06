package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}

	var wg sync.WaitGroup

	// 使用一个永远不会退出的goroutine
	noop := func() { wg.Done(); <-c }

	const numGoroutines = 1e4

	wg.Add(numGoroutines)
	// 计算创建goroutine前消耗的内存总量
	before := memConsumed()
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	// 计算创建goroutine后消耗的内存总量
	after := memConsumed()
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)
}
