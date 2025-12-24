# PrioQueue
<br>
<div align="center">
  <img src="logo.png" alt="PrioQueue Logo" width="600"/>
</div>

A high-performance, thread-safe, generic priority queue implementation for Go.

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.18-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/l00pss/prioqueue)](https://goreportcard.com/report/github.com/l00pss/prioqueue)

## Features

 **High Performance**: Built on Go's container/heap for optimal performance  
 **Thread-Safe**: Concurrent access support with mutex protection  
 **Generic Support**: Type-safe implementation using Go generics  
 **Advanced Operations**: Priority updates, item removal, custom comparators  
 **Flexible Ordering**: Support for both min-heap and max-heap configurations  
 **Well Tested**: Comprehensive test suite with benchmarks  

## Installation

```bash
go get github.com/l00pss/prioqueue
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/l00pss/prioqueue"
)

func main() {
    // Create a new priority queue for strings
    pq := prioqueue.New[string]()
    
    // Add items with priorities (lower number = higher priority)
    pq.Enqueue("Low priority task", 10)
    pq.Enqueue("High priority task", 1) 
    pq.Enqueue("Medium priority task", 5)
    
    // Process items in priority order
    for !pq.IsEmpty() {
        task, ok := pq.Dequeue()
        if ok {
            fmt.Println("Processing:", task)
        }
    }
}
```

## API Reference

### Creating Priority Queues

```go
// Create a min-heap priority queue (lower priority number = higher priority)
pq := prioqueue.New[string]()

// Create a max-heap priority queue (higher priority number = higher priority) 
pq := prioqueue.NewMax[int]()

// Create with custom comparator
comparator := func(a, b string) int {
    return strings.Compare(a, b)
}
pq := prioqueue.NewWithComparator(comparator, false)
```

### Basic Operations

```go
// Add items
item := pq.Enqueue("value", 5)

// Remove highest priority item
value, ok := pq.Dequeue()

// Peek at highest priority item without removing
value, ok := pq.Peek()

// Check if empty
isEmpty := pq.IsEmpty()

// Get size
size := pq.Size()

// Clear all items
pq.Clear()
```

### Advanced Operations

```go
// Update priority of existing item
pq.UpdatePriority(item, newPriority)

// Remove specific item
value, ok := pq.Remove(item)

// Get all items as slice
items := pq.ToSlice()

// String representation
fmt.Println(pq.String())
```

## Examples

### Task Scheduling

```go
type Task struct {
    ID          string
    Description string
    Deadline    time.Time
}

// Create priority queue for tasks
pq := prioqueue.New[Task]()

// Add tasks with priority based on urgency
pq.Enqueue(Task{ID: "1", Description: "Critical bug fix", Deadline: time.Now().Add(1 * time.Hour)}, 1)
pq.Enqueue(Task{ID: "2", Description: "Feature request", Deadline: time.Now().Add(24 * time.Hour)}, 5)
pq.Enqueue(Task{ID: "3", Description: "Code review", Deadline: time.Now().Add(4 * time.Hour)}, 3)

// Process tasks by priority
for !pq.IsEmpty() {
    task, _ := pq.Dequeue()
    fmt.Printf("Processing task: %s\n", task.Description)
}
```

### Custom Priority Logic

```go
type Customer struct {
    Name string
    VIPLevel int
}

// Custom comparator: VIP customers first
vipComparator := func(a, b Customer) int {
    if a.VIPLevel != b.VIPLevel {
        return b.VIPLevel - a.VIPLevel // Higher VIP level first
    }
    return strings.Compare(a.Name, b.Name) // Then alphabetical
}

pq := prioqueue.NewWithComparator(vipComparator, false)

pq.Enqueue(Customer{Name: "Alice", VIPLevel: 1}, 0)
pq.Enqueue(Customer{Name: "Bob", VIPLevel: 3}, 0)
pq.Enqueue(Customer{Name: "Charlie", VIPLevel: 2}, 0)

// Bob (VIP 3) will be processed first
```

### Gaming Leaderboard

```go
type Player struct {
    Name  string
    Score int
}

// Max heap for highest scores first
leaderboard := prioqueue.NewMax[Player]()

leaderboard.Enqueue(Player{Name: "Alice", Score: 1500}, 1500)
leaderboard.Enqueue(Player{Name: "Bob", Score: 2000}, 2000)
leaderboard.Enqueue(Player{Name: "Charlie", Score: 1800}, 1800)

// Print top players
fmt.Println("🏆 Leaderboard:")
for i := 0; i < 3 && !leaderboard.IsEmpty(); i++ {
    player, _ := leaderboard.Dequeue()
    fmt.Printf("%d. %s - %d points\n", i+1, player.Name, player.Score)
}
```

## Performance

The priority queue is built on Go's optimized `container/heap` package:

- **Enqueue**: O(log n)
- **Dequeue**: O(log n) 
- **Peek**: O(1)
- **Update Priority**: O(log n)
- **Remove**: O(log n)

### Benchmarks

```
BenchmarkEnqueue-8           1000000      1043 ns/op
BenchmarkDequeue-8           2000000       854 ns/op  
BenchmarkEnqueueDequeue-8    1000000      1205 ns/op
```

## Thread Safety

All operations are thread-safe and can be used concurrently:

```go
pq := prioqueue.New[int]()

// Safe to use from multiple goroutines
go func() {
    for i := 0; i < 100; i++ {
        pq.Enqueue(i, i)
    }
}()

go func() {
    for i := 0; i < 100; i++ {
        pq.Dequeue()
    }
}()
```

## Testing

Run the comprehensive test suite:

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run benchmarks
go test -bench=. ./...

# Run tests with race detection
go test -race ./...
```


## Requirements

- Go 1.25 or higher (for generics support)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<div align="center">
  Made with ❤️ for the Go
</div>
<div align="center">
  <a href="https://github.com/l00pss/prioqueue">⭐ Star on GitHub ⭐</a>
</div>
