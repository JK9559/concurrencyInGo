/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

type A interface {
	A()
}

type B interface {
	B()
}

type bird struct {
}

type pig struct {
}

func (b *bird) A() {
	fmt.Println("this is A in bird")
}

func (b *bird) B() {
	fmt.Println("this is B in bird")
}

func (p *pig) A() {
	fmt.Println("this is A in pig")
}

func main() {
	// 将接口转换为其他接口
	// 当实现某个接口的类型同时实现了另外一个接口 此时可以再两个接口之间转换
	animals := make(map[string]interface{})

	animals["bird"] = new(bird)

	animals["pig"] = new(pig)

	for name, obj := range animals {
		// 使用类型断言获得a 类型为A 以及是否断言成功的判断
		a, isA := obj.(A)
		b, isB := obj.(B)

		fmt.Printf("name: %s isA: %v isB: %v %T %T\n",
			name, isA, isB, a, b)

		if isA {
			a.A()
		}

		if isB {
			b.B()
		}
	}

	// 将接口转换为其他类型
	p1 := new(pig)

	// pig实现了A接口 所以 可以被转换为A接口类型 保存在a1中
	var a1 A = p1
	// a1中保存的本身就是*pig,所以可以被转换为*pig类型
	p2 := a1.(*pig)

	fmt.Printf("p1=%p p2=%p", p1, p2)

	p3 := new(pig)
	var a2 A = p3

	p4 := a2.(*bird)
	// 这里报错
	fmt.Println(p4)
}
