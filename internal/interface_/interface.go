package interface_

import (
	"fmt"
	"reflect"
)

type H interface {
	//pointerHi()
	valueHi()
}

//func pointerHi(h H) {
//	h.pointerHi()
//}

func valueHi(h H) {
	h.valueHi()
}

type Mongo struct{}

//func (mongo *Mongo) pointerHi() {
//	fmt.Println("pointerHi mongo")
//}

func (mongo Mongo) valueHi() {
	fmt.Println("valueHi mongo")
}

func (mongo *Mongo) pointerHello() {
	fmt.Println("pointerHello mongo")
}

func (mongo Mongo) valueHello() {
	fmt.Println("valueHello mongo")
}

type Pet interface {
	Name() string
}

type Duck struct {
	name string
}

func (d Duck) Name() string {
	return d.name
}

func Interface_() {

	m := Mongo{} // pointer can match [pointer and value] interface, while value only match value interface.
	M := &Mongo{}
	a := 0.0
	const zero = 0.0
	//haha:
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(zero))
	fmt.Println(reflect.TypeOf(a).Kind())
	fmt.Println(reflect.TypeOf(zero).Kind())
	//pointerHi(m)
	//pointerHi(M)
	valueHi(m)
	valueHi(M)
	//goto haha
	m.pointerHello()
	M.pointerHello()
	m.valueHello()
	M.valueHello()
}
