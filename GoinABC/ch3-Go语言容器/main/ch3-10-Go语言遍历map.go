/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {
	scene := make(map[string]int)

	scene["route"] = 66
	scene["brazil"] = 4
	scene["China"] = 960

	for k, v := range scene {
		fmt.Println(k, v)
	}

}
