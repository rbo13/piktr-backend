package inmemory_test

import (
	"piktr/app/domain/user"
	"piktr/app/inmemory"
	"testing"
)

var userRepo user.Repository

func init() {
	userRepo = inmemory.NewInMemoryUserRepository()
}

func TestCreateUser(t *testing.T) {

	// arrange
	user := &user.User{
		ID:       int64(1),
		Email:    "sam@smith.com",
		Username: "sam_smith",
		Password: "notsecurepassword",
	}

	// act
	err := userRepo.Create(user)

	// assert
	if err != nil {
		t.Errorf("Error creating user: %v", err)
	}
}

func TestGetUserByID(t *testing.T) {

	// arrange
	id := int64(1)

	// act
	user, err := userRepo.FindByID(id)

	// assert
	if user == nil {
		t.Errorf("No user found by that ID: %v", user)
	}

	if err != nil {
		t.Errorf("Error finding a user: %v", err)
	}

	t.Log(user)
}
