package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for _, str := range []string{"hello", "welcome", "greetings"} {
		// 这个示例中 goroutine 正在运行一个闭包,在该闭包使用str变量的时候
		// 字符串的迭代已经结束 所以取到的都是最后的一个字符串
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(str)
		}()
	}
	wg.Wait()
}
