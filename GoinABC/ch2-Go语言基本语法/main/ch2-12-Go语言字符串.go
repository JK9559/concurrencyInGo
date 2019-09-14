/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {
	// Go对于字符串支持两种形式的字面值
	// 解释字符串 使用双引号括起来，其中相关的转义字符被替换，也是平时使用很多的一种
	str1 := "abcc\nccdd\thaha"
	fmt.Println(str1)

	// 非解释字符串 使用反引号`括起来 支持换行
	str2 := `abcc\ncc\
			rdd\thaha`
	fmt.Println(str2)
}
