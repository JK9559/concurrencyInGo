package main

import (
	"fmt"
	"sync"
)

func main()  {
	var wg sync.WaitGroup
	str := "hello"
	wg.Add(1)
	// 闭包可以从创建他们的作用域中获取变量 是在原值的引用上运行的
	go func() {
		defer wg.Done()
		str = "welcome"
	}()
	wg.Wait()
	// 可以看到闭包内 改变了str的值
	fmt.Println(str)
}
