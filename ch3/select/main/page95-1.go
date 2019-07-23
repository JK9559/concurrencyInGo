/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"time"
)

/**
如果没有任何可用的channel的时候，使用select语句读取会报错 deadlock
但是如果有了可执行的channel，就会执行可执行的那个 就算其他的deadlock了也可以
*/

func main() {
	var c chan interface{}
	c1 := make(chan interface{})
	c2 := make(chan interface{})
	close(c2)
	select {
	case <-c:
		fmt.Println("selected channel c")
	// time.After函数通过传入time.Duration返回一个数值并写入channel，该channel会返回执行后的时间
	case <-time.After(5 * time.Second):
		fmt.Println("Time out")
	case <-c1:
		fmt.Println("selected channel c1")
	case <-c2:
		fmt.Println("hah")
	}
}
