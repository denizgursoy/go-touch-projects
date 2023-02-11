package repositories

type (
	ResourceRepository struct{}
)

func NewResourceRepository() *ResourceRepository {
	return &ResourceRepository{}
}
