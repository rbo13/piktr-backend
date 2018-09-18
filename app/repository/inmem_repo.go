package repository

import (
	"piktr/app/model"
)

// InMemRepo is the concrete implementation
// of the interface.
// Saving it to memory
type InMemRepo struct {
	data map[int64]*model.User
}

// NewInMemRepo create new repository
func NewInMemRepo() *InMemRepo {
	var data = map[int64]*model.User{}
	return &InMemRepo{
		data: data,
	}
}

// Create ...
func (r *InMemRepo) Create(user *model.User) (*model.User, error) {
	if user == nil {
		return nil, model.ErrNotInserted
	}
	r.data[user.ID] = user
	return r.data[user.ID], nil
}

// FindByID ...
func (r *InMemRepo) FindByID(id int64) (*model.User, error) {
	if id <= 0 {
		return nil, model.ErrIDIsRequired
	}

	if r.data[id] == nil {
		return nil, model.ErrNotFound
	}
	return r.data[id], nil
}

// FindAll ...
func (r *InMemRepo) FindAll() ([]*model.User, error) {
	var d []*model.User
	for _, j := range r.data {
		d = append(d, j)
	}
	return d, nil
}

// Update ...
func (r *InMemRepo) Update(user *model.User) (*model.User, error) {
	if user == nil {
		return nil, model.ErrNotUpdated
	}

	if r.data[user.ID] == user {
		r.data[user.ID] = user
		return r.data[user.ID], nil
	}

	return nil, model.ErrNotFound
}

// Delete ...
func (r *InMemRepo) Delete(id int64) error {
	if id <= 0 {
		return model.ErrIDIsRequired
	}
	r.data[id] = nil
	return nil
}
