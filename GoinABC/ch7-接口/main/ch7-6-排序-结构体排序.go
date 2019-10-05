/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"sort"
)

type HeroKind int

const (
	None = iota
	Tank
	Assassin
	Bala
)

type Hero struct {
	name string
	kind HeroKind
}

type Heroes []Hero

func (h Heroes) Len() int {
	return len(h)
}

func (h Heroes) Less(i, j int) bool {
	if h[i].kind == h[j].kind {
		return h[i].name < h[j].name
	} else {
		return h[i].kind < h[j].kind
	}
}

func (h Heroes) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func main() {
	// 对结构体类型排序
	// 注意 最终排序的是slice 所以实现接口的类型 应该是slice
	var hs = Heroes{
		Hero{
			name: "abc",
			kind: None,
		},
		Hero{
			name: "bcd",
			kind: None,
		},
		Hero{
			name: "abc",
			kind: Tank,
		},
	}

	sort.Sort(hs)

	for i := 0; i < len(hs); i++ {
		fmt.Println(hs[i])
	}

	// 使用sort.Slice排序
	heros := []*Hero{
		{
			name: "Obaha",
			kind: Tank,
		},
		{
			name: "abc",
			kind: None,
		},
		{
			name: "bcd",
			kind: None,
		},
	}

	sort.Slice(heros, func(i, j int) bool {
		if heros[i].kind == heros[j].kind {
			return heros[i].name < heros[j].name
		} else {
			return heros[i].kind < heros[j].kind
		}
	})

	for i := 0; i < len(heros); i++ {
		fmt.Println(*heros[i])
	}
}
