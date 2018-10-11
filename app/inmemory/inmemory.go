package inmemory

import (
	"piktr/app/domain/user"
	"sync"
	"time"
)

type userRepository struct {
	mu    *sync.Mutex
	users map[string]*user.User
}

// NewInMemoryUserRepository ...
func NewInMemoryUserRepository() user.Repository {
	return &userRepository{
		mu:    &sync.Mutex{},
		users: map[string]*user.User{},
	}
}

func (ur *userRepository) Create(u *user.User) error {
	ur.mu.Lock()
	defer ur.mu.Unlock()

	ur.users["1"] = &user.User{
		ID:        u.ID,
		Email:     u.Email,
		Username:  u.Username,
		CreatedAt: time.Now(),
		Password:  u.Password,
	}
	return nil
}

func (ur *userRepository) FindByID(id int64) (*user.User, error) {
	ur.mu.Lock()
	defer ur.mu.Unlock()

	for _, usr := range ur.users {
		if usr.ID == id {
			return usr, nil
		}
	}
	return nil, nil
}

func (ur *userRepository) FindAll() ([]*user.User, error) {
	ur.mu.Lock()
	defer ur.mu.Unlock()

	users := make([]*user.User, len(ur.users))

	for _, user := range ur.users {
		users = append(users, user)
	}

	return users, nil
}

func (ur *userRepository) Update(usr *user.User) error {
	ur.mu.Lock()
	defer ur.mu.Unlock()

	for _, user := range ur.users {
		if user.ID == usr.ID {
			user = usr
		}
	}

	return nil
}

func (ur *userRepository) Delete(id int64) error {
	ur.mu.Lock()
	defer ur.mu.Unlock()

	for _, user := range ur.users {
		if user.ID == id {
			user = nil
		}
	}
	return nil
}
