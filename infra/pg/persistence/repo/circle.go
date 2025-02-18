package repo

import (
	"context"
	circle_domain "go-arch-practice/domain/circles"
	user_domain "go-arch-practice/domain/users"
	"go-arch-practice/infra/pg/persistence/models"

	"github.com/uptrace/bun"
)

type CircleRepository struct {
	db bun.IDB
}

var _ circle_domain.ICircleRepository = (*CircleRepository)(nil)

func NewCircleRepository(db bun.IDB) *CircleRepository {
	return &CircleRepository{
		db: db,
	}
}

func (r *CircleRepository) FindById(id string) (*circle_domain.Circle, error) {
	circleDataModel := &models.Circle{}
	err := r.db.NewSelect().Model(circleDataModel).Where("id = ?", id).Scan(context.Background())
	if err != nil {
		return nil, err
	}

	return r.toDomainModel(*circleDataModel), nil
}

func (r *CircleRepository) FindByName(name string) (*circle_domain.Circle, error) {
	circleDataModel := &models.Circle{}
	err := r.db.NewSelect().Model(circleDataModel).Where("name = ?", name).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return r.toDomainModel(*circleDataModel), nil
}

func (r *CircleRepository) Save(circle *circle_domain.Circle) error {

	circleDataModel := &models.Circle{
		Id:      circle.ID().Value(),
		Name:    circle.Name(),
		QwnerId: circle.Owner().Value(),
	}
	_, err := r.db.NewInsert().
		Model(circleDataModel).
		On("CONFLICT (id) DO UPDATE").
		Exec(context.Background())

	if err != nil {
		return err
	}

	members := circle.Members()
	for _, member := range members {

		_, err := r.db.NewInsert().
			Model(&models.UserToCircle{
				UserId:   member.Value(),
				CircleId: circle.ID().Value(),
			}).
			On("CONFLICT (user_id, circle_id) DO UPDATE").
			Exec(context.Background())
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *CircleRepository) Delete(circle *circle_domain.Circle) error {

	circleDataModel := &models.Circle{
		Id: circle.ID().Value(),
	}
	_, err := r.db.NewDelete().Model(circleDataModel).WherePK().Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (r *CircleRepository) FindAll() ([]circle_domain.Circle, error) {

	circleDataModels := []models.Circle{}
	err := r.db.NewSelect().Model(&circleDataModels).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	circles := make([]circle_domain.Circle, len(circleDataModels))
	for i, circleDataModel := range circleDataModels {
		circles[i] = *r.toDomainModel(circleDataModel)
	}

	return circles, nil
}

func (r *CircleRepository) toDomainModel(from models.Circle) *circle_domain.Circle {
	circleId, _ := circle_domain.NewCircleID(from.Id)

	var ownerId *user_domain.UserID = nil
	if from.QwnerId != "" {
		ownerId, _ = user_domain.NewUserID(from.QwnerId)
	}

	var members []user_domain.UserID
	if from.Members != nil {
		members = make([]user_domain.UserID, len(from.Members))
		for i, member := range from.Members {
			memberId, _ := user_domain.NewUserID(member.Id)
			members[i] = *memberId
		}
	} else {
		members = make([]user_domain.UserID, 0)
	}

	circle, _ := circle_domain.NewCircle(circleId, from.Name, ownerId, members)

	return circle
}
