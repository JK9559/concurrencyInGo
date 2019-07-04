package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for _, str := range []string{"hello", "welcome", "good day"} {
		wg.Add(1)
		go func(str string) {
			defer wg.Done()
			fmt.Println(str)
			// 这里就将当前迭代的变量传递给了闭包。创建了一个字符串的副本
			// 从而确保当goroutine运行时 可以引用适当的字符串
		}(str)
	}
	wg.Wait()
}
