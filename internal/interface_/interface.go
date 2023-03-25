package interface_

import "fmt"

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

func Interface_() {
	m := Mongo{} // pointer can match [pointer and value] interface, while value only match value interface.
	M := &Mongo{}

	//pointerHi(m)
	//pointerHi(M)
	valueHi(m)
	valueHi(M)

	m.pointerHello()
	M.pointerHello()
	m.valueHello()
	M.valueHello()
}
