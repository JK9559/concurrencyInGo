/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {

	// 生成重复数据的生成器
	// 传入一组interface的values 不停止的情况下 依次循环将值放入 valuesStream中
	repeat := func(done <-chan interface{}, values ...interface{}) <-chan interface{} {
		valuesStream := make(chan interface{})
		go func() {
			defer close(valuesStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valuesStream <- v:
					}
				}
			}
		}()
		return valuesStream
	}

	// pipeline Stage
	// 取nums个valuesStream中的值 写入到takeStream里
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

	for num := range take(done, repeat(done, 1, 2, 3, 4), 10) {
		fmt.Printf("%v ", num)
	}

}
