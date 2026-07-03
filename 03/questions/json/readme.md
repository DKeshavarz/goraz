# JSON File Storage 

## Assignment
Implement the `UserStorage` interface to handle persistent JSON storage. Your implementation must encapsulate all file I/O operations and JSON marshaling logic.

📁 **Project Architecture**
```text
project/           
├── main.go               # Orchestrator (Interface defined here)
├── user/
│   ├── model.go          # User struct definition
│   └── json_repo.go      # Implementation (YOU WRITE)
├── data/
│   └── user.json         # Already given
└── go.mod                # Module definition
```

### Requirements

1.  **The Interface (Defined in `main.go`):**
    Your implementation must satisfy the following contract:
    ```go
    type UserStorage interface {
        Load() error                               // Read from file
        Add(u User) error                          // Append user & Save to file
        GetPrettyJSON(id int) (string, error)      // Return pretty-printed string for specific user
    }
    ```

2.  **Implementation Logic (`user/json_repo.go`):**
    *   Implement the `Load()` method to unmarshal `data/user.json` into an internal slice.
    *   Implement `Add()` to update the slice and rewrite the file.
    *   Implement `GetPrettyJSON()` to locate a user by ID and return their data formatted with `json.MarshalIndent`.




### 💡 Tips for implementation:
*   **Struct Tags:** Don't forget your struct tags in `user/model.go` (e.g., `json:"name"`, `json:"age,omitempty"`).
*   **Error Handling:** In `GetPrettyJSON`, return a custom error if the ID requested does not exist in your slice.
*   **JSON Indentation:** Use `json.MarshalIndent(v, "", "  ")` to satisfy the requirement for "pretty" format.