# Boss fight
## Assignment
Build a production-ready Task Management REST API. You must follow a strict layered architecture to ensure the code is testable and maintainable. The application must interact with a PostgreSQL database running in a separate Docker container.

📁 **Project Architecture**
```text
project/           
├── main.go              # Entry point (Dependency Injection happens here)
├── docker-compose.yml   # Orchestrates the PostgreSQL container
├── schema.sql           # Database initialization script
├── model/               # Domain models
│   └── task.go          # Task struct definition
├── repository/          # Data access layer (SQL logic)
│   └── task_repo.go     # TaskRepository implementation
├── service/             # Business logic layer
│   └── task_service.go  # TaskService interface & implementation
├── delivery/            # Transport layer (HTTP)
│   └── http_handler.go  # net/http handlers
└── go.mod
```

### Requirements

1.  **The Model (`model/task.go`):**
    Define a `Task` struct with: `ID (int)`, `Title (string)`, `Description (string)`, and `CreatedAt (time.Time)`. Use appropriate JSON tags.

2.  **The Repository (`repository/task_repo.go`):**
    Implement the `TaskRepository` interface. This layer is the **only** place where SQL queries (`SELECT`, `INSERT`, `DELETE`) should exist.
    *   `FetchAll() ([]model.Task, error)`
    *   `Create(task model.Task) error`
    *   `GetByID(id int) (model.Task, error)`
    *   `Delete(id int) error`

3.  **The Service (`service/task_service.go`):**
    Implement the `TaskService` interface. This layer acts as a bridge. It should call the Repository. *Note: In a real app, this is where you would put validation logic (e.g., "Title cannot be empty").*

4.  **The Delivery (`delivery/http_handler.go`):**
    Use `net/http` to create a router and handlers.
    *   `GET /tasks` $\rightarrow$ returns list of tasks.
    *   `POST /tasks` $\rightarrow$ creates a task from JSON body.
    *   `GET /tasks/{id}` $\rightarrow$ returns a single task.
    *   `DELETE /tasks/{id}` $\rightarrow$ deletes a task.
    *   **Constraint:** All responses must be JSON. No plain text or HTML allowed.

5.  **Infrastructure (`docker-compose.yml`):**
    *   Spin up a `postgres` container.
    *   The Go application runs on your **local machine** (not in Docker), but it must connect to the database inside the Docker container.

---

### 📝 Deliverable
1. A working Go project.
2. A `docker-compose.yml` that launches a Postgres database.
3. A `schema.sql` to create the `tasks` table.
4. Ability to run `go run main.go` and perform all CRUD operations via `curl` or Postman.



> Note: Starter Code may have problem