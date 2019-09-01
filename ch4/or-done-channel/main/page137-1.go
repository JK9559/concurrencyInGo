/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

/*
or-done-channel 尚未理解
1. 初步理解为 当done channel尚未被关闭的时候 myChan可能已经提前于done被关闭了，这时 函数也应该结束
*/
func main() {
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

	done := make(chan interface{})
	myChan := make(chan interface{})

	defer close(myChan)
	defer close(done)

	for val := range orDone(done, myChan) {
		fmt.Println(val)
	}
}
