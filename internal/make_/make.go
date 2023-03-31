package make_

import (
	"fmt"
	"golearning/internal/switch_"
)

// Make_ creates slices, maps, and channels only
func Make_() {
	a := make([]int, 10, 100)
	switch_.Switch(a)
	var mp map[int]int
	fmt.Println(mp[1])

	//mp[1] = 1 //panic
}
