/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

/**
首先明确：chan的默认值为nil 当我们对nil的channel做一些操作时(读 写 关闭)
*/

func main() {
	// 从nil的channel中读取数据
	var dataStream chan interface{}
	<-dataStream
	// 报错
	/**
	fatal error: all goroutines are asleep - deadlock!

	goroutine 1 [chan receive (nil chan)]:
	main.main()
		D:/goWorkSpace/Go/src/concurrencyInGo/ch3/channel/main/page87-1.go:14 +0x30
	*/
	// 向nil的channel中写入数据
	var dataStream1 chan interface{}
	dataStream1 <- struct{}{}
	// 报错
	/**
	fatal error: all goroutines are asleep - deadlock!

	goroutine 1 [chan send (nil chan)]:
	main.main()
		D:/goWorkSpace/Go/src/concurrencyInGo/ch3/channel/main/page87-1.go:25 +0x53
	*/
	// 关闭(close)一个nil的channel
	var dataStream3 chan interface{}
	close(dataStream3)
	// 报错
	/**
	panic: close of nil channel

	goroutine 1 [running]:
	main.main()
		D:/goWorkSpace/Go/src/concurrencyInGo/ch3/channel/main/page87-1.go:36 +0x31
	*/
}
