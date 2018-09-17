package service

import (
	"errors"
	"piktr/app/model"
	"piktr/app/repository"
	"time"
)

// Service invokes the repository
type Service struct {
	repo repository.UserRepository
}

// New create new service
func New(r repository.UserRepository) *Service {
	return &Service{
		repo: r,
	}
}

// Create ...
func (s *Service) Create(user *model.User) (*model.User, error) {
	user.ID = 1
	user.CreatedAt = time.Now()
	return s.repo.Create(user)
}

// FindByID ...
func (s *Service) FindByID(id int64) (*model.User, error) {
	return s.repo.FindByID(id)
}

// FindAll ...
func (s *Service) FindAll() ([]*model.User, error) {
	return s.repo.FindAll()
}

// Update ...
func (s *Service) Update(user *model.User) (*model.User, error) {
	user.ID = 1
	user.UpdatedAt = time.Now()
	return s.repo.Update(user)
}

// Delete ...
func (s *Service) Delete(id int64) error {
	res, err := s.FindByID(id)

	if err != nil {
		return err
	}

	if res.ID != id {
		return errors.New("User cannot be deleted")
	}
	return s.repo.Delete(id)
}
