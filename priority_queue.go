package prioqueue

import (
	"container/heap"
	"fmt"
	"sync"
)

// Comparator function type for custom priority comparison
type Comparator[T any] func(a, b T) int

// Item represents an item in the priority queue with its priority
type Item[T any] struct {
	Value    T
	Priority int
	Index    int // The index of the item in the heap
}

// PriorityQueue is a generic priority queue implementation
type PriorityQueue[T any] struct {
	items      []*Item[T]
	comparator Comparator[T]
	isMaxHeap  bool
	mutex      sync.RWMutex
}

// New creates a new priority queue with default integer comparison (min-heap)
func New[T any]() *PriorityQueue[T] {
	pq := &PriorityQueue[T]{
		items:     make([]*Item[T], 0),
		isMaxHeap: false,
	}
	heap.Init(pq)
	return pq
}

// NewMax creates a new max-heap priority queue
func NewMax[T any]() *PriorityQueue[T] {
	pq := &PriorityQueue[T]{
		items:     make([]*Item[T], 0),
		isMaxHeap: true,
	}
	heap.Init(pq)
	return pq
}

// NewWithComparator creates a new priority queue with custom comparator
func NewWithComparator[T any](comparator Comparator[T], isMaxHeap bool) *PriorityQueue[T] {
	pq := &PriorityQueue[T]{
		items:      make([]*Item[T], 0),
		comparator: comparator,
		isMaxHeap:  isMaxHeap,
	}
	heap.Init(pq)
	return pq
}

// Len returns the number of items in the priority queue
func (pq *PriorityQueue[T]) Len() int {
	return len(pq.items)
}

// Less compares two items based on priority
func (pq *PriorityQueue[T]) Less(i, j int) bool {
	if pq.comparator != nil {
		result := pq.comparator(pq.items[i].Value, pq.items[j].Value)
		if pq.isMaxHeap {
			return result > 0
		}
		return result < 0
	}

	// Default comparison by priority
	if pq.isMaxHeap {
		return pq.items[i].Priority > pq.items[j].Priority
	}
	return pq.items[i].Priority < pq.items[j].Priority
}

// Swap swaps two items in the priority queue
func (pq *PriorityQueue[T]) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	pq.items[i].Index = i
	pq.items[j].Index = j
}

// Push adds an item to the priority queue
func (pq *PriorityQueue[T]) Push(x interface{}) {
	n := len(pq.items)
	item := x.(*Item[T])
	item.Index = n
	pq.items = append(pq.items, item)
}

// Pop removes and returns the highest priority item
func (pq *PriorityQueue[T]) Pop() interface{} {
	old := pq.items
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	pq.items = old[0 : n-1]
	return item
}

// Enqueue adds an item with given priority
func (pq *PriorityQueue[T]) Enqueue(value T, priority int) *Item[T] {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()

	item := &Item[T]{
		Value:    value,
		Priority: priority,
	}
	heap.Push(pq, item)
	return item
}

// Dequeue removes and returns the highest priority item
func (pq *PriorityQueue[T]) Dequeue() (T, bool) {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()

	var zero T
	if len(pq.items) == 0 {
		return zero, false
	}

	item := heap.Pop(pq).(*Item[T])
	return item.Value, true
}

// Peek returns the highest priority item without removing it
func (pq *PriorityQueue[T]) Peek() (T, bool) {
	pq.mutex.RLock()
	defer pq.mutex.RUnlock()

	var zero T
	if len(pq.items) == 0 {
		return zero, false
	}

	return pq.items[0].Value, true
}

// IsEmpty returns true if the priority queue is empty
func (pq *PriorityQueue[T]) IsEmpty() bool {
	pq.mutex.RLock()
	defer pq.mutex.RUnlock()
	return len(pq.items) == 0
}

// Size returns the number of items in the priority queue
func (pq *PriorityQueue[T]) Size() int {
	return pq.Len()
}

// Clear removes all items from the priority queue
func (pq *PriorityQueue[T]) Clear() {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()
	pq.items = pq.items[:0]
	heap.Init(pq)
}

// UpdatePriority updates the priority of an existing item
func (pq *PriorityQueue[T]) UpdatePriority(item *Item[T], newPriority int) {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()

	if item.Index < 0 || item.Index >= len(pq.items) {
		return
	}

	item.Priority = newPriority
	heap.Fix(pq, item.Index)
}

// Remove removes an item from the priority queue
func (pq *PriorityQueue[T]) Remove(item *Item[T]) (T, bool) {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()

	var zero T
	if item.Index < 0 || item.Index >= len(pq.items) {
		return zero, false
	}

	removed := heap.Remove(pq, item.Index).(*Item[T])
	return removed.Value, true
}

// ToSlice returns a copy of all items as a slice (ordered by priority)
func (pq *PriorityQueue[T]) ToSlice() []Item[T] {
	pq.mutex.RLock()
	defer pq.mutex.RUnlock()

	result := make([]Item[T], len(pq.items))
	for i, item := range pq.items {
		result[i] = *item
	}
	return result
}

// String returns a string representation of the priority queue
func (pq *PriorityQueue[T]) String() string {
	pq.mutex.RLock()
	defer pq.mutex.RUnlock()

	if len(pq.items) == 0 {
		return "PriorityQueue[]"
	}

	var result string
	heapType := "min"
	if pq.isMaxHeap {
		heapType = "max"
	}

	result = fmt.Sprintf("PriorityQueue[%s-heap, size=%d]:", heapType, len(pq.items))
	for i, item := range pq.items {
		result += fmt.Sprintf("\n  [%d] Value: %v, Priority: %d", i, item.Value, item.Priority)
	}
	return result
}
