/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

/**
声明和实例化一个channel(读+写、只读、只写)
*/

func main() {
	// 创建一个channel(读+写)
	var dataStream chan interface{}
	dataStream = make(chan interface{})
	fmt.Println(dataStream)

	// 创建一个channel(只读)
	var dataStreamReadOnly <-chan interface{}
	dataStreamReadOnly = make(<-chan interface{})
	fmt.Println(dataStreamReadOnly)

	// 创建一个channel(只写)
	var dataStreamWriteOnly chan<- interface{}
	dataStreamWriteOnly = make(chan<- interface{})
	fmt.Println(dataStreamWriteOnly)

	/**
	通常不会看到单向的channel实例化的。单向的channel可以用作 函数的参数 和 返回类型
	Go会在需要的时候 隐式的将双向的channel转换为单向的channel
	*/
	stringStream := make(chan string)
	go func() {
		stringStream <- "this is a string"
	}()
	fmt.Println(<-stringStream)

	writeStream := make(chan<- interface{})
	readStream := make(<-chan interface{})
	fmt.Println(writeStream)
	fmt.Println(readStream)
	// 以下两个操作会报错
	// <-writeStream : invalid operation: <-writeStream (receive from send-only type chan<- interface {})
	// readStream <- "abc" : invalid operation: readStream <- "abc" (send to receive-only type <-chan interface {})

	// 运行结果
	// 0xc00004a060
	// 0xc00004a0c0
	// 0xc00004a120
	// this is a string
	// 0xc00004a1e0
	// 0xc00004a240

}
