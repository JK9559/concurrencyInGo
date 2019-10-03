/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"runtime"
)

type panicContext struct {
	function string
}

// 函数以保护方式运行
// 调用 recover 可以捕获到 panic 的输入值
func ProtectRun(entry func()) {
	defer func() {
		err := recover()

		switch err.(type) {
		case runtime.Error:
			fmt.Println("runtime.error:", err)
		default:
			fmt.Println("error:", err)
		}

	}()
	entry()
}

func main() {
	// recover函数可以让进去宕机流程panic的goroutine恢复过来
	// 仅在defer中生效
	fmt.Println("运行前")

	ProtectRun(func() {
		fmt.Println("手动宕机前")
		panic(&panicContext{function: "手动触发panic"})
		fmt.Println("手动宕机后")
	})

	ProtectRun(func() {
		fmt.Println("赋值宕机前")
		var a *int
		*a = 1
		fmt.Println("赋值宕机后")
	})
	fmt.Println("运行后")
}

// panic 和 recover的关系
// 有panic无recover 宕机
// 有panic有recover 不会宕机 执行完对应的defer后，从宕机点退出当前函数后 继续执行
