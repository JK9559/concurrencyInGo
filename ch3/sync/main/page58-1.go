package main

import (
	"fmt"
	"sync"
	"time"
)

/**
WaitGroup
*/
func main() {
	var wg sync.WaitGroup

	// 参数为1 表示开始一个goroutine
	wg.Add(1)
	go func() {
		// defer 表示在goroutine退出前执行Done操作，向WaitGroup表示当前的goroutine已经退出了
		defer wg.Done()
		fmt.Println("1st goroutine sleeping")
		time.Sleep(1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine sleeping")
		time.Sleep(2)
	}()

	// Wait将阻塞 main goroutine，直到所有的goroutine表明他们已经退出
	wg.Wait()
	fmt.Println("All goroutine complete!")
}
