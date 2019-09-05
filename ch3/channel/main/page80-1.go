/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {
	// 通过<-读出的返回值 可以选择返回两个值
	stringStream := make(chan string)
	go func() {
		stringStream <- "this is a string"
	}()
	readData1, ok1 := <-stringStream
	fmt.Printf("%v , %v\n", ok1, readData1)

	close(stringStream)

	readData2, ok2 := <-stringStream
	fmt.Printf("%v , %v\n", ok2, readData2)

	/**
	问题：返回的布尔值 表示了什么？
	答案：如果为true，表示了该channel上有新的数据写入，如果为false，表示了该channel已经关闭，此时读到的内容，是默认值
	*/

	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 0; i < 5; i++ {
			intStream <- i
		}
	}()
	// 遍历chan的方式
	// 如果不close chan 那么range会报错
	for integer := range intStream {
		fmt.Printf("%v ", integer)
	}
}
