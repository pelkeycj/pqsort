package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
	"github.com/pelkeycj/pqsort"
)


func main() {
	var count int
	var parallel bool
	flag.IntVar(&count, "count", 100, "number of ints to sort")
	flag.BoolVar(&parallel, "parallel", true, "whether or not to sort in parallel")

	flag.Parse()
	fmt.Printf("Creating slice of %d ints...\n", count)

	arr := buildArr(count)

	fmt.Println("Sorting slice of %d ints with parallel=%v...\n", count, parallel)
	
	start := time.Now()
	if parallel {
		pqsort.Sort(arr)
	} else {
		pqsort.SortSeq(arr)
	}
	end := time.Now()

	elapsed := end.Sub(start)
	fmt.Printf("Time elapsed: %v\n", elapsed)

}

func buildArr(count int) []int {
	arr := make([]int, count)
	for i, _:= range arr {
		arr[i] = i
	}

	for i, _ := range arr {
		j := rand.Intn(count)
		arr[j], arr[i] = arr[i], arr[j]
	}

	return arr
}