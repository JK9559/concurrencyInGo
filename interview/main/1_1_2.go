/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

/*
给定根号2约等于1.414，计算根号2小数点后10位的精确值
*/
func main() {
	low := 1.4
	high := 1.5
	const Accuracy = 0.00000000001
	mid := (high - low) / 2.0
	for high-mid > Accuracy {
		if mid*mid > 2 {
			high = mid
		} else {
			low = mid
		}
		mid = (low + high) / 2.0
	}

	fmt.Println(mid)

}
