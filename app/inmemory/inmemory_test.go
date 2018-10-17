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

	table := []struct {
		input    *user.User
		expected error
	}{
		{
			&user.User{
				ID:       int64(1),
				Email:    "sam@smith.com",
				Username: "sam_smith",
				Password: "notsecurepassword",
			}, nil},
	}

	// arrange
	for _, tt := range table {
		actual := userRepo.Create(tt.input)

		if actual != tt.expected {
			t.Errorf("Expected to be: %v, but got %v", tt.expected, actual)
		}
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

func TestDelete(t *testing.T) {
	id := int64(1)

	err := userRepo.Delete(id)

	if err != nil {
		t.Errorf("Error updating due to: %v", err)
	}
}
