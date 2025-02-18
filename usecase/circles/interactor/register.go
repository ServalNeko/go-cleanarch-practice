package interactor

import (
	"errors"
	domain_circle "go-arch-practice/domain/circles"
	domain_users "go-arch-practice/domain/users"
	"go-arch-practice/usecase/circles/ports"

	"github.com/google/uuid"
)

type CircleRegisterInteractor struct {
	circle_repo    domain_circle.ICircleRepository
	circle_service domain_circle.ICircleService
	user_repo      domain_users.IUserRepository
}

func NewCircleRegisterInteractor(
	circle_repo domain_circle.ICircleRepository,
	user_repo domain_users.IUserRepository,
	circle_service domain_circle.ICircleService,
) *CircleRegisterInteractor {
	return &CircleRegisterInteractor{
		circle_repo:    circle_repo,
		circle_service: circle_service,
		user_repo:      user_repo,
	}
}

var _ ports.ICircleRegisterInputPort = (*CircleRegisterInteractor)(nil)

func (c *CircleRegisterInteractor) Handle(input *ports.CircleRegisterInputData) (*ports.CircleRegisterOutputData, error) {

	owner, err := c.user_repo.FindById(input.OwnerID())
	if err != nil {
		return nil, err
	}
	ownerId := owner.ID()
	circleId, _ := domain_circle.NewCircleID(uuid.NewString())
	circle, _ := domain_circle.NewCircle(circleId, input.Name(), &ownerId, make([]domain_users.UserID, 0))

	exists, _ := c.circle_service.Exists(circle)
	if exists {
		return nil, errors.New("既に存在しているサークルです")
	}

	err = c.circle_repo.Save(circle)
	return ports.NewCircleCreateOutputData(circleId.Value()), err

}
