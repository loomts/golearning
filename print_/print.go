package print_

import (
	"fmt"
	"os"
)

type T struct {
	a int
	b float64
	c string
}
type MyString string

func (m MyString) String() string {
	return fmt.Sprintf("MyString=%s", string(m)) // OK: note conversion.
}

// to string
func (t *T) String() string {
	return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c)
}
func Print_() {
	fmt.Printf("Hello %d\n", 23)
	fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
	fmt.Println("Hello", 23)
	fmt.Println(fmt.Sprint("Hello ", 23)) // make it as a string

	t := &T{7, -2.35, "abc\tdef"}
	
	fmt.Printf("%v\n", t)
	fmt.Printf("%v\n", t)
	fmt.Printf("%+v\n", t)
	fmt.Printf("%#v\n", t)
	var str MyString
	str = "myString"
	fmt.Printf("%v\n", str)
}
