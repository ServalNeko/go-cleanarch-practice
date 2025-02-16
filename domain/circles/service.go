package circles

type ICircleService interface {
	Exists(c *Circle) (bool, error)
}

type CircleService struct {
	repo ICircleRepository
}

var _ ICircleService = (*CircleService)(nil)

func NewCircleService(repo ICircleRepository) *CircleService {
	return &CircleService{
		repo: repo,
	}
}

func (s *CircleService) Exists(c *Circle) (bool, error) {

	t, _ := s.repo.FindByName(c.Name())
	return t != nil, nil
}
