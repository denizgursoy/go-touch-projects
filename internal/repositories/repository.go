package repositories

import "{{.ModuleName}}/internal/services"

type (
	ResourceRepository struct{}
)

func NewResourceRepository() *ResourceRepository {
	return &ResourceRepository{}
}

func (r ResourceRepository) QueryResourceByID(id int64) (*services.Resource, error) {
	return &services.Resource{ID: id}, nil
}
