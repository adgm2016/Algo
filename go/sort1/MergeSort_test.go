package sort1

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{5, 4}
	MergeSort(arr)
	fmt.Println(arr)

	arr = []int{5, 4, 3, 2, 1}
	MergeSort(arr)
	fmt.Println(arr)
}
