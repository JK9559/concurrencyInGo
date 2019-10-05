/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Skill struct {
		Name  string
		Level int
	}

	type Actor struct {
		Name   string
		Age    int
		Skills []Skill
	}

	a := Actor{
		Name: "James",
		Age:  12,
		Skills: []Skill{
			{Name: "Playing", Level: 1},
			{Name: "Playing1", Level: 1},
			{Name: "Playing2", Level: 1},
		},
	}

	if res, err := json.Marshal(a); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println(string(res))
	}
}
