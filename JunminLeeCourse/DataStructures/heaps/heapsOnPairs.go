package heaps

import "fmt"

// MinPairHeap class
type MinPairHeap struct {
	Array [][]int
}

// Insert method
func (h *MinPairHeap) Insert(i []int) {
	h.Array = append(h.Array, i)
	h.MinPairHeapifyUp(len(h.Array) - 1)
}

// Extract method
func (h *MinPairHeap) Extract() []int {
	extracted := h.Array[0]

	l := len(h.Array) - 1

	if len(h.Array) == 0 {
		fmt.Println("cannot extract from empty array")
		return []int{}
	}

	h.Array[l] = h.Array[0]
	h.Array = h.Array[:l]
	h.MinPairHeapifyDown(0)

	return extracted
}

// MinPairHeapifyUp method
func (h *MinPairHeap) MinPairHeapifyUp(index int) {
	for h.Array[parent(index)][0]+h.Array[parent(index)][1] > h.Array[index][0]+h.Array[index][1] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// MinPairHeapifyDown method
func (h *MinPairHeap) MinPairHeapifyDown(index int) {
	lastIndex := len(h.Array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex || h.Array[l][0]+h.Array[l][1] < h.Array[r][0]+h.Array[r][1] {
			childToCompare = l
		} else {
			childToCompare = r
		}
		if h.Array[index][0]+h.Array[index][1] > h.Array[childToCompare][0]+h.Array[childToCompare][1] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

//// get the parent index
//func parent(i int) int {
//	return (i - 1) / 2
//}
//
//// get the left child index
//func left(i int) int {
//	return i*2 + 1
//}
//
//// get the right child index
//func right(i int) int {
//	return i*2 + 2
//}

// swap keys
func (h *MinPairHeap) swap(i1, i2 int) {
	h.Array[i1], h.Array[i2] = h.Array[i2], h.Array[i1]
}
