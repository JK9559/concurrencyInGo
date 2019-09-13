/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {

	orDone := func(done <-chan interface{}, c <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for {
				select {
				case <-done:
					return
				case val, ok := <-c:
					if ok == false {
						return
					}
					select {
					case valStream <- val:
					case <-done:
					}
				}
			}

		}()
		return valStream
	}

	bridge := func(done <-chan interface{}, chanStream <-chan <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for {
				var stream <-chan interface{}
				select {
				case mayStream, ok := <-chanStream:
					if ok == false {
						return
					}
					stream = mayStream
				case <-done:
					return
				}
				for val := range orDone(done, stream) {
					select {
					case valStream <- val:
					case <-done:
					}
				}
			}
		}()
		return valStream
	}

	genVals := func() <-chan <-chan interface{} {
		chanStream := make(chan (<-chan interface{}))
		go func() {
			defer close(chanStream)
			for i := 0; i < 10; i++ {
				// 这里声明了一个channel 缓冲区长度为1 因为这里是模拟生成10个有数据的channel，再生成一个
				// 数据为channel的channel
				stream := make(chan interface{}, 1)
				stream <- i
				close(stream)
				chanStream <- stream
			}
		}()
		return chanStream
	}

	for v := range bridge(nil, genVals()) {
		fmt.Printf("%v ", v)
	}
}
