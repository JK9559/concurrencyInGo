/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {
	mp1 := make(map[string]int)

	mp1["abc"] = 1
	mp1["cde"] = 2
	mp1["fgh"] = 3

	delete(mp1, "abc")

	for l, v := range mp1 {
		fmt.Println(l, v)
	}
}
