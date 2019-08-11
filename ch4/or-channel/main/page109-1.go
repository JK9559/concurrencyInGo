/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"time"
)

func main() {
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
				// 执行了or 但是执行结果为阻塞 因为orDone未被读取被阻塞 之后close后 才释放，所以是先 channels[1]后 执行完毕
				// 这个go func 最后释放orDone
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
			//defer fmt.Println("this is:", after)
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()

	<-or(
		sig(2*time.Hour),
		//sig(1*time.Minute),
		//sig(1*time.Hour),
		sig(1*time.Second),
		sig(2*time.Minute),
	)

	fmt.Printf("done after %v", time.Since(start))
}
