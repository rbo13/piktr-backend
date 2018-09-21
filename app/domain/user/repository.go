package user

// Repository defines
// the list of operations
// for the user.
type Repository interface {
	Create(user *User) error
	FindByID(id int64) (*User, error)
	FindAll() ([]*User, error)
	Update(user *User) error
	Delete(id int64) error
}
