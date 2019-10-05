/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"sort"
	"strings"
)

var original = []string{
	"Nonmetals",
	"        Ruropium",
	"    Hydrogen",
	"    Carbon",
	"    Nitrogen",
	"    Oxygen",
	"Inner Transitionals",
	"    Lanthanides",
	"        Europium",
	"        Cerium",
	"    Actinides",
	"        Uranium",
	"        Plutonium",
	"        Curium",
	"Alkali Metals",
	"    Lithium",
	"    Sodium",
	"    Potassium",
}

type Entry struct {
	key      string
	value    string
	children Entries
}

// golang中的函数 全都是值传递的，只不过区别在于 如果传了指针 就可以对原值进行修改，也就相当于引用传递，另外传递指针可以节省值拷贝的空间
// 值传递为 函数内部修改的入参 实际上是外部调用方入参的拷贝

type Entries []Entry

// 想要使用sort.Sort() 必须要实现以下三个方法 Len Less Swap
func (entries Entries) Len() int {
	return len(entries)
}

func (entries Entries) Less(i, j int) bool {
	return entries[i].key < entries[j].key
}

func (entries Entries) Swap(i, j int) {
	entries[i], entries[j] = entries[j], entries[i]
}

func main() {
	fmt.Println("|     Original      |       Sorted      |")
	fmt.Println("|-------------------|-------------------|")
	sorted := SortedIndentedStrings(original)
	for i := range original {
		fmt.Printf("|%-19s|%-19s|\n", original[i], sorted[i])
	}
}

func SortedIndentedStrings(slice []string) []string {
	entries := populateEntries(slice)
	fmt.Println(entries)
	return sortedEntries(entries)
}

func populateEntries(slice []string) Entries {
	// 获取传入原始数据缩进的宽度和单位字符串
	indent, indentSize := computeIndent(slice)
	// 初始化一个结果集
	entries := make(Entries, 0)
	// 遍历每个字符串
	for _, item := range slice {
		// 初始化指向每个字符串头部指针i为0 level为第0级
		i, level := 0, 0
		for strings.HasPrefix(item[i:], indent) {
			// 后移指针1个indentSize
			i += indentSize
			// level加一
			level++
		}
		// 将key值统一化处理
		key := strings.ToLower(strings.TrimSpace(item))
		addEntry(level, key, item, &entries)
	}
	return entries
}

func computeIndent(slice []string) (string, int) {
	// 遍历切片的每个字符串
	for _, item := range slice {
		// 如果字符串是以' '或者'\t'开头的
		if len(item) > 0 && (item[0] == ' ' || item[0] == '\t') {
			whiteSpace := rune(item[0])
			for i, char := range item[1:] {
				if char != whiteSpace {
					i++
					return strings.Repeat(string(whiteSpace), i), i
				}
			}
		}
	}
	return "", 0
}

func addEntry(level int, key, value string, entries *Entries) {
	// 如果为第0级 在当前级别新增一个节点 因为如果为递归，当前级别为上一级别的children
	if level == 0 {
		*entries = append(*entries, Entry{
			key:      key,
			value:    value,
			children: make(Entries, 0),
		})
		// 否则级别减去1 继续寻找上一级别
	} else {
		addEntry(level-1, key, value, &((*entries)[entries.Len()-1].children))
	}
}

func sortedEntries(entries Entries) []string {
	var indentedSlice []string
	sort.Sort(entries)
	for _, entry := range entries {
		// 递归进行排序 注意是传地址 如果不传地址 可能造成本身的内容未被更改
		populateIndentedStrings(entry, &indentedSlice)
	}
	return indentedSlice
}

func populateIndentedStrings(entry Entry, indentedSlice *[]string) {
	*indentedSlice = append(*indentedSlice, entry.value)
	sort.Sort(entry.children)
	for _, child := range entry.children {
		populateIndentedStrings(child, indentedSlice)
	}
}
