/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

type People struct {
	name  string
	child *People
}

type Address struct {
	Province    string
	City        string
	ZipCode     int
	PhoneNumber string
}

func printMsg(msg *struct {
	id   int
	data string
}) {
	fmt.Printf("%T\n", msg)
}

// 重点重点： 结构体成员中只能包含结构体的指针类型，包含非指针类型会引起编译错误。

func main() {
	// 有两种初始化形式 "键值对"形式 和 多值列表的形式
	// 1. 键值对形式 键值对的填充是可选的，不需要初始化的字段可以不填入初始化列表中。
	relation := &People{
		name: "grandpa",
		child: &People{
			name: "father",
			child: &People{
				name:  "me",
				child: nil,
			},
		},
	}

	fmt.Println(relation)

	// 2. 使用多个值的列表初始化结构体
	// 注意点: 1. 必须初始化所有字段 2. 填充顺序必须与声明顺序一直 3. 键值对形式和值列表形式不能混用
	addr := &Address{
		"四川",
		"成都",
		100001,
		"00",
	}
	fmt.Println(addr)

	// 额外话题：初始化匿名结构体
	// 匿名结构体没有type定义 如
	struct1 := struct {
		Str string
		X   int
	}{
		Str: "yes",
		X:   0,
	}
	fmt.Println(struct1)

	// 接收匿名结构体的函数
	// 这里 注意 使用了取地址符
	msg := &struct {
		id   int
		data string
	}{
		id:   1024,
		data: "hello",
	}

	printMsg(msg)

}
