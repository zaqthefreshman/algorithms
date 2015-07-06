package mergesort

func MergeSort(unsorted []int) []int {
	if length := len(unsorted); length >= 2 {
		mid := length / 2
		//fmt.Println(unsorted, "mid:", mid)
		//fmt.Println("left", unsorted[:mid])
		//fmt.Println("right", unsorted[mid:])
		left := MergeSort(unsorted[:mid])
		right := MergeSort(unsorted[mid:])
		return merge(left, right)
	}

	return unsorted
}

func merge(left, right []int) (output []int) {
	length := len(left) + len(right)
	output = make([]int, length)

	//fmt.Println("merge length", length)

	i, j := 0, 0
	for k := 0; k < length; k++ {
		//fmt.Println("k:", k, "i:", i, "j:", j)
		if len(left) <= i {
			if len(right) > j {
				output[k] = right[j]
				j++
				continue
			}
		}
		if len(right) <= j {
			if len(left) > i {
				output[k] = left[i]
				i++
				continue
			}
		}
		if left[i] < right[j] {
			output[k] = left[i]
			i++
		} else {
			output[k] = right[j]
			j++
		}
	}

	return
}
