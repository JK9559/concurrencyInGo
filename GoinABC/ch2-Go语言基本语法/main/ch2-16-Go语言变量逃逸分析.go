/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {
	// 堆和栈
	// 栈可以用于内存分配，栈的分配和回收速度都非常快
	// 堆分配内存时，需要先找到一块足够的内存，再存放变量，可能导致的问题就是，
	// 虽然有足够的空间 但是各个空间不连续，无法分配

	// 对比C/C++，需要自己学习如何进行内存分配，例如局部变量使用栈，全局变量 结构体类型使用堆分配
	// Go将整个过程整合到编译器中，命名为"变量逃逸分析"

	// go run -gcflags "-m -l" GoinABC\ch2-Go语言基本语法\main\ch2-16-Go语言变量逃逸分析.go
	// 执行上述命令 -gcflags 参数是编译参数 -m 表示进行内存分配分析 -l 表示避免程序内联，也就是避免进行程序优化。
	dummy := func(b int) int {
		var c int
		c = b
		return c
	}

	void := func() {

	}

	var a int

	void()

	fmt.Println(a, dummy(0))

	// 对于变量逃逸的总结：
	// 1. 如上文所说 变量在栈上分配比在堆上分配效率更高
	// 2. 栈上分配的数据不需要GC来处理
	// 3. 但是在堆上的数据需要GC来处理
	// 4. 逃逸分析本身就是决定变量在哪里分配
	// 5. 变量逃逸分析发生在编译阶段

	// 更多实例：
	// todo 学习更多关于变量逃逸的实例
}
