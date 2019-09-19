/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

type notifier1 interface {
	notify2()
}

type user2 struct {
	name  string
	email string
}

func (u *user2) notify2() {
	fmt.Printf("send %s<%s>\n", u.name, u.email)
}

type admin2 struct {
	user2
	level string
}

func (a *admin2) notify2() {
	fmt.Printf("send admin %s<%s>\n", a.name, a.email)
}

func main() {
	ad := admin2{
		user2: user2{
			name:  "xiaohu",
			email: "xiaohu@ab.com",
		},
		level: "super",
	}

	// 先外后内
	sendNotification1(&ad)

	// 先外后内
	ad.user2.notify2()

	// 先外后内
	ad.notify2()

}

func sendNotification1(n notifier1) {
	n.notify2()
}
