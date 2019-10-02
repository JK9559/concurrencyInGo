/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 列表是一种非连续存储的容器，由多个节点组成，节点通过一些遍历记录彼此之间的联系
	// 列表有多种实现方法 单链表 双链表等

	// Go中列表使用container/list包来实现 内部实现原理是双链表，
	// 能够高效的进行任意位置的元素插入和删除

	// 初始化列表
	// 方法一 通过container/list 包的New方法 初始化list
	mylist := list.New()
	// 方法二 通过声明初始化list
	var mylist1 list.List

	// 注意 上面两种声明方式 得到的list类型是不同的
	// *list.List list.List
	fmt.Printf("%T %T\n", mylist, mylist1)

	// 在列表中插入元素
	ele1 := mylist.PushBack("aha")
	ele2 := mylist1.PushFront(667)

	mylist.PushBackList(&mylist1)
	mylist1.PushFrontList(mylist)

	mylist.InsertAfter(6, ele1)
	mylist1.InsertBefore(777, ele2)

	// 遍历
	for i := mylist.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	fmt.Println()
	for i := mylist1.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	fmt.Println()
	mylist.Remove(ele1)
	mylist1.Remove(ele2)

	for i := mylist.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	fmt.Println()
	for i := mylist1.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

}
