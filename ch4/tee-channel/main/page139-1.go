/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {

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

	orDone := func(done, c <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for {
				select {
				case <-done:
					return
				case v, ok := <-c:
					if ok == false {
						return
					}
					select {
					case valStream <- v:
					case <-done:
					}
				}
			}
		}()
		return valStream
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

	tee := func(done <-chan interface{}, in <-chan interface{}) (<-chan interface{}, <-chan interface{}) {
		out1 := make(chan interface{})
		out2 := make(chan interface{})

		go func() {
			defer close(out1)
			defer close(out2)

			for val := range orDone(done, in) {
				// 这步创建了俩chan 和out1 和 out2的数据共享 也就是更改了out11, out22
				// 也会更改out1, out2 的值
				var out11, out22 = out1, out2
				for i := 0; i < 2; i++ {
					select {
					case <-done:
					case out11 <- val:
						// 这步将out11 的状态置为 nil(注意 chan的状态有nil 关闭 激活状态) 并非out1的状态
						out11 = nil
					case out22 <- val:
						out22 = nil
					}
				}
			}
		}()
		return out1, out2
	}

	done := make(chan interface{})
	defer close(done)

	out1, out2 := tee(done, take(done, repeat(done, 1, 2, 3, 4, 6, 7, 8), 4))

	for val1 := range out1 {
		fmt.Printf("out1: %v, out2: %v\n", val1, <-out2)
	}

}
