/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"
	"testing"
	"time"
)

/**
command:
go test -benchtime=10s -bench=. D:\goWorkSpace\Go\src\concurrencyInGo\ch3\pool\main\page73-1_test.go
result:
goos: windows
goarch: amd64
BenchmarkNetworkRequest-8              5        2005849860 ns/op
PASS
ok      command-line-arguments  13.193s
*/
// 2005849860 ns/op
// 1305447860 ns/op
func init() {
	daemonStarted := startNetworkDaemon()
	daemonStarted.Wait()
}

func BenchmarkNetworkRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			b.Fatalf("cannot dial host: %v\n", err)
		}
		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("cannot Read: %v\n", err)
		}
		conn.Close()
	}
}

func startNetworkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("cannot listen: %v\n", err)
		}
		defer server.Close()

		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("cannot accept connection: %v", err)
				continue
			}
			connectToService()
			_, _ = fmt.Fprintln(conn, "")
			_ = conn.Close()
		}
	}()
	return &wg
}

func connectToService() interface{} {
	time.Sleep(1 * time.Second)
	return struct{}{}
}
