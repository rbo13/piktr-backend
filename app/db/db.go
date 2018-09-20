package db

import (
	"log"
	"piktr/app/model"

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

// Setup ...
func Setup(db *gorm.DB, dbName string) error {
	err := db.Debug().Exec("CREATE DATABASE IF NOT EXISTS " + dbName + " DEFAULT CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci;").Error
	if err != nil {
		return err
	}
	err = db.Debug().Exec("USE " + dbName + ";").Error

	if err != nil {
		return err
	}

	return nil
}

// MigrateTables handles migrating of table
func (sql *MySQLDB) MigrateTables() {
	log.Println("Migrating Tables...")
	sql.db.AutoMigrate(&model.User{})
}
