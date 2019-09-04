/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

/*
以下两段代码说明了两个问题：
1. range函数是值传递 也就是说 得到的k, v都是无法在循环中直接修改的 也就是说虽然在k==0时都修改了array或者slice 但是此时的v都是1
2. 数组是值传递的，类似于在函数里给range入参了一个数组。参考 https://blog.csdn.net/bobodem/article/details/80187558
3. slice虽然改变了 但是不能严格说是引用传递。因为slice存在一个底层数组 外部表现可能未修改，但是底层数组可能已经修改了
参考 https://blog.csdn.net/bobodem/article/details/80188126
*/
func main() {
	// slice print
	a := []int{1, 2, 3}
	for k, v := range a {
		if k == 0 {
			a[0], a[1] = 100, 200
			fmt.Print(a)
		}
		a[k] = 100 + v
	}
	// [100 200 3][101 300 103]
	fmt.Print(a)

	fmt.Println()

	// array
	b := [3]int{1, 2, 3}
	for k, v := range b {
		if k == 0 {
			b[0], b[1] = 100, 200
			fmt.Print(b)
		}
		b[k] = 100 + v
	}
	// [100 200 3][101 102 103]
	fmt.Print(b)
}
