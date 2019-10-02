/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"sort"
)

type statistics struct {
	numbers []float64
	mean    float64
	median  float64
}

func getStats(numbers []float64) (stats statistics) {
	stats.numbers = numbers
	// 对原数组进行了升序排序 对参数进行了更改
	sort.Float64s(stats.numbers)
	stats.mean = sum(numbers) / float64(len(numbers))
	stats.median = median(numbers)
	return stats
}

func sum(numbers []float64) (totals float64) {
	for _, x := range numbers {
		totals += x
	}
	return totals
}

func median(numbers []float64) float64 {
	middle := len(numbers) / 2
	result := numbers[middle]
	if len(numbers)%2 == 0 {
		result = (result + numbers[middle-1]) / 2
	}
	return result
}

func main() {
	var num = []float64{1.0, 1.9, 2.5, 0.1, 0.3}
	var stats = getStats(num)
	fmt.Println(stats.numbers)
	fmt.Println(stats.mean)
	fmt.Println(stats.median)
}
