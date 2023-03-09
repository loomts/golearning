package interface_

import (
	"bufio"
	"fmt"
	"net/http"
)

type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type Counter int
type Chan chan *http.Request

func (ch Chan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//ch <- req
	fmt.Fprintf(w, "url sent")
	fmt.Fprintf(w, "%v", req.URL)
	//fmt.Fprintf(w, "%v", <-ch)
}

func Interface_() {
	// sometime, implements an interface needs runtime check
	ctr := new(Chan)

	http.Handle("/handler", ctr) // register handler
	(&http.Server{Addr: "localhost:8000"}).ListenAndServe()
}

// ReadWriter is the struct that implement interface that combines the Reader and Writer interfaces.
type ReadWriter struct {
	Reader *bufio.Reader
	Writer *bufio.Writer
}

func (rw *ReadWriter) Read(p []byte) (n int, err error) {
	return rw.Reader.Read(p)
}
