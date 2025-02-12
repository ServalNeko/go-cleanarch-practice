package users

import domain "go-arch-practice/domain/users"

type IUserRepository interface {
	FindById(id string) (*domain.User, error)
	FindByName(name string) (*domain.User, error)
	FindByIds(ids []string) ([]domain.User, error)
	FindAll() ([]domain.User, error)
	Save(user *domain.User) error
	Update(user *domain.User) error
	Delete(user *domain.User) error
}
