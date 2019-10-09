/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

type dic struct {
	data map[interface{}]interface{}
}

func newDic() *dic {
	d := &dic{}
	d.Clear()
	return d
}

func (d *dic) Get(key interface{}) interface{} {
	return d.data[key]
}

func (d *dic) Set(key, value interface{}) {
	d.data[key] = value
}

func (d *dic) Visit(callback func(k, v interface{}) bool) {
	if callback == nil {
		return
	}
	for k, v := range d.data {
		if !callback(k, v) {
			return
		}
	}
}

func (d *dic) Clear() {
	d.data = make(map[interface{}]interface{})
}

func main() {
	dict := newDic()
	dict.Set("My abc", 40)
	dict.Set("Ohhh", 50)
	dict.Set("Emmm", 60)

	a := dict.Get("hah")
	fmt.Println(a)

	dict.Visit(func(k, v interface{}) bool {
		fmt.Printf("%v %v\n", k.(string), v.(int))
		return true
	})
}
