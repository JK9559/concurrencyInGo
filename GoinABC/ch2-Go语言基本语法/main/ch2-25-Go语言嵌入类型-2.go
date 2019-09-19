/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

type notifier interface {
	notify1()
}

type user1 struct {
	name  string
	email string
}

// 因为这里是指针类型 所以下面的调用使用指针调用
func (u *user1) notify1() {
	fmt.Printf("sending to %s <%s>\n", u.name, u.email)
}

type admin1 struct {
	user1
	level string
}

func main() {
	ad := admin1{
		user1: user1{
			name:  "xiaoli",
			email: "xiaoli@abc.com",
		},
		level: "super",
	}

	// 这里可以看到使用了admin类型，但是代码中可以看到admin并没有实现notifier接口
	// 这是因为 内部类型的提升，内部类型实现的接口会自动提升到外部类型。
	// 也就是意味着，因为内部类型的实现 外部类型也同样实现了这个接口
	sendNotification(&ad)

}

func sendNotification(n notifier) {
	n.notify1()
}
