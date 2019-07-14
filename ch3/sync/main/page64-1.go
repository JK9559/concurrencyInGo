package main

import (
	"fmt"
	"sync"
	"time"
)

/**
总结1 使用Cond的时候 加锁时需要注意：首先明确判断条件是什么，然后在使用条件的临界区加锁和解锁
，也就是说针对条件的操作必须在临界区(上锁的范围之内)
*/
func main() {
	/**
	  代码功能：不断向队列内添加元素并异步(通过sleep实现)出队，保持队列长度为2，如果长度大于等于2
	  ，则等待出队完成再继续添加元素
	*/
	c := sync.NewCond(&sync.Mutex{})
	// 创建长度为0的slice 容量为10
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		// 进入临界区 以便我们可以修改条件的临界区
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Removed from queue")
		// 退出条件临界区 因为我们已经对条件操作结束了
		c.L.Unlock()
		// 向wait发出信号 表示条件已经被触发
		c.Signal()
	}

	for i := 0; i < 100; i++ {
		// 锁定了条件(也就是说在条件所在的那层进行锁定) 因为在进入Locker的时候，执行wait会自动执行unlock
		// 进入临界区
		c.L.Lock()
		// =========>注意这里使用了循环<=========
		// 主要是因为接受到的信号并不一定是你所期待的信号
		for len(queue) >= 2 {
			fmt.Println("Yes,Length now is: ", len(queue))
			// wait暂停了main goroutine直到一个信号条件到达
			c.Wait()
		}
		fmt.Println("Adding to Queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second)
		// 为条件Locker进行解锁操作，当wait退出的时候 会自动执行lock
		// 退出条件临界区
		c.L.Unlock()
	}
}
