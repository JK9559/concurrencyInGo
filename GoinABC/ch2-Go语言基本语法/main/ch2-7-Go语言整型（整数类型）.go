/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

func main() {
	// Go 语言的整型 包含了有符号和无符号类型
	// 有符号 int8 int16 int32 int64 int
	// 无符号 uint8 uint16 uint32 uint64 uint
	// 两类的前四种分别对应了8 16 32 64bit大小的有符号/无符号整数
	// 最后两种 int uint 是对应特定CPU平台机器字大小的 32或者64bit

	// 多数情况 只需要使用int一种即可，处理速度也是最快的

	// rune 和 int32 等价
	// byte 和 int8 等价
	// uintptr 一种无符号的整数类型 没有指定具体的bit大小 但是足以容纳指针
	// 只有在底层编程才需要

	// 另外 即使int类型的大小是32bit 但是在需要将int当做int32类型的地方
	// 也需要一个现实类型的转换

	// 一个n bit的有符号数的值域为 -2^(n-1)~ 2^(n-1) - 1
	// 一个n bit的无符号数的值域为 0~ 2^n - 1

	// 在逻辑处理中 对于整型的长度没有要求
	// 但是针对二进制传输 和读写文件的结构描述时，为了保证文件结构不会受到不同的编译目标平台的
	// 字节长度影响，不要使用int和uint
}
