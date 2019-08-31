/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 生成器
	// 该生成器为重复调用函数的生成器
	repeatFn := func(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
		valuesStream := make(chan interface{})
		go func() {
			defer close(valuesStream)
			for {
				select {
				case <-done:
					return
				case valuesStream <- fn():
				}
			}
		}()
		return valuesStream
	}

	take := func(done <-chan interface{}, valuesStream <-chan interface{}, nums int) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < nums; i++ {
				select {
				case <-done:
					return
				// 这样写的话 print出来的结果为0xc00004a0c0 原因未知
				// case takeStream <-valusStream:
				// 这样写的话 print出来的结果为正常
				case v := <-valuesStream:
					takeStream <- v
				}
			}
		}()
		return takeStream
	}

	done := make(chan interface{})
	defer close(done)

	// 随机生成一个1到29的int数字
	randF := func() interface{} { return rand.Int()%29 + 1 }

	for num := range take(done, repeatFn(done, randF), 10) {
		fmt.Println(num)
	}
}
