package main

import (
	"fmt"
	"log"

	"github.com/DKeshavarz/goraz/module3/json/user"
)

// UserStorage defines the behavior for our persistence layer
type UserStorage interface {
	Load() error
	Add(u user.User) error
	GetPrettyJSON(id int) (string, error)
}

func main() {
	
	// 1. Initialize implementation (Dependency Injection)
	// You will need to write the NewJSONRepo function in user/json_repo.go
	var repo UserStorage = nil/*user.NewJSONRepo("data/user.json")*/

	// 2. Load Data
	err := repo.Load()
	if err != nil {
		log.Fatalf("Failed to load data: %v", err)
	}

	// 3. Validate/Check data
	fmt.Println("Successfully loaded data from file.")

	// 4. Print specific user
	prettyUser, err := repo.GetPrettyJSON(1)
	if err != nil {
		fmt.Printf("Could not find user: %v\n", err)
	} else {
		fmt.Printf("User 1 Data:\n%s\n", prettyUser)
	}

	// 5. Add a new user
	newUser := user.User{/*ID: 3, Name: "Charlie", Email: "charlie@example.com"*/}

	err = repo.Add(newUser)
	if err != nil {
		log.Fatalf("Failed to add user: %v", err)
	}

	fmt.Println("New user added and file updated successfully.")
}
