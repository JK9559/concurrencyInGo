/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"reflect"
)

type T struct {
}

type Error struct {
	errCode uint8
}

func (e *Error) Error() string {
	switch e.errCode {
	case 1:
		return "1"
	case 2:
		return "2"
	case 3:
		return "3"
	default:
		return "4"

	}
}

func main() {

	// Go中 指针 函数 interface slice channel map 的零值都是nil
	// interface
	// 一个interface在没有进行初始化的时候 对应的值是nil
	var v interface{}
	fmt.Println(v, &v)

	var k *T
	var i interface{}
	i = k
	fmt.Println(i)

	var e *Error
	fmt.Println(reflect.TypeOf(e))
	fmt.Println(e == nil)
	checkError(e)

}

func checkError(err error) {
	fmt.Println(reflect.TypeOf(err))
	fmt.Println(err == (*Error)(nil))
	// 这里因为传入的参数为*Error类型 接收类型为error 所以内部的err为(*Error,nil) 即类型为*Error 值为nil
	// 和nil比较 必定不相等
	// 所以可以将nil转换为*Error 来进行比较
	if err != nil {
		panic(err)
	}
}
