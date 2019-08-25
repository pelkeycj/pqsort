package pqsort

import (
	"sync"
)

const (
	threshold = 10000
)

func Sort(arr []int) {
	var wg sync.WaitGroup
	sort(arr, &wg, 0, len(arr)-1)
	wg.Wait()
}

func SortSeq(arr []int) {
	sortSeq(arr, 0, len(arr)-1)
}


func sortSeq(arr []int, left, right int) {
	if left >= right {
		return
	}

	pIdx := partition(arr, left, right)
	sortSeq(arr, left, pIdx-1)
	sortSeq(arr, pIdx+1, right)
	return
}

func sort(arr []int, wg *sync.WaitGroup, left, right int) {
	if left >= right {
		return
	}

	pIdx := partition(arr, left, right)

	// sort sequentially if too few items
	if right - left < threshold {
		sortSeq(arr, left, pIdx-1)
		sortSeq(arr, pIdx+1, right)
	} else {
		wg.Add(2)

		go func() {
			defer wg.Done()
			sort(arr, wg, left, pIdx-1)
		}()

		go func() {
			defer wg.Done()
			sort(arr, wg, pIdx+1, right)
		}()
	}

	return
}

func partition(arr []int, left, right int) int {
	pVal := arr[right]
	i := left

	for j := i; j < right; j++ {
		if arr[j] < pVal {
			arr[i], arr[j] = arr[j], arr[i]
			i += 1
		}
	}
	arr[i], arr[right] = arr[right], arr[i]
	return i
}
