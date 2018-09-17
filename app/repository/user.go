package repository

import "piktr/app/model"

// UserRepository defines the list
// of operation for user
type UserRepository interface {
	Create(user *model.User) (*model.User, error)
	FindByID(id int64) (*model.User, error)
	FindAll() ([]*model.User, error)
	Update(user *model.User) (*model.User, error)
	Delete(id int64) error
}
