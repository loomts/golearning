package fetchURL_

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

func FetchURL() {
	wg := sync.WaitGroup{}
	for _, url := range os.Args[1:] {
		wg.Add(1)
		url := url
		go func() {
			defer wg.Done()
			fetch(url)
		}()
	}
	wg.Wait()
	fmt.Println("okay!")
}

func fetch(url string) {
	start := time.Now()
	defer fmt.Printf("fetch %s used time: %.2f\n", url, time.Since(start).Seconds())
	if resp, err := http.Get(url); err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	} else {
		b, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
