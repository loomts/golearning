package switch_

import (
	"fmt"
	"time"
)

func Switch(t any) {
	switch {
	case false:
		fmt.Println(false)
	case true:
		fmt.Println(true)
	default:
		fmt.Println("default")
	}
	select {
	case <-time.After(time.Second):
		fmt.Println("after 1s")
	}
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t) // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", t) // t has type bool
	case int:
		fmt.Printf("integer %d\n", t) // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t has type *int
	case string:
		fmt.Printf("string %v\n", t)
	}
}
