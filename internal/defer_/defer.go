package defer_

import (
	"fmt"
	"io"
	"os"
)

// Contents returns the file's contents as a string.
func Contents(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close() // f.Close will run when we're finished.

	var result []byte
	buf := make([]byte, 100)
	for {
		n, err := f.Read(buf[0:])
		result = append(result, buf[0:n]...) // append is discussed later.
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err // f will be closed if we return here.
		}
	}
	return string(result), nil // f will be closed if we return here.
}

func LIFO() {
	for i := 0; i < 5; i++ {
		func() {
			defer fmt.Printf("%d ", i)
		}()
		// deferred functions are executed in LIFO order, so this code will cause 4 3 2 1 0 to be printed when the function returns
		defer fmt.Printf("%d ", i)
	}
	fmt.Println()
}

// Trace deferred
func Trace() {
	b()
}

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func data() {
	a := make([]int, 10, 100)
	fmt.Println(a)
}
