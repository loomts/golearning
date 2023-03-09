package timer_

import (
	"fmt"
	"time"
)

func Timer_() {
	timer := time.NewTimer(3 * time.Second)
	timerr := time.NewTimer(2 * time.Second)
	//stopT := time.NewTimer(6 * time.Second)
	for {
		select {
		default:
			timer.Reset(100 * time.Second)
		case <-timer.C:
			fmt.Println("pass 3s")
			timer.Reset(3 * time.Second)
		case <-timerr.C:
			fmt.Println("pass 2s")
			timerr.Reset(2 * time.Second)
		}
	}

}
