/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b int
	c float32
	int
	innerS
}

func main() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 1
	outer.in1 = 2
	outer.in2 = 3
	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)

	outer2 := &outerS{
		b:      6,
		c:      8.5,
		int:    11,
		innerS: innerS{5, 10},
	}

	fmt.Println(*outer2)
}
