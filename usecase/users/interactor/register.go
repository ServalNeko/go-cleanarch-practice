package interactor

import (
	domain "go-arch-practice/domain/users"
	"go-arch-practice/usecase/users"
	"go-arch-practice/usecase/users/ports"

	"github.com/google/uuid"
)

type UserRegisterInteractor struct {
	userRepo users.IUserRepository
}

var _ ports.IUserRegisterInputPort = (*UserRegisterInteractor)(nil)

func NewUserRegisterInteractor(userRepo users.IUserRepository) *UserRegisterInteractor {
	return &UserRegisterInteractor{
		userRepo: userRepo,
	}
}

func (u *UserRegisterInteractor) Handle(input *ports.UserRegisterInputData) (*ports.UserRegisterOutputData, error) {
	// ユーザ生成はファクトリで行った方がいいけど、今回は省略
	id := uuid.New().String()
	userID, _ := domain.NewUserID(id)
	user, err := domain.NewUser(userID, input.Name(), domain.NomalUser)
	if err != nil {
		return nil, err
	}

	err = u.userRepo.Save(user)
	if err != nil {
		return nil, err
	}

	return ports.NewUserRegisterOutputData(user.ID().Value()), nil
}
