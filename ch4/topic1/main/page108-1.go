/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
关于go func的执行 为什么有的时候 只需要sleep一段时间就可以有运行的机会
首先 明确定义了一个go func就意味着 当调用到这里的时候 gf是开辟出了一个新的分支 而这个分支是否有机会执行
取决于调起他的协程执行的快慢(也就是是否会等他)，如果调起他的协程执行的时间久 那么就会在调起他的协程结束之前
调用gf 这时 gf也就会执行。可以看到 每次是阻塞了main协程 才调起了gf 最终因为gf使得main不再阻塞 从而使main继续执行
*/

func main() {
	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream is exiting")
			defer close(randStream)
			for {
				select {
				case randStream <- rand.Int():
				case <-done:
					return
				}
			}
		}()
		return randStream
	}

	done := make(chan interface{})
	randStream := newRandStream(done)

	fmt.Println("3 random ints:")
	for i := 0; i < 3; i++ {
		fmt.Println(<-randStream)
	}
	close(done)
	// 想想这里是为什么，如果不睡眠 就会像之前的例子一样 不会对channel进行close 从而泄露
	// 因为这里需要通过sleep来模拟阻塞main协程 使gf能有机会执行完
	time.Sleep(1 * time.Second)
}
