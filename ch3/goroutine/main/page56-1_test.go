package main

import (
	"sync"
	"testing"
)

/*
使用一个cpu情况下
command : go test -bench=. -cpu=1 D:\goWorkSpace\Go\src\concurrencyInGo\ch3\goroutine\main\page56-1_test.go

D:\goWorkSpace\Go\src\concurrencyInGo>go test -bench=. -cpu=1 D:\goWorkSpace\Go\src\concurrencyInGo\ch3\goroutine\main\page56-1_test.go
goos: windows
goarch: amd64
BenchmarkContextSwitch   5000000               315 ns/op
PASS
ok      command-line-arguments  2.594s
*/
/**
切换上下文时间测试
*/
func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})

	var token struct{}
	sender := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			c <- token
		}
	}
	receiver := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			<-c
		}
	}

	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer()
	close(begin)
	wg.Wait()
}
