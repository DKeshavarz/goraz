package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/DKeshavarz/goraz/module3/taskmanger/delivery"
	"github.com/DKeshavarz/goraz/module3/taskmanger/repository"
	"github.com/DKeshavarz/goraz/module3/taskmanger/service"
	_ "github.com/lib/pq"
)

func main() {
	// 1. DB connection setup
	connStr := "user=postgres password=mysecretpassword dbname=tasksdb host=localhost port=5432 sslmode=disable" // Use you credential
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}

	// 2. Wire Layers
	repo := repository.NewPostgresRepo(db)
	svc := service.NewTaskService(repo)
	handler := delivery.NewTaskHandler(svc)

	// 3. Routing (Splitting collection vs. individual member endpoints)
	http.HandleFunc("/tasks", handler.TasksHandler)       // Matches exactly "/tasks"
	http.HandleFunc("/tasks/", handler.TaskMemberHandler) // Matches "/tasks/<id>"

	// 4. Start Server
	fmt.Println("Server running on port :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
