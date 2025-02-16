package users

import (
	"context"
	domain "go-arch-practice/domain/users"
	"go-arch-practice/infra/pg/persistence/models"

	"github.com/uptrace/bun"
)

type UserRepository struct {
	db bun.IDB
}

var _ domain.IUserRepository = (*UserRepository)(nil)

func NewUserRepository(db bun.IDB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindById(id string) (*domain.User, error) {
	userDataModel := &models.UserDataModel{}
	err := r.db.NewSelect().Model(userDataModel).Where("id = ?", id).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return r.toDomainModel(*userDataModel), nil
}

func (r *UserRepository) FindByName(name string) (*domain.User, error) {
	userDataModel := &models.UserDataModel{}
	err := r.db.NewSelect().Model(userDataModel).Where("name = ?", name).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return r.toDomainModel(*userDataModel), nil
}

func (r *UserRepository) FindByIds(ids []string) ([]domain.User, error) {
	userDataModels := []models.UserDataModel{}
	err := r.db.NewSelect().Model(&userDataModels).Where("id IN (?)", ids).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	users := make([]domain.User, len(userDataModels))
	for i, userDataModel := range userDataModels {
		users[i] = *r.toDomainModel(userDataModel)
	}
	return users, nil
}

func (r *UserRepository) FindAll() ([]domain.User, error) {
	userDataModels := []models.UserDataModel{}
	err := r.db.NewSelect().Model(&userDataModels).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	users := make([]domain.User, len(userDataModels))
	for i, userDataModel := range userDataModels {
		users[i] = *r.toDomainModel(userDataModel)
	}
	return users, nil
}

func (r *UserRepository) Save(user *domain.User) error {
	userDataModel := &models.UserDataModel{
		Id:       user.ID().Value(),
		Name:     user.Name(),
		UserType: int(user.UserType()),
	}
	_, err := r.db.NewInsert().Model(userDataModel).Exec(context.Background())
	return err
}

func (r *UserRepository) Update(user *domain.User) error {
	userDataModel := &models.UserDataModel{
		Id:       user.ID().Value(),
		Name:     user.Name(),
		UserType: int(user.UserType()),
	}
	_, err := r.db.NewUpdate().Model(userDataModel).WherePK().Exec(context.Background())
	return err
}

func (r *UserRepository) Delete(user *domain.User) error {
	userDataModel := &models.UserDataModel{
		Id:       user.ID().Value(),
		Name:     user.Name(),
		UserType: int(user.UserType()),
	}
	_, err := r.db.NewDelete().Model(userDataModel).WherePK().Exec(context.Background())
	return err
}

func (r *UserRepository) toDomainModel(from models.UserDataModel) *domain.User {
	id, _ := domain.NewUserID(from.Id)
	u, _ := domain.NewUser(id, from.Name, domain.UserType(from.UserType))
	return u
}
