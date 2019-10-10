package main

import (
	"fmt"
	"sync"
)

/**
sync.Mutex 互斥锁
*/

func main() {
	var count int
	// 互斥锁
	var lock sync.Mutex

	increment := func() {
		// 我们请求对临界区的独占 使用互斥锁来解决
		lock.Lock()
		// 声明已经完成了对临界区的独占使用 释放互斥锁
		defer lock.Unlock()
		count++
		fmt.Printf("Incrementing:%d\n", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrementing:%d\n", count)
	}

	var arithmetic sync.WaitGroup
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()
	fmt.Println("Arithmetic complete")

	/*
		Incrementing:1
		Decrementing:0
		Decrementing:-1
		Incrementing:0
		Decrementing:-1
		Incrementing:0
		Incrementing:1
		Incrementing:2
		Incrementing:3
		Decrementing:2
		Decrementing:1
		Decrementing:0
		Arithmetic complete
	*/
}
