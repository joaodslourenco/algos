package sortAlgos

func MergeSort(numbers []int) {
	if len(numbers) <= 1 {
		return
	}

	middleIndex := len(numbers) / 2

	leftSlice := make([]int, middleIndex)
	copy(leftSlice, numbers[:middleIndex])

	rightSlice := make([]int, len(numbers)-middleIndex)
	copy(rightSlice, numbers[middleIndex:])

	MergeSort(leftSlice)
	MergeSort(rightSlice)

	merge(leftSlice, rightSlice, numbers)
}

func merge(leftSlice, rightSlice, slice []int) {
	i, l, r := 0, 0, 0

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

}
