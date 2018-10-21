package photo

// Repository defines
// the list of operations
// for the user.
type Repository interface {
	Create(user *Photo) error
	FindByID(id int64) (*Photo, error)
	FindAll() ([]*Photo, error)
	Update(user *Photo) error
	Delete(id int64) error
}
