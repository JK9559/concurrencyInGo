/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {
	// Go语言有两种复数类型 complex64(32位实数和虚数) complex128(64位实数和虚数)
	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(x * y)
	fmt.Println(real(x * y))
	fmt.Println(imag(x * y))
}
