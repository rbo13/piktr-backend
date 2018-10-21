package photo

import (
	"time"
)

// Photo represents
// a photo of our application
type Photo struct {
	ID        int64      `json:"id" db:"id"`
	UniqueURL string     `json:"unique_url" db:"unique_url"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"update_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

// TableName returns a string
// for the name of our table
// in pluralized form
func (Photo) TableName() string {
	return "photos"
}
