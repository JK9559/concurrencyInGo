/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

type interA interface {
	inA1(int) int
	inA2(int) int
}

// 多个类型能实现相同的接口
type interB interface {
	inB1(float32) float32
}

// 一个类型能实现多个接口
type structA struct {
}

func (s *structA) inA1(a int) int {
	return a
}

func (s *structA) inA2(a int) int {
	return a
}

func (s *structA) inB1(b float32) float32 {
	return b
}

type structB struct {
}

func (s *structB) inB1(b float32) float32 {
	return b
}

func main() {

}
