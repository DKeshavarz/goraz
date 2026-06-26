package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

// User represents a user in our system
type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

// UserRepository defines the interface for user data operations
type userRepository interface {
	// Save stores a new user. Returns error if user already exists.
	Save(user User) error

	// FindByID retrieves a user by ID. Returns error if not found.
	FindByID(id int) (User, error)

	// Delete removes a user by ID. Returns error if not found.
	Delete(id int) error

	// List returns all users. Returns empty slice if none exist.
	List() ([]User, error)
}

func main() {

	// Choose implementation
	fmt.Println("\nSelect repository implementation:")
	fmt.Println("  1. Slice-based Repository")
	fmt.Println("  2. Map-based Repository")



	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	var repo userRepository
	switch choice {
	case "1":
		fmt.Println("\nSLICE REPOSITORY")
		// repo = sliceRepo.New()
	case "2":
		fmt.Println("\nMAP REPOSITORY")
		// repo = mapRepo.New()

	default:
		fmt.Println("Invalid choice. You disappointed golang")
	}

	runTests(repo)
}

func runTests(repo userRepository) {
	fmt.Println("=== RUNNING TESTS ===")

	// Test 0: check null
	fmt.Println("📋 Test 0: check list is null")
	if repo == nil {
		fmt.Println("❌ Repo is Null")
		return
	}

	// Test 1: List empty repository (should return empty slice)
	fmt.Println("📋 Test 1: List empty repository")
	users, err := repo.List()
	if err != nil {
		fmt.Printf("  ❌ List failed: %v\n", err)
	} else if len(users) == 0 {
		fmt.Println("  ✅ Empty list returned correctly")
	} else {
		fmt.Printf("  ❌ Expected empty list, got %d users\n", len(users))
	}

	// Test 2: Add three users
	fmt.Println("\n📝 Test 2: Add three users")
	user1 := User{ID: 1, Name: "Alice", Email: "alice@email.com", Age: 28}
	user2 := User{ID: 2, Name: "Bob", Email: "bob@email.com", Age: 34}
	user3 := User{ID: 3, Name: "Charlie", Email: "charlie@email.com", Age: 22}

	err = repo.Save(user1)
	if err != nil {
		fmt.Printf("  ❌ Failed to save Alice: %v\n", err)
	} else {
		fmt.Println("  ✅ Saved Alice (ID: 1)")
	}

	err = repo.Save(user2)
	if err != nil {
		fmt.Printf("  ❌ Failed to save Bob: %v\n", err)
	} else {
		fmt.Println("  ✅ Saved Bob (ID: 2)")
	}

	err = repo.Save(user3)
	if err != nil {
		fmt.Printf("  ❌ Failed to save Charlie: %v\n", err)
	} else {
		fmt.Println("  ✅ Saved Charlie (ID: 3)")
	}

	// Test 3: Get one user by ID
	fmt.Println("\n🔍 Test 3: Find user by ID (ID: 2)")
	foundUser, err := repo.FindByID(2)
	if err != nil {
		fmt.Printf("  ❌ Failed to find user: %v\n", err)
	} else if foundUser == user2 {
		fmt.Printf("  ✅ Found correct user: %s (ID: %d)\n", foundUser.Name, foundUser.ID)
	} else {
		fmt.Printf("  ❌ Found wrong user. Expected %+v, got %+v\n", user2, foundUser)
	}

	// Test 4: Delete the user
	fmt.Println("\n🗑️  Test 4: Delete user (ID: 2)")
	err = repo.Delete(2)
	if err != nil {
		fmt.Printf("  ❌ Failed to delete: %v\n", err)
	} else {
		fmt.Println("  ✅ Deleted user ID: 2")
	}

	// Test 5: Find the deleted user (should fail)
	fmt.Println("\n🔍 Test 5: Find deleted user (ID: 2)")
	_, err = repo.FindByID(2)
	if err != nil {
		fmt.Printf("  ✅ Correctly got error: %v\n", err)
	} else {
		fmt.Println("  ❌ Should have returned error for deleted user")
	}

	// Test 6: Delete the deleted user again (should fail)
	fmt.Println("\n🗑️  Test 6: Delete deleted user again (ID: 2)")
	err = repo.Delete(2)
	if err != nil {
		fmt.Printf("  ✅ Correctly got error: %v\n", err)
	} else {
		fmt.Println("  ❌ Should have returned error for already deleted user")
	}

	// Test 7: List all users and verify
	fmt.Println("\n📋 Test 7: List all users and verify")
	allUsers, err := repo.List()
	if err != nil {
		fmt.Printf("  ❌ List failed: %v\n", err)
	} else {
		expectedUsers := []User{user1, user3}
		if len(allUsers) == len(expectedUsers) {
			fmt.Printf("  ✅ Correct number of users: %d\n", len(allUsers))

			// Check each user matches
			match := true
			for i, expected := range expectedUsers {
				if i >= len(allUsers) {
					match = false
					break
				}
				if !reflect.DeepEqual(allUsers[i], expected) {
					match = false
					break
				}
			}

			if match {
				fmt.Println("  ✅ All users match expected:")
				for _, u := range allUsers {
					fmt.Printf("     • ID: %d | Name: %s | Email: %s | Age: %d\n",
						u.ID, u.Name, u.Email, u.Age)
				}
			} else {
				fmt.Println("  ❌ Users don't match expected")
				fmt.Printf("  Expected: %+v\n", expectedUsers)
				fmt.Printf("  Got: %+v\n", allUsers)
			}
		} else {
			fmt.Printf("  ❌ Wrong number of users. Expected %d, got %d\n",
				len(expectedUsers), len(allUsers))
		}
	}

	// Test 8: Add user with duplicate ID (should fail)
	fmt.Println("\n📝 Test 8: Add user with duplicate ID")
	duplicateUser := User{ID: 1, Name: "Duplicate", Email: "dup@email.com", Age: 99}
	err = repo.Save(duplicateUser)
	if err != nil {
		fmt.Printf("  ✅ Correctly rejected duplicate ID: %v\n", err)
	} else {
		fmt.Println("  ❌ Should have rejected duplicate ID")
	}

	fmt.Println("=== ALL TESTS COMPLETED ===")
}
