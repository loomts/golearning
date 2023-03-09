package redeclear_

import (
	"fmt"
	"log"
	"os"
)

func Redeclear_() {
	name := "file"
	f, err := os.Open(name)
	if err != nil {
		log.Fatalf("%v", err)
	}
	d, err := f.Stat()
	if err != nil {
		f.Close()
		log.Fatalf("%v", err)
	}
	fmt.Println(d)
}
