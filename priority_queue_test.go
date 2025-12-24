package prioqueue

import (
	"fmt"
	"testing"
)

func TestNewPriorityQueue(t *testing.T) {
	pq := New[int]()
	if pq == nil {
		t.Fatal("Expected priority queue to be created")
	}
	if !pq.IsEmpty() {
		t.Error("Expected new priority queue to be empty")
	}
	if pq.Size() != 0 {
		t.Error("Expected new priority queue size to be 0")
	}
}

func TestNewMaxPriorityQueue(t *testing.T) {
	pq := NewMax[int]()
	if pq == nil {
		t.Fatal("Expected max priority queue to be created")
	}
	if !pq.isMaxHeap {
		t.Error("Expected priority queue to be max heap")
	}
}

func TestEnqueueDequeue(t *testing.T) {
	pq := New[string]()

	// Test enqueue
	pq.Enqueue("low", 3)
	pq.Enqueue("high", 1)
	pq.Enqueue("medium", 2)

	if pq.Size() != 3 {
		t.Errorf("Expected size 3, got %d", pq.Size())
	}

	// Test dequeue (min-heap, so lowest priority first)
	value, ok := pq.Dequeue()
	if !ok || value != "high" {
		t.Errorf("Expected 'high', got %v", value)
	}

	value, ok = pq.Dequeue()
	if !ok || value != "medium" {
		t.Errorf("Expected 'medium', got %v", value)
	}

	value, ok = pq.Dequeue()
	if !ok || value != "low" {
		t.Errorf("Expected 'low', got %v", value)
	}

	// Test empty dequeue
	_, ok = pq.Dequeue()
	if ok {
		t.Error("Expected dequeue from empty queue to return false")
	}
}

func TestMaxHeapBehavior(t *testing.T) {
	pq := NewMax[int]()

	pq.Enqueue(1, 1)
	pq.Enqueue(2, 2)
	pq.Enqueue(3, 3)

	// Max heap should return highest priority first
	value, ok := pq.Dequeue()
	if !ok || value != 3 {
		t.Errorf("Expected 3, got %v", value)
	}

	value, ok = pq.Dequeue()
	if !ok || value != 2 {
		t.Errorf("Expected 2, got %v", value)
	}
}

func TestPeek(t *testing.T) {
	pq := New[string]()

	// Test peek on empty queue
	_, ok := pq.Peek()
	if ok {
		t.Error("Expected peek on empty queue to return false")
	}

	pq.Enqueue("first", 1)
	pq.Enqueue("second", 2)

	// Test peek
	value, ok := pq.Peek()
	if !ok || value != "first" {
		t.Errorf("Expected 'first', got %v", value)
	}

	// Size should remain the same after peek
	if pq.Size() != 2 {
		t.Errorf("Expected size 2 after peek, got %d", pq.Size())
	}
}

func TestUpdatePriority(t *testing.T) {
	pq := New[string]()

	item1 := pq.Enqueue("item1", 3)
	item2 := pq.Enqueue("item2", 1)
	pq.Enqueue("item3", 2)

	// Update priority of item1 to make it highest priority
	pq.UpdatePriority(item1, 0)

	value, ok := pq.Dequeue()
	if !ok || value != "item1" {
		t.Errorf("Expected 'item1' after priority update, got %v", value)
	}

	// Update priority of item2 to lowest
	pq.UpdatePriority(item2, 5)

	value, ok = pq.Dequeue()
	if !ok || value != "item3" {
		t.Errorf("Expected 'item3', got %v", value)
	}
}

func TestRemove(t *testing.T) {
	pq := New[string]()

	item1 := pq.Enqueue("item1", 1)
	pq.Enqueue("item2", 2)
	pq.Enqueue("item3", 3)

	// Remove item1
	value, ok := pq.Remove(item1)
	if !ok || value != "item1" {
		t.Errorf("Expected to remove 'item1', got %v", value)
	}

	if pq.Size() != 2 {
		t.Errorf("Expected size 2 after removal, got %d", pq.Size())
	}

	// Next dequeue should be item2 (lowest remaining priority)
	value, ok = pq.Dequeue()
	if !ok || value != "item2" {
		t.Errorf("Expected 'item2' after removal, got %v", value)
	}
}

func TestClear(t *testing.T) {
	pq := New[int]()

	pq.Enqueue(1, 1)
	pq.Enqueue(2, 2)
	pq.Enqueue(3, 3)

	if pq.Size() != 3 {
		t.Errorf("Expected size 3, got %d", pq.Size())
	}

	pq.Clear()

	if !pq.IsEmpty() {
		t.Error("Expected queue to be empty after clear")
	}
	if pq.Size() != 0 {
		t.Errorf("Expected size 0 after clear, got %d", pq.Size())
	}
}

