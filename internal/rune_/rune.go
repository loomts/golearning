package rune_

import "fmt"

func Rune() {
	s := "我是dakta"
	r := []rune(s)
	fmt.Println(r)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%v", string(r[i]))
	}
}
