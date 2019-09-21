/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {
	// Go没有针对切片元素提供专门的接口或者语法去做删除动作。
	// 有三种情况 1从头删除 2从尾删除 3从中间位置删除，其中从尾部最快
	// 1 从头部
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("%p\n", a)
	fmt.Println(a)
	i := 1
	// 从头部删除i个元素
	// 删头法1 不会导致内存结构发生变化 可通过观察地址得到
	a = append(a[:0], a[i:]...)
	fmt.Printf("%p\n", a)
	fmt.Println(a)

	// 删头法2 直接移动了头部指针 导致内存结构发生变化
	a = a[i:]
	fmt.Printf("%p\n", a)
	fmt.Println(a)

	// 删头法3 利用copy函数返回值为两切片较小的一个的长度 并且是后一个值拷贝到前一个(地址不变化)
	// 所以这个方法 内存结构也不会发生改变
	a = a[:copy(a, a[i:])]
	// 那么为什么不直接copy而是 另外再截取一段 可以运行下面的注释
	// copy(a, a[i:])
	fmt.Printf("%p\n", a)
	fmt.Println(a)

	// 从尾部删除i个元素
	a = a[:len(a)-1]
	fmt.Printf("%p\n", a)
	fmt.Println(a)

	// 从中间删除i位置开始(包括第i个位置) 删除N个元素
	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("%p\n", b)
	fmt.Println(b)
	i = 3
	N := 2
	b = append(b[:i], b[i+N:]...)
	fmt.Printf("%p\n", b)
	fmt.Println(b)
	b = b[:i+copy(b[i:], b[i+N:])]
	fmt.Printf("%p\n", b)
	fmt.Println(b)

	// 总结 可以看到 上面删除几乎都用到了append 所以 删除本质就是
	// 将切片先按照规则切开 然后再拼回去
}
