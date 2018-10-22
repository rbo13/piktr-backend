package photo

import (
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

// NewPhotoService returns the Service
// interface defined
func NewPhotoService(repo Repository) Service {
	return &photoService{
		repo,
	}
}

func (p *photoService) CreatePhoto(ph *Photo) error {
	ph.CreatedAt = time.Now()
	ph.UniqueURL = p.GenerateUniqueURL()
	return p.repo.Create(ph)
}

func (p *photoService) FindPhotoByID(id int64) (*Photo, error) {
	return p.repo.FindByID(id)
}

func (p *photoService) FindAllPhotos() ([]*Photo, error) {
	return p.repo.FindAll()
}

func (p *photoService) UpdatePhoto(ph *Photo) error {
	ph.UpdatedAt = time.Now()
	return p.repo.Update(ph)
}

func (p *photoService) DeletePhotoByID(id int64) error {
	return p.repo.Delete(id)
}

func (p *photoService) GenerateUniqueURL() string {
	return ""
}
