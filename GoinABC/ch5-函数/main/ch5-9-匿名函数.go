/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"flag"
	"fmt"
)

func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}

var skillParam = flag.String("skill", "", "skill to perform")

func main() {
	// 匿名函数就是没有定义名字的普通函数
	// 可以在定义时调用匿名函数
	func(data int) {
		fmt.Println("it is ", data)
	}(100)

	// 将匿名函数赋值给变量
	funNm := func(data int) {
		fmt.Println("ha ", data)
	}
	funNm(200)

	// 匿名函数用作回调函数
	visit([]int{1, 2, 3, 4}, func(i int) {
		fmt.Println(i)
	})

	// 使用匿名函数实现操作封装
	flag.Parse()

	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angle fly")
		},
	}

	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		fmt.Println("skill not found")
	}
}

// go run ch5-9-匿名函数.go --skill=fly
