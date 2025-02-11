package interactor

import (
	"go-arch-practice/domain/users"
	"go-arch-practice/usecase/users/ports"
)

type UserDeleteInteractor struct {
	userRepo users.IUserRepository
}

var _ ports.IUserDeleteInputPort = (*UserDeleteInteractor)(nil)

func NewUserDeleteInteractor(userRepo users.IUserRepository) *UserDeleteInteractor {
	return &UserDeleteInteractor{
		userRepo: userRepo,
	}
}

func (u *UserDeleteInteractor) Handle(input *ports.UserDeleteInputData) (*ports.UserDeleteOutputData, error) {
	user, err := u.userRepo.FindById(input.ID())
	if err != nil {
		return nil, err
	}

	err = u.userRepo.Delete(user)
	if err != nil {
		return nil, err
	}

	return ports.NewUserDeleteOutputData(user.ID()), nil
}
