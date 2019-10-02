/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"sync"
)

func main() {
	// 使用了sync.Map
	// 特性：
	// 1. 无需初始化 直接声明即可
	var synMap sync.Map
	// 2. sync.Map 使用不同于map,不能使用map的方式进行取值和设置操作。
	// 需要使用sync.Map的方法进行调用

	// 将键值对保存到sync.Map
	synMap.Store("greece", 97)
	synMap.Store("london", 99)
	synMap.Store("egypt", 100)

	// 从sync.Map中取值
	fmt.Println(synMap.Load("london"))

	// 从sync.Map中删除数据
	synMap.Delete("london")

	// 遍历sync.Map的键值对
	synMap.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})

	// sync.Map 没有提供获取 map 数量的方法，
	// 替代方法是获取时遍历自行计算数量。
	// sync.Map 为了保证并发安全有一些性能损失，
	// 因此在非并发情况下，使用 map 相比使用 sync.Map 会有更好的性能。

}