func TestCustomComparator(t *testing.T) {
	// Custom comparator for strings (by length)
	lengthComparator := func(a, b string) int {
		if len(a) < len(b) {
			return -1
		} else if len(a) > len(b) {
			return 1
		}
		return 0
	}

	pq := NewWithComparator(lengthComparator, false) // min-heap by string length

	pq.Enqueue("hello", 0)   // length 5
	pq.Enqueue("hi", 0)      // length 2
	pq.Enqueue("goodbye", 0) // length 7

	// Should dequeue shortest string first
	value, ok := pq.Dequeue()
	if !ok || value != "hi" {
		t.Errorf("Expected 'hi', got %v", value)
	}

	value, ok = pq.Dequeue()
	if !ok || value != "hello" {
		t.Errorf("Expected 'hello', got %v", value)
	}

	value, ok = pq.Dequeue()
	if !ok || value != "goodbye" {
		t.Errorf("Expected 'goodbye', got %v", value)
	}
}

func TestToSlice(t *testing.T) {
	pq := New[int]()

	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	pq.Enqueue(30, 3)

	items := pq.ToSlice()
	if len(items) != 3 {
		t.Errorf("Expected 3 items in slice, got %d", len(items))
	}

	// Verify original queue is unchanged
	if pq.Size() != 3 {
		t.Errorf("Expected original queue size 3, got %d", pq.Size())
	}
}

func TestString(t *testing.T) {
	pq := New[string]()

	// Test empty queue string representation
	str := pq.String()
	if str != "PriorityQueue[]" {
		t.Errorf("Expected 'PriorityQueue[]', got %s", str)
	}

	pq.Enqueue("test", 1)
	str = pq.String()
	if str == "PriorityQueue[]" {
		t.Error("Expected non-empty string representation")
	}
}

func TestConcurrentAccess(t *testing.T) {
	pq := New[int]()
	done := make(chan bool, 2)

	// Producer goroutine
	go func() {
		for i := 0; i < 100; i++ {
			pq.Enqueue(i, i)
		}
		done <- true
	}()

	// Consumer goroutine
	go func() {
		count := 0
		for count < 100 {
			if _, ok := pq.Dequeue(); ok {
				count++
			}
		}
		done <- true
	}()

	// Wait for both goroutines to complete
	<-done
	<-done

	if !pq.IsEmpty() {
		t.Errorf("Expected queue to be empty after concurrent operations, size: %d", pq.Size())
	}
}

func BenchmarkEnqueue(b *testing.B) {
	pq := New[int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pq.Enqueue(i, i)
	}
}

func BenchmarkDequeue(b *testing.B) {
	pq := New[int]()

	// Pre-fill the queue
	for i := 0; i < b.N; i++ {
		pq.Enqueue(i, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pq.Dequeue()
	}
}

func BenchmarkEnqueueDequeue(b *testing.B) {
	pq := New[int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pq.Enqueue(i, i)
		if i%2 == 0 {
			pq.Dequeue()
		}
	}
}

// Test with different data types
func TestGenericTypes(t *testing.T) {
	// Test with struct
	type Person struct {
		Name string
		Age  int
	}

	pq := New[Person]()
	pq.Enqueue(Person{Name: "Alice", Age: 30}, 2)
	pq.Enqueue(Person{Name: "Bob", Age: 25}, 1)

	person, ok := pq.Dequeue()
	if !ok || person.Name != "Bob" {
		t.Errorf("Expected Bob, got %v", person.Name)
	}

	// Test with interface{}
	pqAny := New[interface{}]()
	pqAny.Enqueue("string", 1)
	pqAny.Enqueue(42, 2)
	pqAny.Enqueue([]int{1, 2, 3}, 3)

	if pqAny.Size() != 3 {
		t.Errorf("Expected size 3, got %d", pqAny.Size())
	}
}

func ExamplePriorityQueue() {
	// Create a new min-heap priority queue for strings
	pq := New[string]()

	// Add items with different priorities
	pq.Enqueue("low priority task", 10)
	pq.Enqueue("high priority task", 1)
	pq.Enqueue("medium priority task", 5)

	// Process items in priority order
	for !pq.IsEmpty() {
		task, _ := pq.Dequeue()
		fmt.Println("Processing:", task)
	}
	// Output:
	// Processing: high priority task
	// Processing: medium priority task
	// Processing: low priority task
}

func ExamplePriorityQueue_maxHeap() {
	// Create a max-heap priority queue
	pq := NewMax[int]()

	pq.Enqueue(1, 1)
	pq.Enqueue(2, 2)
	pq.Enqueue(3, 3)

	for !pq.IsEmpty() {
		value, _ := pq.Dequeue()
		fmt.Println("Value:", value)
	}
	// Output:
	// Value: 3
	// Value: 2
	// Value: 1
}
