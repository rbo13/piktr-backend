package repository

import (
	"piktr/app/model"

	"github.com/jinzhu/gorm"
)

// SQLDb is the concrete
// implementation of the
// UserRepository interface
// using mysql
type SQLDb struct {
	db *gorm.DB
}

// NewSQLDb returns
// a pointer to SQLDb
func NewSQLDb(db *gorm.DB) *SQLDb {
	return &SQLDb{
		db: db,
	}
}

// Create ...
func (sql *SQLDb) Create(user *model.User) (*model.User, error) {
	if user == nil {
		return nil, model.ErrNotInserted
	}
	err := sql.db.Debug().Create(&user).Error

	if err != nil {
		return nil, model.ErrNotInserted
	}
	return user, nil
}

// FindByID ...
func (sql *SQLDb) FindByID(id int64) (*model.User, error) {
	user := new(model.User)

	if id <= 0 {
		return nil, model.ErrIDIsRequired
	}

	err := sql.db.Debug().First(&user, id).Error

	if err != nil {
		return nil, model.ErrNotInserted
	}
	return user, nil
}

// FindAll ...
func (sql *SQLDb) FindAll() ([]*model.User, error) {
	var user []*model.User

	err := sql.db.Debug().Find(&user).Error

	if err != nil {
		return nil, model.ErrNotInserted
	}
	return user, nil
}

// Update ...
func (sql *SQLDb) Update(user *model.User) (*model.User, error) {
	if user == nil {
		return nil, model.ErrNotInserted
	}
	err := sql.db.Save(&user).Error
	if err != nil {
		return nil, model.ErrNotUpdated
	}
	return user, nil
}

// Delete ...
func (sql *SQLDb) Delete(id int64) error {
	if id <= 0 {
		return model.ErrIDIsRequired
	}

	// lets get the current user first
	user, err := sql.FindByID(id)
	if err != nil {
		return model.ErrNotFound
	}

	return sql.db.Debug().Delete(&user, id).Error
}
