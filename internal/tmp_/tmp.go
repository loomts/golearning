package tmp_

import (
	"fmt"
	"time"
)

func Tmp() {
	ch := make(chan string, 10)
	go fmt.Println(<-ch)
	go func() { fmt.Println(<-ch) }()
	ch <- "asdf"
	time.Sleep(1000000)
}
