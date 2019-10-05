/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

// 事件注册
// 事件系统需要对外提供一个注册入口。这个入口传入注册的事件名和对应事件名称的响应函数
// 事件注册的过程就是将事件名和响应函数关联并且保存起来

var eventByName = make(map[string][]func(interface{}))

// 注册事件 提供事件名和回调函数
func RegisterEvent(name string, callback func(interface{})) {
	// 通过名字查找事件列表
	list := eventByName[name]

	// 在列表切片中添加函数
	list = append(list, callback)

	// 保存新的事件列表
	eventByName[name] = list
}

// 调用事件
func CallEvent(name string, param interface{}) {
	// 通过名字查找事件列表
	list := eventByName[name]

	// 遍历该事件的所有回调
	for _, callback := range list {
		// 传入参数调用回调
		callback(param)
	}
}

type Actor struct {
}

func (a *Actor) OnEvent(param interface{}) {
	fmt.Println("actor event:", param)
}

func GlobalEvent(param interface{}) {
	fmt.Println("global event:", param)
}

func main() {
	a := new(Actor)

	RegisterEvent("OnSkill", a.OnEvent)

	RegisterEvent("OnSkill", GlobalEvent)

	CallEvent("OnSkill", 100)
}
