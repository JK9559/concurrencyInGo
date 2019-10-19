/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"sync"
)

/**
问：为什么使用Pool 而不是在运行时实例化对象？
答：Pool可以进行预分配内存，通过这个例子可以看出来，通过Pool获取的实例都在使用完毕时候通过defer Put回了Pool里，及时回收并重复利用
   并且最终只分配了额外的4kb内存。如果不使用Pool，只是依靠运行时进行对象的实例化，可能一直在分配内存，而没有得到回收，就导致内存占用
   越来越高。最终溢出？
*/
func main() {
	var numCalCreated int
	calcPool := &sync.Pool{
		New: func() interface{} {
			numCalCreated++
			mem := make([]byte, 1024)
			return &mem
		},
	}

	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	fmt.Printf("%d instance were created!\n", numCalCreated)

	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()
			mem := calcPool.Get().(*[]byte)
			//time.Sleep(1 * time.Millisecond)
			defer calcPool.Put(mem)
		}()
	}
	wg.Wait()
	fmt.Printf("%d instance were created!", numCalCreated)

	/*
		4 instance were created!
		8 instance were created!
	*/
}
