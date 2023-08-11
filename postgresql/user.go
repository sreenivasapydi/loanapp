package postgresql

import (
	"context"
	"fmt"
	"sort"
	"strconv"

	"loanapp"

	"github.com/icrowley/fake"
)

// Ensure service implements interface.
var _ loanapp.UserService = (*UserService)(nil)

// UserService represents a service for managing users.
type UserService struct {
	db    *DB
	users map[string]*loanapp.User
}

// NewUserService returns a new instance of UserService.
func NewUserService(db *DB) *UserService {
	s := &UserService{
		db:    db,
		users: make(map[string]*loanapp.User),
	}
	loadUsers(s)
	return s
}

// FindUserByID retrieves a user by ID along with their associated auth objects.
// Returns ENOTFOUND if user does not exist.
func (s *UserService) FindUserByID(ctx context.Context, id string) (*loanapp.User, error) {
	user := s.users[id]
	return user, nil
}

// FindUsers retrieves a list of users by filter. Also returns total count of
// matching users which may differ from returned results if filter.Limit is specified.
func (s *UserService) FindUsers(ctx context.Context, filter loanapp.UserFilter) ([]*loanapp.User, int, error) {
	var userIds = make([]string, 0, len(s.users))
	for _, u := range s.users {
		userIds = append(userIds, u.ID)
	}
	sort.Strings(userIds)
	var users = make([]*loanapp.User, 0, len(s.users))
	for _, uid := range userIds {
		users = append(users, s.users[uid])
	}

	return users, 0, nil
}

// CreateUser creates a new user. This is only used for testing since users are
// typically created during the OAuth creation process in AuthService.CreateAuth().
func (s *UserService) CreateUser(ctx context.Context, user *loanapp.User) (*loanapp.User, error) {
	s.users[user.ID] = user
	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id string, upd loanapp.UserUpdate) (*loanapp.User, error) {
	user, exists := s.users[id]

	if !exists {
		return &loanapp.User{}, &loanapp.Error{Code: loanapp.ENOTFOUND, Message: "User not found"}
	}
	// update fields
	if v := upd.Name; v != nil {
		user.Name = *v
	}
	if v := upd.Email; v != nil {
		user.Email = *v
	}

	return user, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	fmt.Printf("=== user user %s", id)

	_, exists := s.users[id]
	fmt.Printf("=== user user %s %v", id, exists)

	if !exists {
		return &loanapp.Error{Code: loanapp.ENOTFOUND, Message: "User not found"}
	}

	delete(s.users, id)
	return nil
}

func (s *UserService) Ping(ctx context.Context, id string) (string, error) {
	fmt.Printf("==== hello %s", id)
	return id, nil
}

func loadUsers(s *UserService) {
	for i := 101; i <= 110; i++ {
		var u loanapp.User
		u.ID = strconv.Itoa(i)
		u.Name = fake.FirstName()
		u.Email = fake.EmailAddress()
		s.users[u.ID] = &u
	}
}
