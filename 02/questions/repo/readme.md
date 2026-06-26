# User Repository Implementation
## Assignment
Implement the UserRepository interface using two different storage strategies:
- Slice-based (sliceRepo package)
- Map-based (mapRepo package)

📁 Project Architecture
```
project/
├── user.go              # Interface + User struct (GIVEN)
├── main.go              # Test suite (GIVEN)
├── sliceRepo/
│   └── slice_repo.go    # Slice implementation (YOU WRITE)
├── mapRepo/
│   └── map_repo.go      # Map implementation (YOU WRITE)
└── go.mod               # Module definition
```