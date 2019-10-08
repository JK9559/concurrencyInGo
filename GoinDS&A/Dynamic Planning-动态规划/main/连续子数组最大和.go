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

func longestSubArrayMaxSum(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 0 {
		return 0
	}
	var resLen, resj, res = 1, 0, nums[0]

	var arrLen = len(nums)
	var dp []int = make([]int, arrLen)

	dp[0] = nums[0]
	for i := 1; i < arrLen; i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		if dp[i] > res {
			resLen++
			if nums[i] == dp[i] {
				resLen = 1
			}
			resj = i
			res = dp[i]
			//fmt.Println(i)
		}
	}

	fmt.Println(dp)
	fmt.Println(resj-resLen+1, resj, nums[resj-resLen+1:resj+1])
	return res

}

func main() {
	fmt.Println(longestSubArrayMaxSum([]int{1, -2, 2, 3}))
}
