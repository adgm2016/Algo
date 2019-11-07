package recursion

import "testing"

func TestRangeType_RangeAll(t *testing.T) {
	slice1 := NewRangeArray(4)
	for i := 0; i < 4; i++ {
		slice1.val[i] = i + 1
	}
	slice1.RangeAll(0)

	slice2 := NewRangeArray(3)
	slice2.val[0] = "a"
	slice2.val[1] = "b"
	slice2.val[2] = "c"
	slice2.RangeAll(0)
}
