package main

import (
	"fmt"
	"sync"
)

func main()  {
	// page50 - 1
	var wg sync.WaitGroup
	sayHello := func() {
		defer wg.Done()
		fmt.Println("hello goroutine")
	}
	wg.Add(1)
	go sayHello()
	// 这里阻塞了main goroutine
	wg.Wait()
}
