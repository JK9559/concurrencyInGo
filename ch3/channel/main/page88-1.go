/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {
	/**
	首先是关于excel表格中的注意事项，作用如下：
	1. 说明对于channel的操作规则略复杂，不能在使用的时候毫无规划
	2. 说明使用时 需要有条理 方便维护

	解决办法为：分配channel的所有权
	channel的拥有者 所享有的权限：1. 实例化channel 2. 写操作 3. 关闭 4. 转移channel所有权给另一个goroutine
	channel的使用者 只有读取channel的权限

	*/
	// channel拥有者
	// =====>>>尽量保持channel的所有权范围很小<<<=====
	chanOwner := func() <-chan int {
		// 1. 初始化channel
		resultStream := make(chan int, 5)
		go func() {
			// 3. 关闭channel
			defer close(resultStream)
			// 2. 写入channel
			for i := 0; i < 6; i++ {
				resultStream <- i
			}
		}()
		// 4. 暴露只读channel
		return resultStream
	}
	// channel使用者
	resultStream := chanOwner()
	for rst := range resultStream {
		fmt.Println(rst)
	}

	fmt.Println("Done receive!")
}
