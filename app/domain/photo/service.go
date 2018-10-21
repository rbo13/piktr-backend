package photo

import (
	"crypto/sha256"
	"encoding/base64"
	"time"
)

// Service is an interface
// that defines the service
// that implements the repository
type Service interface {
	CreatePhoto(ph *Photo) error
	FindPhotoByID(id int64) (*Photo, error)
	FindAllPhotos() ([]*Photo, error)
	UpdatePhoto(ph *Photo) error
	DeletePhotoByID(id int64) error
}

type photoService struct {
	repo Repository
}

// NewUserService returns the Service
// interface defined
func NewUserService(repo Repository) Service {
	return &userService{
		repo,
	}
}

func (u *userService) CreateUser(user *User) error {
	user.CreatedAt = time.Now()
	user.Password = u.HashPassword(user.Password)
	return u.repo.Create(user)
}

func (u *userService) FindUserByID(id int64) (*User, error) {
	return u.repo.FindByID(id)
}

func (u *userService) FindAllUsers() ([]*User, error) {
	return u.repo.FindAll()
}

func (u *userService) UpdateUser(user *User) error {
	user.UpdatedAt = time.Now()
	return u.repo.Update(user)
}

func (u *userService) DeleteUserByID(id int64) error {
	return u.repo.Delete(id)
}

func (u *userService) HashPassword(plainPassword string) string {
	s := sha256.New()
	s.Write([]byte(plainPassword))
	return base64.URLEncoding.EncodeToString(s.Sum(nil))
}
