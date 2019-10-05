/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"sort"
)

func main() {
	// 使用sort排序时 可以使用sort.Interface排序
	// 需要实现 Len Less Swap 三个方法

	// 其他 也可以使用以下已经提供的方法
	names := sort.StringSlice{
		"abc",
		"a123",
		"edf",
	}
	sort.Sort(names)
	fmt.Println(names)

	names1 := []string{
		"abc",
		"a123",
		"edf",
	}
	sort.Strings(names1)
	fmt.Println(names1)

	ints := sort.IntSlice{
		1, 33, 45,
	}
	sort.Sort(ints)
	fmt.Println(ints)
}
