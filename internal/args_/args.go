package args_

import (
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "the greeting object.")
}
func Args() {
	flag.Parse()
	fmt.Println("pointerHello, " + name + "!")
}
