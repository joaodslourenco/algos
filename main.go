package main

import "fmt"

func mergeSort(numbers []int) {
	if len(numbers) <= 1 {
		return
	}

	middleIndex := len(numbers) / 2

	leftSlice := make([]int, middleIndex)
	copy(leftSlice, numbers[:middleIndex])

	rightSlice := make([]int, len(numbers)-middleIndex)
	copy(rightSlice, numbers[middleIndex:])

	fmt.Println("left", leftSlice, "right", rightSlice)

	mergeSort(leftSlice)
	mergeSort(rightSlice)

	merge(leftSlice, rightSlice, numbers)
}

func merge(leftSlice, rightSlice, slice []int) {
	i, l, r := 0, 0, 0

	fmt.Println(slice, leftSlice, rightSlice)

	for l < len(leftSlice) && r < len(rightSlice) {
		if leftSlice[l] < rightSlice[r] {
			slice[i] = leftSlice[l]
			l++
		} else {
			slice[i] = rightSlice[r]
			r++
		}
		i++
	}

	for l < len(leftSlice) {
		slice[i] = leftSlice[l]
		l++
		i++
	}
	for r < len(rightSlice) {
		slice[i] = rightSlice[r]
		r++
		i++
	}
	fmt.Println(slice)
}

func main() {
	list := []int{6, 3, 11, 9, 12, 2, 1, 10, 8, 7, 5, 4}
	fmt.Println("Unsorted:", list)

	mergeSort(list)
	fmt.Println("Sorted:", list)
}
