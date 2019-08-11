/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 这里的checkStatus 在进行错误处理的时候 直接print了 最好的还是保持函数功能
	// 将错误包装返回到上一级处理
	checkStatue := func(done <-chan interface{}, urls ...string) <-chan *http.Response {
		responses := make(chan *http.Response)
		go func() {
			defer close(responses)
			for _, url := range urls {
				resp, err := http.Get(url)
				if err != nil {
					fmt.Println(err)
					continue
				}
				select {
				case <-done:
					return
				case responses <- resp:
				}
			}
		}()
		return responses
	}

	done := make(chan interface{})
	defer close(done)
	urls := []string{"https://www.google.com", "https://www.baidu.com"}
	for response := range checkStatue(done, urls...) {
		fmt.Printf("Response: %v\n", response.Status)
	}
}
