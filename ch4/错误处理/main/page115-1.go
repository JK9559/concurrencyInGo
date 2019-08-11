/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"net/http"
)

func main() {
	type Result struct {
		Error    error
		Response *http.Response
	}
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

	errCount := 0
	urls := []string{"https://www.baidu.com", "a", "b", "c", "d", "e"}
	for res := range checkStatus(done, urls...) {
		if res.Error != nil {
			fmt.Printf("error: %v\n", res.Error)
			errCount++
			if errCount >= 3 {
				fmt.Println("Too many Err breaking!")
				break
			}
			continue
		}
		fmt.Printf("resp: %v\n", res.Response.Status)
	}
}
