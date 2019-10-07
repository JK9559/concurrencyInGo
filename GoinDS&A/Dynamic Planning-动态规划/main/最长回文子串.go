/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func longestPalindrome(s string) string {
	if 0 == len(s) {
		return ""
	}
	if 1 == len(s) {
		return s
	}
	var resi, resj, res = 0, 0, 0
	var strLen = len(s)
	var dp = make([][]int, strLen)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, strLen)
		dp[i][i] = 1
	}

	for i := strLen - 1; i >= 0; i-- {
		for j := i + 1; j < strLen; j++ {
			if (dp[i+1][j-1] == 1 && s[i] == s[j]) || (j == i+1 && s[i+1] == s[i]) {
				dp[i][j] = 1
				if res < j-i+1 {
					resi = i
					resj = j
					res = resj - resi + 1
				}
			}
		}
	}
	fmt.Println(dp)
	return s[resi : resj+1]

}

func main() {
	var str = "aaaa"
	res := longestPalindrome(str)
	fmt.Println(res)

}
