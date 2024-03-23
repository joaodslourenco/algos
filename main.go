package main

import (
	sortAlgos "algos/merge_sort"
	"fmt"
)

func main() {
	list := []int{6, 3, 11, 9, 12, 2, 1, 10, 8, 7, 5, 4}

	fmt.Println("Unsorted:", list)
	sortAlgos.MergeSort(list)
	fmt.Println("Sorted:", list)
}
