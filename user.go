package loanapp

import (
	"context"
)

// User represents a user in the system. Users are typically created via OAuth
// using the AuthService but users can also be created directly for testing.
type User struct {
	ID string `json:"id"`

	// User's preferred name & email.
	Name  string `json:"name"`
	Email string `json:"email"`

	// Timestamps for user creation & last update.
	// 	CreatedAt time.Time `json:"createdAt"`
	// 	UpdatedAt time.Time `json:"updatedAt"`
}

// Validate returns an error if the user contains invalid fields.
// This only performs basic validation.
func (u *User) Validate() error {
	if u.Name == "" {
		return Errorf(EINVALID, "User name required.")
	}
	return nil
}

// UserService represents a service for managing users.
type UserService interface {
	// Retrieves a user by ID.
	// Returns ENOTFOUND if user does not exist.
	FindUserByID(ctx context.Context, id string) (*User, error)

	// Retrieves a list of users by filter. Also returns total count of matching
	// users which may differ from returned results if filter.Limit is specified.
	FindUsers(ctx context.Context, filter UserFilter) ([]*User, int, error)

	// Creates a new user.
	CreateUser(ctx context.Context, user *User) (*User, error)

	// Updates a user object. Returns EUNAUTHORIZED if current user is not
	// the user that is being updated. Returns ENOTFOUND if user does not exist.
	UpdateUser(ctx context.Context, id string, upd UserUpdate) (*User, error)

	// Permanently deletes a user. Returns EUNAUTHORIZED
	// if current user is not the user being deleted. Returns ENOTFOUND if
	// user does not exist.
	DeleteUser(ctx context.Context, id string) error

	Ping(ctx context.Context, id string) (string, error)
}

// UserFilter represents a filter passed to FindUsers().
type UserFilter struct {
	// Filtering fields.
	ID     *string `json:"id"`
	Email  *string `json:"email"`
	APIKey *string `json:"apiKey"`

	// Restrict to subset of results.
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

// UserUpdate represents a set of fields to be updated via UpdateUser().
type UserUpdate struct {
	Name  *string `json:"name"`
	Email *string `json:"email"`
}
