/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

type Cat struct {
	Color string
	Name  string
}

// 多种方式创建和初始化结构体——模拟构造函数重载
func NewCatByName(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

func NewCatByColor(color string) *Cat {
	return &Cat{
		Color: color,
	}
}

// 带有父子关系的结构体的构造和初始化——模拟父级构造调用
type BlackCat struct {
	// 相当于BlackCat的基类
	Cat
}

func NewCat(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

func NewBlackCat(color string) *BlackCat {
	cat := &BlackCat{}
	cat.Color = color
	return cat
}

func main() {
	// Go的类型和结构体没有构造函数的功能
	// 可以使用结构体初始化的过程来实现构造函数的功能

}
