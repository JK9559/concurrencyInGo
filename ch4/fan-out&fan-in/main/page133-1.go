/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

/*
找到50000000... 中的10个随机素数
*/
func main() {
	/*isPrime := func(num int) bool {
		if num == 2 || num == 3 {
			return true
		}

		if num%6 != 1 && num%6 != 5 {
			return false
		}

		sqrt := int(math.Sqrt(float64(num)))

		for i := 5; i <= sqrt; i += 6 {
			if num%(i) == 0 || num%(i+2) == 0 {
				return false
			}
		}
		return true
	}*/

	isPrimeSlow := func(num int) bool {
		//sqrt := int(math.Sqrt(float64(num)))

		for i := 2; i < num; i++ {
			if num%i == 0 {
				return false
			}
		}
		return true
	}

	primeFinder := func(done <-chan interface{}, valuesStream <-chan int) <-chan int {
		primeFindStream := make(chan int)
		go func() {
			defer close(primeFindStream)
			for {
				select {
				case <-done:
					return
				case v := <-valuesStream:
					if isPrimeSlow(v) {
						primeFindStream <- v
					}
				}
			}
		}()
		return primeFindStream
	}

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

	take := func(done <-chan interface{}, valuesStream <-chan int, nums int) <-chan int {
		takeStream := make(chan int)
		go func() {
			defer close(takeStream)
			for i := 0; i < nums; i++ {
				select {
				case <-done:
					return
				case v := <-valuesStream:
					takeStream <- v
				}
			}
		}()
		return takeStream
	}

	toInt := func(done <-chan interface{}, valuesStream <-chan interface{}) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for {
				select {
				case <-done:
					return
				case v := <-valuesStream:
					intStream <- v.(int)
				}
			}
		}()
		return intStream
	}

	fanIn := func(done <-chan interface{}, channels ...<-chan int) <-chan int {
		var wg sync.WaitGroup
		multiplexedStream := make(chan int)
		multiplex := func(c <-chan int) {
			defer wg.Done()
			for i := range c {
				select {
				case <-done:
					return
				case multiplexedStream <- i:
				}
			}
		}

		wg.Add(len(channels))
		for _, v := range channels {
			go multiplex(v)
		}

		go func() {
			wg.Wait()
			close(multiplexedStream)
		}()

		return multiplexedStream
	}

	done := make(chan interface{})
	defer close(done)

	randF := func() interface{} { return rand.Intn(50000000) }

	randIntStream := toInt(done, repeatFn(done, randF))
	start := time.Now()
	fmt.Println("Primes:")
	for num := range take(done, primeFinder(done, randIntStream), 10) {
		fmt.Printf("\t%d\n", num)
	}
	fmt.Printf("Search takes: %v\n", time.Since(start))

	done1 := make(chan interface{})
	defer close(done1)
	numFinders := runtime.NumCPU()
	finders := make([]<-chan int, numFinders)
	for i := 0; i < numFinders; i++ {
		finders[i] = primeFinder(done1, randIntStream)
	}

	start1 := time.Now()
	for num := range take(done1, fanIn(done1, finders...), 10) {
		fmt.Printf("\t%d\n", num)
	}
	fmt.Printf("Search takes: %v\n", time.Since(start1))

}
