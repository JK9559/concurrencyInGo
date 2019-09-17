/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"time"
)

type My = time.Duration

// 这里会报错 提示了在一个非本地类型上定义新方法，time.Duration是在time包定义的，在main包中使用
// 不在同一个包中 所以不能给不在一个包中的类型定义方法
// 解决方法 1 将上面的类型别名改为类型定义 2 将main包中的别名添加到time包中
func (m My) hhha(a string) {
	fmt.Println("a")
}

func main() {
	// 类型别名和类型定义
	// 下面是类型定义
	type NewInt int
	// 下面是类型别名
	type IntAlias = int

	var a NewInt = 2
	var b IntAlias = 3
	fmt.Printf("%T %T\n", a, b)
	// 通过上面代码可以看出区别
	// 类型定义 表示了main包下定义了类型NewInt 是实际存在的新的类型 具有int的特性
	// 类型别名 表示了IntAlias这个类型只会在代码文本中出现 编译完成时只有int没有IntAlias

}
