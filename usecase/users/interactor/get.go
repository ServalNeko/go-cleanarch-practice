package interactor

import (
	domain "go-arch-practice/domain/users"
	"go-arch-practice/usecase/users/ports"
)

type UserGetInteractor struct {
	userRepo domain.IUserRepository
}

var _ ports.IUserGetInputPort = (*UserGetInteractor)(nil)

func NewUserGetInteractor(userRepo domain.IUserRepository) *UserGetInteractor {
	return &UserGetInteractor{
		userRepo: userRepo,
	}
}

func (u *UserGetInteractor) Handle(input *ports.UserGetInputData) (*ports.UserGetOutputData, error) {
	user, err := u.userRepo.FindById(input.ID())
	if err != nil {
		return nil, err
	}

	var data = ports.NewUserData(user.ID().Value(), user.Name())

	return ports.NewUserGetOutputData(data), nil
}
