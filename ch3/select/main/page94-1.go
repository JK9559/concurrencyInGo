/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

/**
当多个channel有数据可供给下游读取时,可以看到程序的输出大概有一半的情况是从c1读取 一半是从c2读取
Golang运行的时候在一组case语句中进行了伪随机选择。因为对于select来说无法确定哪个channel更优先
*/

func main() {
	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})
	close(c2)

	var c1Count, c2Count int
	for i := 0; i < 1000; i++ {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		}
	}
	fmt.Printf("c1Count is: %d\nc2Count is: %d", c1Count, c2Count)
}
