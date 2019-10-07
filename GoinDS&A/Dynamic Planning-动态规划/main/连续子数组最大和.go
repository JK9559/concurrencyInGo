/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func longestSubArrayMaxSum(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	if len(arr) == 0 {
		return 0
	}
	var resLen, resj, res = 1, 0, arr[0]

	var arrLen = len(arr)
	var dp []int = make([]int, arrLen)

	dp[0] = arr[0]
	for i := 1; i < arrLen; i++ {
		dp[i] = max(arr[i], dp[i-1]+arr[i])
		if dp[i] > res {
			resLen++
			if arr[i] == dp[i] {
				resLen = 1
			}
			resj = i
			res = dp[i]
			//fmt.Println(i)
		}
	}

	fmt.Println(dp)
	fmt.Println(resj-resLen+1, resj, arr[resj-resLen+1:resj+1])
	return res

}

func main() {
	fmt.Println(longestSubArrayMaxSum([]int{1, -2, 2, 3}))
}
