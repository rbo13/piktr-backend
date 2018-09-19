package db

import (
	"github.com/jinzhu/gorm"
)

// MySQLDB defines
// the SQL database
type MySQLDB struct {
	db *gorm.DB
}

// New returns a pointer
// to MySQLDB
func New(db *gorm.DB) *MySQLDB {
	return &MySQLDB{
		db: db,
	}
}
