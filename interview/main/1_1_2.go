/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

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
