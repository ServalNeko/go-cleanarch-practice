package interactor

import (
	circle_domain "go-arch-practice/domain/circles"
	user_domain "go-arch-practice/domain/users"
	"go-arch-practice/usecase/circles/ports"
)

type CircleJoinInteractor struct {
	circle_repo circle_domain.ICircleRepository
	user_repo   user_domain.IUserRepository
}

var _ ports.ICircleJoinInputPort = (*CircleJoinInteractor)(nil)

func NewCircleJoinInteractor(
	circle_repo circle_domain.ICircleRepository,
	user_repo user_domain.IUserRepository,
) *CircleJoinInteractor {
	return &CircleJoinInteractor{
		circle_repo: circle_repo,
		user_repo:   user_repo,
	}
}

func (c *CircleJoinInteractor) Handle(input *ports.CircleJoinInputData) (*ports.CircleJoinOutputData, error) {

	circle, err := c.circle_repo.FindById(input.CircleID())
	if err != nil {
		return nil, err
	}

	user, err := c.user_repo.FindById(input.MemberID())
	if err != nil {
		return nil, err
	}

	fullSpec := circle_domain.NewCircleFullSpecification(c.user_repo)
	err = circle.Join(*user, fullSpec)
	if err != nil {
		return nil, err
	}

	err = c.circle_repo.Save(circle)

	return ports.NewCircleJoinOutputData(), err
}
