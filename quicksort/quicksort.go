package quicksort

import "math/rand"

func QuickSort(unsorted []int) []int {
	length := len(unsorted)
	if length <= 1 {
		return unsorted
	}

	pivot := rand.Int() % length
	i := partition(unsorted, pivot)
	QuickSort(unsorted[:i])
	QuickSort(unsorted[i+1:])

	return unsorted
}

func partition(a []int, l int) int {
	p := a[l]
	i := 1
	a[l], a[0] = a[0], a[l]
	for j := 1; j < len(a); j++ {
		if a[j] < p {
			a[j], a[i] = a[i], a[j]
			i++
		}
	}
	a[0], a[i-1] = a[i-1], a[0]
	return i - 1
}
