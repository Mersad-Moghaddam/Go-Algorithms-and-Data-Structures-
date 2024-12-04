package heap

import (
	"errors"

	"github.com/TheAlgorithms/Go/constraints"
)

// Heap represents a generic binary heap implementation.
// The heap maintains elements in a specific order determined by a comparator function.
type Heap[T any] struct {
	heaps    []T               // Slice to store heap elements.
	lessFunc func(a, b T) bool // Comparator function to define heap ordering.
}

// New creates a new Heap instance for ordered types.
// It uses the default comparator (a < b) for the heap's ordering.
func New[T constraints.Ordered]() *Heap[T] {
	defaultLess := func(a, b T) bool {
		return a < b
	}
	h, _ := NewAny[T](defaultLess) // Error is ignored as defaultLess is valid.
	return h
}

// NewAny creates a new Heap instance for any type T.
// The caller must provide a valid comparator function (less).
func NewAny[T any](less func(a, b T) bool) (*Heap[T], error) {
	if less == nil {
		return nil, errors.New("less function is required to define heap ordering")
	}
	return &Heap[T]{
		lessFunc: less,
	}, nil
}

// Push adds a new element to the heap.
// Complexity: O(log n), where n is the number of elements in the heap.
func (h *Heap[T]) Push(element T) {
	h.heaps = append(h.heaps, element) // Add the element at the end.
	h.up(len(h.heaps) - 1)             // Restore the heap property.
}

// Top returns the smallest element (based on lessFunc) from the heap.
// Panics if the heap is empty.
func (h *Heap[T]) Top() T {
	if h.Empty() {
		panic("cannot retrieve top element from an empty heap")
	}
	return h.heaps[0]
}

// Pop removes the smallest element (based on lessFunc) from the heap.
// Complexity: O(log n), where n is the number of elements in the heap.
func (h *Heap[T]) Pop() {
	if h.Empty() {
		return
	}

	// Replace the root with the last element and shrink the slice.
	h.swap(0, len(h.heaps)-1)
	h.heaps = h.heaps[:len(h.heaps)-1]

	// Restore the heap property by "sinking down" the root.
	if len(h.heaps) > 0 {
		h.down(0)
	}
}

// Empty checks whether the heap is empty.
func (h *Heap[T]) Empty() bool {
	return len(h.heaps) == 0
}

// Size returns the number of elements currently in the heap.
func (h *Heap[T]) Size() int {
	return len(h.heaps)
}

// swap exchanges elements at indices i and j in the heap.
func (h *Heap[T]) swap(i, j int) {
	h.heaps[i], h.heaps[j] = h.heaps[j], h.heaps[i]
}

// up restores the heap property by "bubbling up" the element at the given index.
func (h *Heap[T]) up(index int) {
	if index <= 0 {
		return
	}

	// Find the parent index.
	parent := (index - 1) / 2

	// Compare the element with its parent; if it's in the wrong order, swap them.
	if h.lessFunc(h.heaps[index], h.heaps[parent]) {
		h.swap(index, parent)
		h.up(parent) // Recursively adjust the parent.
	}
}

// down restores the heap property by "sinking down" the element at the given index.
func (h *Heap[T]) down(index int) {
	smallest := index
	left, right := 2*index+1, 2*index+2

	// Check if the left child exists and is smaller than the current element.
	if left < len(h.heaps) && h.lessFunc(h.heaps[left], h.heaps[smallest]) {
		smallest = left
	}

	// Check if the right child exists and is smaller than the current element.
	if right < len(h.heaps) && h.lessFunc(h.heaps[right], h.heaps[smallest]) {
		smallest = right
	}

	// If the smallest element is not the parent, swap and continue sinking.
	if smallest != index {
		h.swap(index, smallest)
		h.down(smallest) // Recursively adjust the subtree.
	}
}
