package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/l00pss/prioqueue"
)

// Task represents a task with deadline
type Task struct {
	ID          string
	Description string
	Deadline    time.Time
}

// Customer represents a customer with VIP level
type Customer struct {
	Name     string
	VIPLevel int
}

// Player represents a game player with score
type Player struct {
	Name  string
	Score int
}

func main() {
	fmt.Println("* PrioQueue Examples\n")

	// Example 1: Basic usage with strings
	basicExample()

	// Example 2: Task scheduling
	taskSchedulingExample()

	// Example 3: Custom comparator for VIP customers
	vipCustomerExample()

	// Example 4: Gaming leaderboard with max heap
	gamingLeaderboardExample()

	// Example 5: Advanced operations
	advancedOperationsExample()
}

func basicExample() {
	fmt.Println("📋 Example 1: Basic Priority Queue Usage")
	fmt.Println("======================================")

	// Create a new priority queue for strings
	pq := prioqueue.New[string]()

	// Add items with different priorities (lower number = higher priority)
	pq.Enqueue("Low priority task", 10)
	pq.Enqueue("High priority task", 1)
	pq.Enqueue("Medium priority task", 5)
	pq.Enqueue("Very high priority task", 0)

	fmt.Printf("Queue size: %d\n", pq.Size())

	// Process items in priority order
	fmt.Println("Processing tasks in priority order:")
	for !pq.IsEmpty() {
		task, ok := pq.Dequeue()
		if ok {
			fmt.Printf("  * Processing: %s\n", task)
		}
	}
	fmt.Println()
}

func taskSchedulingExample() {
	fmt.Println("-- Example 2: Task Scheduling System")
	fmt.Println("===================================")

	// Create priority queue for tasks
	pq := prioqueue.New[Task]()

	// Add tasks with priority based on urgency
	now := time.Now()
	tasks := []Task{
		{ID: "1", Description: "Critical security patch", Deadline: now.Add(1 * time.Hour)},
		{ID: "2", Description: "Feature development", Deadline: now.Add(24 * time.Hour)},
		{ID: "3", Description: "Code review", Deadline: now.Add(4 * time.Hour)},
		{ID: "4", Description: "Emergency bug fix", Deadline: now.Add(30 * time.Minute)},
		{ID: "5", Description: "Documentation update", Deadline: now.Add(48 * time.Hour)},
	}

	// Enqueue tasks with priority based on deadline urgency
	for _, task := range tasks {
		priority := int(task.Deadline.Sub(now).Minutes()) // Sooner deadline = lower number = higher priority
		pq.Enqueue(task, priority)
		fmt.Printf("  📝 Added: %s (deadline in %v)\n", task.Description, task.Deadline.Sub(now).Round(time.Minute))
	}

	fmt.Println("\nProcessing tasks by deadline urgency:")
	for !pq.IsEmpty() {
		task, _ := pq.Dequeue()
		remaining := task.Deadline.Sub(now).Round(time.Minute)
		fmt.Printf("  🔧 Processing: %s (deadline in %v)\n", task.Description, remaining)
	}
	fmt.Println()
}

