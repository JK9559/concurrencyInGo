/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

// Go允许用户拓展或者修改已经有的类型的行为。
// 有利于代码的复用，也有利于修改已有类型来符合新的类型。
// 是通过嵌入类型来完成的。嵌入类型是将已有的类型直接声明在新的结构体中。
// 被嵌入的类型 被称为 新的外部类型的内部类型

type user struct {
	name  string
	email string
}

// 使用user类型的指针接收者声明为notify的方法
func (u *user) notify() {
	fmt.Printf("sending user email to %s<%s>\n",
		u.name,
		u.email)
}

type admin struct {
	// 嵌入类型
	user
	level string
}

func main() {
	ad := admin{
		user: user{
			name:  "xiaoming",
			email: "ming@abc.com",
		},
		level: "super",
	}

	// 可以直接访问内部类型的方法
	ad.user.notify()

	// 内部类型的方法也被提升到外部类型
	ad.notify()

	// 将user嵌入了admin，我们可以说 user是外部类型admin的内部类型
	//
}
