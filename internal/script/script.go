package script

import (
	"fmt"
	"strings"
)

func Script() {
	s0 := "123 345"
	s := strings.Split(s0, " ")
	fmt.Printf("%v %v", s[0], s[1])
}
