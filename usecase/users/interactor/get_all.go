package interactor

import (
	"go-arch-practice/usecase/users"
	"go-arch-practice/usecase/users/ports"
)

type UserGetAllInteractor struct {
	userRepo users.IUserRepository
}

var _ ports.IUserGetAllInputPort = (*UserGetAllInteractor)(nil)

func NewUserGetAllInteractor(userRepo users.IUserRepository) *UserGetAllInteractor {
	return &UserGetAllInteractor{
		userRepo: userRepo,
	}
}

func (u *UserGetAllInteractor) Handle() (*ports.UserGetAllOutputData, error) {
	users, err := u.userRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var data = make([]ports.UserData, len(users))
	for i, user := range users {
		data[i] = *ports.NewUserData(user.ID().Value(), user.Name())
	}

	return ports.NewUserGetAllOutputData(data), nil
}
