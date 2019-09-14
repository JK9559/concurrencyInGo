/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {
	// 字符串的每个元素被称为字符，但是字符并不是Go的一个类型
	// 字符只是整数的特殊用例

	// Go的字符有两种：
	// 1. uint8类型 或者byte类型，代表了ASCII码的一个字符
	var ch1 byte = 'A'
	var ch2 byte = 65
	var ch3 byte = '\x41'
	var ch4 byte = '\377'
	fmt.Println(ch1, ch2, ch3, ch4)
	// 2. uint32类型 或者rune类型，代表了utf-8字符(Unicode)。当处理中文、日文等其他复合字符时
	// 需要使用rune类型
	// 因为Unicode至少会占用2个字节(16bit) 所以使用int16或者int表示
	// 如果需要使用4个字节 会加上\U前缀紧跟着8位长度的十六进制数字(8*4=32)
	// 如果使用2字节 会加上\u前缀紧跟着4位长度的十六进制数字(4*4=16)

	var chh1 int = '\u0041'
	var chh2 int = '\u03B2'
	var chh3 int = '\U00101234'
	// 十进制int表示
	fmt.Printf("%d - %d - %d\n", chh1, chh2, chh3)
	// 字符形式
	fmt.Printf("%c - %c - %c\n", chh1, chh2, chh3)
	// 大写十六进制
	fmt.Printf("%X - %x - %X\n", chh1, chh2, chh3)
	// Unicode代码点表示格式为 U+hhhh
	fmt.Printf("%U - %U - %U\n", chh1, chh2, chh3)
}
