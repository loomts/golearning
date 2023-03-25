package slice_

import (
	"fmt"
)

var container = []string{"zero", "one"}

func Slide_() {

	//mp := make(map[int]struct{})
	//a := make([]int, 13)
	//for cap(a) < 5000 {
	//	a = append(a, 1)
	//	mp[cap(a)] = struct{}{}
	//}
	//var b []float64
	//for k, _ := range mp {
	//	b = append(b, float64(k))
	//}
	//sort.Float64s(b)
	//for i := 0; i < len(b)-1; i++ {
	//	fmt.Println(b[i+1]/b[i], b[i])
	//}
	mp := map[int]int{}
	mp[1] = 1123
	fmt.Println(mp[1])
	a := make([]int, 1)
	a[0] = 1
	fmt.Printf("%p\n", &a)
	fmt.Printf("%v\n", a)
	func(b []int) {
		fmt.Printf("%p\n", &b)
		b[0] = 2
	}(a)
	fmt.Printf("%v\n", a)

	// fmt.Printf("%v\n", v)
	//var a []int
	//b := v // shallow copy
	//a = make([]int, 3)
	//copy(a, v) // deep copy
	//fmt.Printf("sourcePointer = %p, pointer = %p, v = %v, len = %v, cap = %v\n", v, &v, v, len(v), cap(v))
	//fmt.Printf("sourcePointer = %p, pointer = %p, b = %v, len = %v, cap = %v\n", b, &b, b, len(b), cap(b))
	//fmt.Printf("sourcePointer = %p, pointer = %p, a = %v, len = %v, cap = %v\n", a, &a, a, len(a), cap(a))
	//type LinesOfText [][]byte
	//text := LinesOfText{
	//	[]byte("Now is the time"),
	//	[]byte("for all good gophers"),
	//	[]byte("to bring some fun to the party."),
	//}
	//for _, bytes := range text {
	//	print(string(bytes))
	//}
	//var mp map[string]bool
	//fmt.Printf("\nval = %v, %p\n", mp, mp)
	//mp = make(map[string]bool)
	//fmt.Printf("\nval = %v, %p\n", mp, mp)
	//mp["1"] = false
	//mp["2"] = true
	//second, exist := mp["1"]
	//fmt.Println(second, exist)
	//second, exist = mp["3"]
	//fmt.Println(second, exist)
	//mp := make(map[int][]int)
	//mp[2] = []int{}
	//for k, v := range mp {
	//	fmt.Print("------", k, v)
	//}
	//container := map[int]string{0: "zero", 1: "one"}
	//fmt.Printf("The first element:%q.\n", container[0])
}
