package user

import (
	"time"
)

// User represents
// a user of our application
type User struct {
	ID        int64      `json:"id" db:"id"`
	Email     string     `json:"email" sql:"type:varchar(100)" gorm:"unique_index" db:"email"`
	Username  string     `json:"username" sql:"type:varchar(50)" gorm:"unique_index" db:"username"`
	Password  string     `json:"-" db:"password"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"update_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

// TableName returns a string
// for the name of our table
// in pluralized form
func (User) TableName() string {
	return "users"
}
