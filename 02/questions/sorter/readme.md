# Sorting Strategies with Strategy Pattern

## Problem

Implement different sorting strategies for a list of users using the Strategy Pattern.

You need to create sorting strategies that can sort users by:
- Name (alphabetical)
- Age (ascending)
- ID (ascending)
- Email (alphabetical)


## Architecture

```
project/
├── main.go              # GIVEN - SortStrategy, main()
├── user/
│   └── user.go          # GIVEN - Contains User
├── sorter/
│   └── sorter.go        # YOU WRITE - Strategy implementations
└── go.mod
```

## Task

Create a `sorter` package that implements the `SortStrategy` interface:

```go
type SortStrategy interface {
    Sort(users []User) []User
}
```

Implement these strategies:
- `SortByName`
- `SortByAge`
- `SortByID`
- `SortByEmail`



## Example Output

```
Original Users:
─────────────────────────────────────
ID: 3 | Name: Charlie   | Age: 22 | Email: charlie@email.com
ID: 1 | Name: Alice     | Age: 28 | Email: alice@email.com
ID: 4 | Name: Diana     | Age: 30 | Email: diana@email.com
ID: 2 | Name: Bob       | Age: 34 | Email: bob@email.com

Sorted by Name (A-Z):
─────────────────────────────────────
ID: 1 | Name: Alice     | Age: 28 | Email: alice@email.com
ID: 2 | Name: Bob       | Age: 34 | Email: bob@email.com
ID: 3 | Name: Charlie   | Age: 22 | Email: charlie@email.com
ID: 4 | Name: Diana     | Age: 30 | Email: diana@email.com
```
