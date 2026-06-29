package user

type UserRepository interface {
	// Save stores a new user. Returns error if user already exists.
	Save(user User) error

	// FindByID retrieves a user by ID. Returns error if not found.
	FindByID(id int) (User, error)

	// Delete removes a user by ID. Returns error if not found.
	Delete(id int) error

	// List returns all users. Returns empty slice if none exist.
	List() ([]User, error)
}