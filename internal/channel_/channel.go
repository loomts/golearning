package channel_

import (
	"fmt"
)

func Channel() {
	var ch chan int
	ch = make(chan int) // if not ,deadlock
	go func() {
		//ch <- 100
		close(ch) // return zero
		//ch <- 100
	}()
	fmt.Println(<-ch)

}
