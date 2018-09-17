package service_test

import (
	"piktr/app/model"
	"piktr/app/repository"
	"piktr/app/service"
	"testing"
)

func TestCreate(t *testing.T) {
	repo := repository.NewInMemRepo()
	service := service.New(repo)
	userData := &model.User{
		Email:    "rborofeo",
		Username: "rbo13",
		Password: "notasecurepassword",
	}
	user, err := service.Create(userData)

	if err != nil {
		t.Fatal(err)
	}
	if user == nil {
		t.Errorf("user is empty: %v", user)
	}
	t.Log(user)
}
