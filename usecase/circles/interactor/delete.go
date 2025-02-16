package interactor

import (
	domain "go-arch-practice/domain/circles"
	"go-arch-practice/usecase/circles/ports"
)

type CircleDeleteInteractor struct {
	repo domain.ICircleRepository
}

var _ ports.ICircleDeleteInputPort = (*CircleDeleteInteractor)(nil)

func NewCircleDeleteInteractor(repo domain.ICircleRepository) *CircleDeleteInteractor {
	return &CircleDeleteInteractor{
		repo: repo,
	}
}

func (c *CircleDeleteInteractor) Handle(input *ports.CircleDeleteInputData) (*ports.CircleDeleteOutputData, error) {

	circle, err := c.repo.FindById(input.CircleID())
	if err != nil {
		return nil, err
	}

	err = c.repo.Delete(circle)
	if err != nil {
		return nil, err
	}

	return ports.NewCircleDeleteOutputData(), nil

}
