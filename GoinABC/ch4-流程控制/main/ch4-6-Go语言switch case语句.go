/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {
	// Go switch-case的表达式不需要为常量 整数 是从上到下进行求值
	// 可以一分支 多值
	var a = "a"
	switch a {
	case "a", "b", "c":
		fmt.Println("yes")
	}
	// 可以表达式
	var r = 11
	switch {
	case r > 10 && r < 20:
		fmt.Println("yes")
	}
	// 类型switch
	var t interface{}
	var p *int
	t = p
	fmt.Println(t == nil)
	switch t := t.(type) {
	case bool:
		fmt.Printf("%T\n", t)
	case int:
		fmt.Printf("%T\n", t)
	case *int:
		fmt.Printf("%T\n", t)
	default:
		fmt.Printf("%T\n", t)
	}
}
