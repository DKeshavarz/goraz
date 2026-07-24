package main

import (
	"fmt"
	"goraz/module2/sorter/user"
)

// SortStrategy defines the interface for sorting strategies
type SortStrategy interface {
	Sort(users []user.User) []user.User
}

// Context
type UserSorter struct {
	strategy SortStrategy
}

func NewUserSorter(strategy SortStrategy) *UserSorter {
	return &UserSorter{strategy: strategy}
}

func (us *UserSorter) SetStrategy(strategy SortStrategy) {
	us.strategy = strategy
}

func (us *UserSorter) ExecuteSort(users []user.User) []user.User {
	return us.strategy.Sort(users)
}

// ========= MAIN ===============

func main() {
	// Sample users
	users := []user.User{
		{ID: 3, Name: "Charlie", Email: "charlie@email.com", Age: 22},
		{ID: 1, Name: "Alice", Email: "alice@email.com", Age: 28},
		{ID: 4, Name: "Diana", Email: "diana@email.com", Age: 30},
		{ID: 2, Name: "Bob", Email: "bob@email.com", Age: 34},
	}

	printUsers("Original Users", users)

	// Create sorter with default strategy
	sorter := NewUserSorter(nil /*...*/)

	// Sort by Name
	sortedByName := sorter.ExecuteSort(users)
	printUsers("Sorted by Name (A-Z)", sortedByName)

	// Change strategy to SortByAge
	// sorter.SetStrategy(...)
	sortedByAge := sorter.ExecuteSort(users)
	printUsers("Sorted by Age (Ascending)", sortedByAge)

	// Change strategy to SortByID
	// sorter.SetStrategy(...)
	sortedByID := sorter.ExecuteSort(users)
	printUsers("Sorted by ID (Ascending)", sortedByID)

	// Change strategy to SortByEmail
	// sorter.SetStrategy(...)
	sortedByEmail := sorter.ExecuteSort(users)
	printUsers("Sorted by Email (A-Z)", sortedByEmail)

	fmt.Println("\n✅ All sorting strategies completed!")
}

// ========== HELPER FUNCTIONS ==========

func printUsers(title string, users []user.User) {
	fmt.Printf("\n%s:\n", title)
	fmt.Println("─────────────────────────────────────")
	for _, u := range users {
		fmt.Printf("ID: %d | Name: %-10s | Age: %2d | Email: %s\n",
			u.ID, u.Name, u.Age, u.Email)
	}
}
