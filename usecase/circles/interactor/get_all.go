package interactor

import (
	domain "go-arch-practice/domain/circles"
	"go-arch-practice/usecase/circles/ports"
)

type CiecleGetAllInteractor struct {
	repo domain.ICircleRepository
}

var _ ports.ICircleGetAllInputPort = (*CiecleGetAllInteractor)(nil)

func NewCiecleGetAllInteractor(repo domain.ICircleRepository) *CiecleGetAllInteractor {
	return &CiecleGetAllInteractor{
		repo: repo,
	}
}

func (c *CiecleGetAllInteractor) Handle() (*ports.CircleGetAllOutputData, error) {

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
