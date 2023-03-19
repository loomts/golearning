package slide_

import "fmt"

func Slide_() {
	// v := make([]int, 3, 10)
	// for i, _ := range v {
	// 	v[i] = i
	// }
	// fmt.Printf("%v\n", v)
	// v = append(v[:5], v[2:]...)
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
	mp := make(map[int][]int)
	mp[2] = []int{}
	for k, v := range mp {
		fmt.Print("------", k, v)
	}
}
