package interactor

import (
	domain "go-arch-practice/domain/circles"
	"go-arch-practice/usecase/circles/ports"
)

type CircleGetInteractor struct {
	circleRepo domain.ICircleRepository
}

var _ ports.ICircleGetInputPort = (*CircleGetInteractor)(nil)

func NewCircleGetInteractor(circleRepo domain.ICircleRepository) *CircleGetInteractor {
	return &CircleGetInteractor{
		circleRepo: circleRepo,
	}
}

func (c *CircleGetInteractor) Handle(input *ports.CircleGetInputData) (*ports.CircleGetOutputData, error) {
	circle, err := c.circleRepo.FindById(input.ID())
	if err != nil {
		return nil, err
	}
	data := ports.NewCircleData(circle.ID().Value(), circle.Name(), circle.Owner().Value(), circle.MemberIDs())

	return ports.NewCircleGetOutputData(data), nil
}
