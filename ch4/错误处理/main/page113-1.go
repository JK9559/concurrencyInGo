/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"net/http"
)

// 这里将错误包装起来 在函数返回时对错误的处理可以更加灵活
type Result struct {
	Error    error
	Response *http.Response
}

func main() {
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
		results := make(chan Result)
		go func() {
			defer close(results)
			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				result = Result{Error: err, Response: resp}
				select {
				case <-done:
					return
				case results <- result:
				}
			}
		}()
		return results
	}

	done := make(chan interface{})
	defer close(done)
	urls := []string{"https://www.google.com", "https://www.baidu.com"}
	for res := range checkStatus(done, urls...) {
		if res.Error != nil {
			fmt.Printf("error: %v\n", res.Error)
			continue
		}
		fmt.Printf("Response: %v\n", res.Response.Status)
	}

}
