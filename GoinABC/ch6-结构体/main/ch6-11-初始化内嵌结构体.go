/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

type Wheel struct {
	Size int
}

type Engine struct {
	Power int
	Type  string
}

type Car struct {
	Wheel
	Engine
}

type Car1 struct {
	Wheel
	Engine1 struct {
		Power int
		Type  string
	}
}

func main() {
	// 结构体内嵌初始化，将结构体内嵌的类型作为字段名像普通结构体一样进行初始化
	c := Car{
		Wheel: Wheel{
			Size: 20,
		},
		Engine: Engine{
			Power: 100,
			Type:  "V1.5",
		},
	}

	fmt.Printf("%v\n", c)

	// 有些时候 可能会将结构体直接定义在嵌入的结构体中
	c1 := Car1{
		Wheel: Wheel{
			Size: 100,
		},
		Engine1: struct {
			Power int
			Type  string
		}{
			Power: 100,
			Type:  "V1.6",
		},
	}

	fmt.Printf("%v\n", c1)
}
