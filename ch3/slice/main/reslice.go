/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

func main() {
	// 通过 截取 创建slice。基于已有的slice创建slice对象，被称为reslice。关键点在于
	// 新的slice和老的slice共用底层数组，新老slice对底层数组的更改会影响到彼此。
	// 基于数组(老)创建的slice(新)，对数组或者slice元素的更改会影响到彼此。

	// 新老slice或者老数组新slice相互影响的前提是两者共用底层数组，如果因为执行了append操作
	// 而使新的slice底层数组扩容，那么新slice的底层数组就移动到了新的位置，两者之间就不会继续影响了

	// 截取方式如下：
	// data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// slice := data[low, high, max]

	// 其中data可以是数组 可以是slice 截取元素为 [low, high) 之间的元素，最大容量max也是开区间，
	// 默认的容量为取到底层数组结尾，max >= high >= low. 并且high max必须在老数组或者老slice容量内

	// 例子如下：
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// s1 是从 slice的 2(包括)到5(不包括) 长度为3 容量未设置 默认到内置数组结尾，容量为8
	s1 := slice[2:5]
	// s2 是从 s1的 2(包括)到6(不包括) 长度为4 容量到7(不包括) 容量为5
	s2 := s1[2:6:7]
	fmt.Println(">============<")
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("slice:", slice)

	// s2 后新增元素100 容量够 不需要扩容，修改原底层数组，影响了slice和s1，使他们对应的底层数组都发生了变更
	// 可以看到输出结果在slice体现了，虽然看不到s1的结果，但是对应的底层数组也发生了改动
	s2 = append(s2, 100)

	fmt.Println(">======[s2.append 100]======<")
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("slice:", slice)

	// s2 后新增元素200 容量不够 扩容，使用了新的底层数组，并且把容量翻倍变为10，这次修改是在新的底层数组修改
	// 不影响slice 和 s1
	s2 = append(s2, 200)

	fmt.Println(">======[s2.append 200]======<")
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("slice:", slice)

	// 修改s1位置2的元素 影响到了slice 但是对于s2没影响
	s1[2] = 20

	fmt.Println(">======[s1[2] = 20]======<")
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("slice:", slice)
}
