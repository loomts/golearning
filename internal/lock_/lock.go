package lock_

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	mu sync.Mutex
	ch chan int
}

func (t1 Task) inL(seq int) {
	t1.mu.Lock()
	t1.ch <- 1
	t1.ch <- 1
	t1.ch <- 1
	t1.ch <- 1
	t1.ch <- 1
	t1.ch <- 1
	t1.ch <- 1
	t1.ch <- 1
	t1.ch <- 1
	t1.ch <- 1
	t1.ch <- 1
	t1.ch <- 1
	t1.ch <- 1
	t1.ch <- 1
	t1.ch <- 1
	fmt.Printf("%vst ch<-1\n", seq)
	t1.mu.Unlock()
}

func (t1 Task) out() {
	for {
		select {
		case <-t1.ch:
			fmt.Println("ch->1")
			t1.mu.Lock()
			time.Sleep(time.Second)
			t1.mu.Unlock()
			fmt.Println("out")
		}
	}
}

func Lock_() {
	// task := Task{
	// 	mu:    sync.Mutex{},
	// 	peers: nil,
	// }
	// for i := 0; i < 10000; i++ {
	// 	task.peers = append(task.peers, i+i)
	// }
	// var wg sync.WaitGroup
	// wg.Add(len(task.peers))
	// task.mu.Lock()
	// for i, i2 := range task.peers {
	// 	go func(i int, i2 int) {
	// 		defer wg.Done()
	// 		task.mu.Lock()
	// 		fmt.Printf("%v %v\n", i, i2)
	// 		task.mu.Unlock()
	// 	}(i, i2)
	// }
	// task.mu.Unlock()
	// wg.Wait()
	t1 := Task{mu: sync.Mutex{}, ch: make(chan int, 1)}
	go t1.inL(1)
	// go t1.inL(2)
	// go t1.inL(3)
	// go t1.inL(4)
	// go t1.inL(5)
	// go t1.inL(6)
	go t1.out()
	time.Sleep(time.Hour)
}
