package interactor

import (
	domain "go-arch-practice/domain/circles"
	"go-arch-practice/usecase/circles/ports"
)

type CircleGetAllInteractor struct {
	repo domain.ICircleRepository
}

var _ ports.ICircleGetAllInputPort = (*CircleGetAllInteractor)(nil)

func NewCircleGetAllInteractor(repo domain.ICircleRepository) *CircleGetAllInteractor {
	return &CircleGetAllInteractor{
		repo: repo,
	}
}

func (c *CircleGetAllInteractor) Handle() (*ports.CircleGetAllOutputData, error) {

	ciecles, err := c.repo.FindAll()
	if err != nil {
		return nil, err
	}

	data := make([]ports.CircleData, len(ciecles))
	for i, c := range ciecles {
		data[i] = *ports.NewCircleData(c.ID().String(), c.Name(), c.Owner().String(), c.MemberIDs())
	}
	output := ports.NewCircleGetAllOutputData(data)

	return output, nil
}
