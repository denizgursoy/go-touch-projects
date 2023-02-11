package services

type (
	ResourceService struct {
		r ResourceRepository
	}

	ResourceRepository interface {
		QueryResourceByID(id int64) (*Resource, error)
	}
)

func NewResourceService(repository ResourceRepository) *ResourceService {
	return &ResourceService{
		r: repository,
	}
}

func (r ResourceService) GetResourceByID(id int64) (*Resource, error) {
	return r.r.QueryResourceByID(id)

}
