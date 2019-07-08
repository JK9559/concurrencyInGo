package main

import (
	"fmt"
	"math"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

func main() {
	producer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		for i := 5; i > 0; i-- {
			l.Lock()
			time.Sleep(1)
			l.Unlock()
		}
	}

	observer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		l.Lock()
		time.Sleep(1)
		defer l.Unlock()
	}

	test := func(count int, mutex, rwMutex sync.Locker) time.Duration {
		var wg sync.WaitGroup
		wg.Add(count + 1)
		begTestTime := time.Now()
		go producer(&wg, mutex)
		for i := count; i > 0; i-- {
			go observer(&wg, rwMutex)
		}

		wg.Wait()
		return time.Since(begTestTime)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex
	_, _ = fmt.Fprintf(tw, "Readers\tMutex\tRWMutex\n")
	for i := 0; i < 12; i++ {
		count := int(math.Pow(2, float64(i)))
		_, _ = fmt.Fprintf(tw, "%d\t%v\t%v\n", count, test(count, &m, &m), test(count, &m, m.RLocker()))
	}
}

/**
Readers  Mutex       RWMutex
1        9.019ms     8.0081ms
2        11.11ms     8.1782ms
4        11.9666ms   8.0448ms
8        17.0116ms   7.0156ms
16       24.9952ms   8.0207ms
32       54.2249ms   7.0468ms
64       94.26ms     9.073ms
128      173.1679ms  8.9677ms
256      362.7473ms  9.0071ms
512      716.3931ms  10.0738ms
1024     1.3972396s  9.1015ms
2048     2.7758371s  9.1122ms
*/
