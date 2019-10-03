/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"errors"
	"fmt"
)

// 自定义一个错误
var err = errors.New("this is an error")

func div(div1, div2 int) (int, error) {
	if div2 == 0 {
		return 0, err
	}
	return div1 / div2, nil
}

func main() {

	fmt.Println(div(1, 0))
}
