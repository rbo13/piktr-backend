package mysql

import (
	"database/sql"
	"testing"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func recordStats(db *sql.DB, userID int64) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	if _, err = tx.Exec("SELECT * FROM users WHERE id = ?", userID); err != nil {
		return
	}
	return
}

func TestShouldGetUser(t *testing.T) {
	// open database stub
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("SELECT (.+) FROM users WHERE id (.+)").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// now we execute our method
	if err = recordStats(db, 1); err != nil {
		t.Errorf("error looking for user with that id of %d: %v", 1, err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}

}
