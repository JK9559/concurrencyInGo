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
go test -benchtime=10s -bench=. D:\goWorkSpace\Go\src\concurrencyInGo\ch3\pool\main\page75-1_test.go
result:
goos: windows
goarch: amd64
BenchmarkNetworkRequestByCache-8              10        1406239980 ns/op
PASS
ok      command-line-arguments  24.975s
结论：比每次自己申请快一点点 应该是win的影响？
*/

func init() {
	daemonStarted := startNetworkDaemonByCache()
	daemonStarted.Wait()
}

func BenchmarkNetworkRequestByCache(b *testing.B) {
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

func warmServiceConnCache() *sync.Pool {
	p := &sync.Pool{
		New: connectToServiceByCache,
	}
	for i := 0; i < 10; i++ {
		p.Put(p.New)
	}
	return p
}

func startNetworkDaemonByCache() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		connPool := warmServiceConnCache()
		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("cannot listen: %v\n", err)
		}
		defer server.Close()

		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				log.Fatalf("cannot accept connection: %v\n", err)
			}
			svrConn := connPool.Get()
			_, _ = fmt.Fprintln(conn, "")
			connPool.Put(svrConn)
			conn.Close()
		}
	}()
	return &wg
}

func connectToServiceByCache() interface{} {
	time.Sleep(1 * time.Second)
	return struct{}{}
}
