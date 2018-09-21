package mysql

import (
	"errors"
	"piktr/app/domain/user"

	"github.com/jinzhu/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// NewMySQLUserRepository returns the interface
// of user.Repository
func NewMySQLUserRepository(db *gorm.DB) user.Repository {
	return &userRepository{
		db,
	}
}

func (u *userRepository) Create(user *user.User) error {
	return u.db.Debug().Create(&user).Error
}

func (u *userRepository) FindByID(id int64) (*user.User, error) {
	user := user.User{}
	if id <= 0 {
		return nil, errors.New("ID is required")
	}

	err := u.db.Debug().First(&user, id).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) FindAll() ([]*user.User, error) {
	users := []*user.User{}

	err := u.db.Debug().Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepository) Update(user *user.User) error {
	if user == nil {
		return errors.New("User is required")
	}

	err := u.db.Debug().Save(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepository) Delete(id int64) error {
	if id <= 0 {
		return errors.New("UserID is required")
	}

	// Find a user first
	user, err := u.FindByID(id)

	if err != nil && user == nil {
		return errors.New("User not found")
	}

	return u.db.Debug().Delete(&user).Error
}
