/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"runtime"
)

func main() {
	status := runtime.MemStats{}
	runtime.ReadMemStats(&status)
	// 获取当前的内存状态
	fmt.Printf("%dkb\n", status.Alloc/1024)
}
