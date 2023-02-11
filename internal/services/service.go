package services

type (
	ResourceService struct{}
)

func NewResourceService() *ResourceService {
	return &ResourceService{}
}

func (r ResourceService) GetResourceByID(id int64) (*Resource, error) {
	return &Resource{"Resource-1"}, nil

}
