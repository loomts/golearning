package make_

import (
	"golearning/internal/switch_"
)

// Make_ creates slices, maps, and channels only
func Make_() {
	a := make([]int, 10, 100)
	switch_.GetType(a)

}
