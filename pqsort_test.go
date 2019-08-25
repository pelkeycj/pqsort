package pqsort

import (
	"testing"
)

func TestPqsort(t *testing.T) {
	arr := []int{5, 2, 1, 4, 3}
	Sort(arr)
	checkSorted(arr, t)
}

func TestEmpty(t *testing.T) {
	arr := []int{}
	Sort(arr)
	checkSorted(arr, t)
}

func checkSorted(arr []int, t *testing.T) {
	if len(arr) <= 1 {
		return
	}

	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			t.Errorf("got %v > %v", arr[i-1], arr[i])
		}
	}
}
