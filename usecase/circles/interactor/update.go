package interactor

import (
	"errors"
	domain "go-arch-practice/domain/circles"
	"go-arch-practice/usecase/circles/ports"
)

type CircleUpdateInteractor struct {
	repo    domain.ICircleRepository
	service domain.ICircleService
}

var _ ports.ICircleUpdatePort = (*CircleUpdateInteractor)(nil)

func (c *CircleUpdateInteractor) Handle(input *ports.CircleUpdateInputData) (*ports.CircleUpdateOutputData, error) {

	circle, err := c.repo.FindById(input.ID())
	if err != nil {
		return nil, err
	}

	if input.Name() != "" {

		circle.ChangeName(input.Name())
		exist, _ := c.service.Exists(circle)
		if exist {
			return nil, errors.New("既に存在しているサークルです")
		}

	}

	err = c.repo.Update(circle)
	if err != nil {
		return nil, err
	}

	return ports.NewCircleUpdateOutputData(), nil

}
