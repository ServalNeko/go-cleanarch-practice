package interactor

import (
	domain "go-arch-practice/domain/users"
	"go-arch-practice/usecase/users/ports"
)

type UserUpdateInteractor struct {
	userRepo domain.IUserRepository
}

var _ ports.IUserUpdateInputPort = (*UserUpdateInteractor)(nil)

func NewUserUpdateInteractor(userRepo domain.IUserRepository) *UserUpdateInteractor {
	return &UserUpdateInteractor{
		userRepo: userRepo,
	}
}

func (u *UserUpdateInteractor) Handle(input *ports.UserUpdateInputData) (*ports.UserUpdateOutputData, error) {
	user, err := u.userRepo.FindById(input.ID())
	if err != nil {
		return nil, err
	}

	user.ChangeName(input.Name())

	err = u.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	return ports.NewUserUpdateOutputData(), nil
}