func vipCustomerExample() {
	fmt.Println("-- Example 3: VIP Customer Service Queue")
	fmt.Println("=======================================")

	// Custom comparator: VIP customers first, then alphabetical
	vipComparator := func(a, b Customer) int {
		if a.VIPLevel != b.VIPLevel {
			return b.VIPLevel - a.VIPLevel // Higher VIP level first
		}
		return strings.Compare(a.Name, b.Name) // Then alphabetical
	}

	pq := prioqueue.NewWithComparator(vipComparator, false)

	customers := []Customer{
		{Name: "Alice Johnson", VIPLevel: 1},
		{Name: "Bob Wilson", VIPLevel: 3},
		{Name: "Charlie Brown", VIPLevel: 2},
		{Name: "Diana Prince", VIPLevel: 3},
		{Name: "Eve Adams", VIPLevel: 1},
	}

	fmt.Println("Adding customers to queue:")
	for _, customer := range customers {
		pq.Enqueue(customer, 0) // Priority is determined by comparator
		vipStatus := "Regular"
		if customer.VIPLevel == 2 {
			vipStatus = "Silver"
		} else if customer.VIPLevel == 3 {
			vipStatus = "Gold"
		}
		fmt.Printf("  👤 %s (%s VIP)\n", customer.Name, vipStatus)
	}

	fmt.Println("\nServing customers by VIP priority:")
	for !pq.IsEmpty() {
		customer, _ := pq.Dequeue()
		vipStatus := "Regular"
		if customer.VIPLevel == 2 {
			vipStatus = "Silver"
		} else if customer.VIPLevel == 3 {
			vipStatus = "Gold"
		}
		fmt.Printf("  -> Now serving: %s (%s VIP)\n", customer.Name, vipStatus)
	}
	fmt.Println()
}

func gamingLeaderboardExample() {
	fmt.Println("-- Example 4: Gaming Leaderboard (Max Heap)")
	fmt.Println("===========================================")

	// Max heap for highest scores first
	leaderboard := prioqueue.NewMax[Player]()

	players := []Player{
		{Name: "Alice", Score: 1500},
		{Name: "Bob", Score: 2000},
		{Name: "Charlie", Score: 1800},
		{Name: "Diana", Score: 2200},
		{Name: "Eve", Score: 1600},
	}

	fmt.Println("Adding players to leaderboard:")
	for _, player := range players {
		leaderboard.Enqueue(player, player.Score)
		fmt.Printf("  * %s: %d points\n", player.Name, player.Score)
	}

	fmt.Println("\n* Final Leaderboard (Top to Bottom):")
	position := 1
	for !leaderboard.IsEmpty() {
		player, _ := leaderboard.Dequeue()
		medal := "3"
		if position == 1 {
			medal = "1"
		} else if position == 2 {
			medal = "2"
		}
		fmt.Printf("  %s %d. %s - %d points\n", medal, position, player.Name, player.Score)
		position++
	}
	fmt.Println()
}

func advancedOperationsExample() {
	fmt.Println("⚡ Example 5: Advanced Operations")
	fmt.Println("================================")

	pq := prioqueue.New[string]()

	// Add some items
	item1 := pq.Enqueue("Task A", 5)
	item2 := pq.Enqueue("Task B", 3)
	item3 := pq.Enqueue("Task C", 7)

	fmt.Printf("Initial queue size: %d\n", pq.Size())

	// Peek at top item
	if topItem, ok := pq.Peek(); ok {
		fmt.Printf("Top priority item: %s\n", topItem)
	}

	// Update priority of Task A to make it highest priority
	fmt.Println("Updating Task A priority from 5 to 1...")
	pq.UpdatePriority(item1, 1)

	// Peek again
	if topItem, ok := pq.Peek(); ok {
		fmt.Printf("New top priority item: %s\n", topItem)
	}

	// Show current priority of Task B
	fmt.Printf("Task B current priority: %d\n", item2.Priority)

	// Remove specific item
	fmt.Println("Removing Task C...")
	if removed, ok := pq.Remove(item3); ok {
		fmt.Printf("Removed: %s\n", removed)
	}

	fmt.Printf("Queue size after removal: %d\n", pq.Size())

	// Get all remaining items as slice
	items := pq.ToSlice()
	fmt.Println("Remaining items:")
	for _, item := range items {
		fmt.Printf("  - %s (priority: %d)\n", item.Value, item.Priority)
	}

	// Print queue representation
	fmt.Println("\nQueue string representation:")
	fmt.Println(pq.String())

	// Clear queue
	pq.Clear()
	fmt.Printf("Queue size after clear: %d\n", pq.Size())
	fmt.Printf("Is empty: %v\n", pq.IsEmpty())
}
