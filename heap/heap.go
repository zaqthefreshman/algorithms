package heap

type MaxHeap []int

func Heap(a []int) MaxHeap {
	h := MaxHeap(a)
	for i := len(h) / 2; i >= 0; i-- {
		h.maxHeapify(i)
	}
	return h
}

func (h MaxHeap) maxHeapify(i int) {
	l, r := i*2+1, i*2+2
	large := i
	if l < len(h) && h[l] >= h[large] {
		large = l
	}
	if r < len(h) && h[r] >= h[large] {
		large = r
	}

	if large != i {
		h[large], h[i] = h[i], h[large]
		h.maxHeapify(large)
	}
}
