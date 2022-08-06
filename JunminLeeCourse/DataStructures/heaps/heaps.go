package heaps

import "fmt"

// MaxHeap class
type MaxHeap struct {
	array []int
}

// Insert method
func (h *MaxHeap) Insert(i int) {
	h.array = append(h.array, i)
	h.MaxHeapifyUp(len(h.array) - 1)
}

// Extract method
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]

	l := len(h.array) - 1

	if len(h.array) == 0 {
		fmt.Println("cannot extract from empty array")
		return -1
	}

	h.array[0] = h.array[l]
	h.array = h.array[:l]
	h.MaxHeapifyDown(0)

	return extracted
}

// MaxHeapifyUp method
func (h *MaxHeap) MaxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// MaxHeapifyDown method
func (h *MaxHeap) MaxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex || h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

// get the parent index
func parent(i int) int {
	return (i - 1) / 2
}

// get the left child index
func left(i int) int {
	return i*2 + 1
}

// get the right child index
func right(i int) int {
	return i*2 + 2
}

// swap keys
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}
