/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"time"
)

/*
Go or-channel模式 为多个协程并发 如果有一个协程结束 那么这一组协程都将结束
*/

func main() {
	// 声明一个函数 接收参数为 <-chan的切片，返回参数为<-chan
	var or func(channels ...<-chan interface{}) <-chan interface{}

	or = func(channels ...<-chan interface{}) <-chan interface{} {
		fmt.Println("place1 len of channels", len(channels))
		switch len(channels) {
		case 0:
			return nil
		case 1:
			//fmt.Println(channels[0])
			return channels[0]
		}

		orDone := make(chan interface{})
		//fmt.Println(orDone)
		go func() {
			fmt.Println("place2 len of channels", len(channels))
			defer fmt.Println("closed")
			defer close(orDone)
			switch len(channels) {
			case 2:
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default:
				select {
				case <-channels[0]:
					fmt.Println("this is chan 0")
				case <-channels[1]:
					fmt.Println("this is chan 1")
				case <-channels[2]:
					fmt.Println("this is chan 2")
				case <-or(append(channels[3:], orDone)...):
					fmt.Println("this is chan 3")
				}
			}
		}()
		return orDone
	}

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			//fmt.Println("this func is:", after)
			defer fmt.Println("this is:", after)
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()

	<-or(
		sig(2*time.Hour),
		//sig(1*time.Minute),
		sig(1*time.Hour),
		sig(2*time.Minute),
		sig(1*time.Second),
	)

	// 执行流程是
	// 1. 传入3个只读channel 长度为3 创建orDone 进入go func的default分支。
	// 2. 分别执行default内的4个case,前三个case由于睡眠时间未到，全为阻塞状态
	// 3. 第四个分支调递归调用or函数，此时入参只有orDone 阻塞，所以在等待1秒钟之后执行完毕释放orDone

	// 当传入为4个channel的时候
	// 1. 长度为4 创建orDone 进入go func的default分支。
	// 2. 分别执行default内的4个case,前三个case由于睡眠时间未到，全部阻塞
	// 3. 第四个分支递归调用or函数，入参为1second chan和orDone，读取1second后 递归结束
	fmt.Printf("done after %v", time.Since(start))
}
