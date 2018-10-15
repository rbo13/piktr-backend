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

	if user != nil {
		t.Log(user)
	}

}

func TestFindAllUsers(t *testing.T) {
	// arrange

	// act
	users, err := userRepo.FindAll()
	// assert
	if err != nil {
		t.Errorf("Error occurred due to: %v", err)
	}
	if users == nil {
		t.Errorf("Users cannot be empty: %v", users)
	}
	if len(users) <= 0 {
		t.Errorf("Users length is: %d", len(users))
	}
	if len(users) > 0 {
		t.Log(&users)
	}
}

func TestUpdate(t *testing.T) {
	id := 1

	user := &user.User{
		ID:       int64(id),
		Email:    "sam@smith.com",
		Username: "richard_burk",
	}

	err := userRepo.Update(user)

	if err != nil {
		t.Errorf("Error updating due to: %v", err)
	}
}
