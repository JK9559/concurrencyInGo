/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

type Flying struct {
}

func (f *Flying) Fly() {
	fmt.Println("can fly")
}

type Walkable struct {
}

func (w *Walkable) Walk() {
	fmt.Println("can walk")
}

type Human struct {
	Walkable
}

type Bird struct {
	Walkable
	Flying
}

func main() {
	b := new(Bird)
	// 等效的
	b.Fly()
	b.Flying.Fly()
	b.Walk()
	b.Walkable.Walk()

	h := new(Human)
	h.Walkable.Walk()
	h.Walk()
}
