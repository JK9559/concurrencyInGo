/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {
	// slice print
	a := []int{1, 2, 3}
	for k, v := range a {
		if k == 0 {
			a[0], a[1] = 100, 200
			fmt.Print(a)
		}
		a[k] = 100 + v
	}
	// [100 200 3][101 300 103]
	fmt.Print(a)

	fmt.Println()

	// array
	b := [3]int{1, 2, 3}
	for k, v := range b {
		if k == 0 {
			b[0], b[1] = 100, 200
			fmt.Print(b)
		}
		b[k] = 100 + v
	}
	// [100 200 3][101 102 103]
	fmt.Print(b)
}
