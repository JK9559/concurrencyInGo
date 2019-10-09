/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

// 类型分支判断基本类型
func printType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println(v, "is int")
	case string:
		fmt.Println(v, "is string")
	case bool:
		fmt.Println(v, "is bool")
	default:
		// 这里没default的情况下 传入未枚举的类型 不报错
		fmt.Println("i do not know")
	}
}

// 使用类型分支判断接口类型
type elePay struct {
}

func (e *elePay) CanUseFaceID() {
}

type cash struct {
}

func (c *cash) Stolen() {
}

type canUseFaceID interface {
	CanUseFaceID()
}

type canBeStolen interface {
	Stolen()
}

func print(payMethod interface{}) {
	switch payMethod.(type) {
	case canUseFaceID:
		fmt.Printf("%T\n", payMethod)
	case canBeStolen:
		fmt.Printf("%T\n", payMethod)
	}
}

func main() {
	// type-switch 语法
	// switch 接口变量.(type) {
	//    case 类型1:
	//        // 变量是类型1时的处理
	//    case 类型2:
	//        // 变量是类型2时的处理
	//    …
	//    default:
	//        // 变量不是所有case中列举的类型时的处理
	// }
	printType(102)
	printType("ha")
	printType(true)
	printType(1.2)

	print(new(elePay))
	print(new(cash))
}
