/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func main() {
	//sl := make([]int, 2, 3)
	//for i := 0; i < len(sl); i++ {
	//	sl[i] = 1
	//	fmt.Println(sl)
	//	fmt.Println(len(sl))
	//	fmt.Println(cap(sl))
	//}
	//
	//sl = append(sl, 1)
	//fmt.Println(sl)
	//fmt.Println(len(sl))
	//fmt.Println(cap(sl))
	//sl = append(sl, 1)
	//fmt.Println(sl)
	//fmt.Println(len(sl))
	//fmt.Println(cap(sl))
	//sl = append(sl, 1)
	//fmt.Println(sl)
	//fmt.Println(len(sl))
	//fmt.Println(cap(sl))
	//sl = append(sl, 1)
	//fmt.Println(sl)
	//fmt.Println(len(sl))
	//fmt.Println(cap(sl))

	fmt.Println(">===========<")

	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	// 错误写法
	for _, stu := range stus {
		// 因为这里的stu是个临时变量 取这个变量的地址始终只有一个，不是实际slice的地址
		m[stu.Name] = &stu
	}

	for k, v := range m {
		println(k, "=>", v.Name)
	}

	// 正确
	for i := 0; i < len(stus); i++ {
		m[stus[i].Name] = &stus[i]
	}
	for k, v := range m {
		println(k, "=>", v.Name)
	}

}
