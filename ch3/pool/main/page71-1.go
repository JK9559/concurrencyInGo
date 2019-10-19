/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"sync"
)

/**
关于Pool模式 是一种创建和提供可供固定数量实例或者Pool实例的方法，用于创建昂贵的场景。
以便只创建固定数量的实例。
*/
func main() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating a new instance")
			return struct{}{}
		},
	}
	// 当调用Get时，先检查池中是否有可用的实例返回给调用者，如果没有调用New方法创建一个新的实例。
	myPool.Get()
	instance := myPool.Get()
	// 当调用完成时，调用者调用Put方法吧工作的实例归还给池中，以供其他进程使用
	myPool.Put(instance)
	myPool.Get()

	/*
		Creating a new instance
		Creating a new instance
	*/
}
