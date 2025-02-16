package circles

import "go-arch-practice/domain/users"

type ICircleFullSpecification interface {
	IsSatisfiedBy(circle *Circle) bool
}

type CircleFullSpecification struct {
	repo users.IUserRepository
}

var _ ICircleFullSpecification = (*CircleFullSpecification)(nil)

func NewCircleFullSpecification(repo users.IUserRepository) *CircleFullSpecification {
	return &CircleFullSpecification{repo: repo}
}

func (s *CircleFullSpecification) IsSatisfiedBy(circle *Circle) bool {
	return circle.CountMembers(true) > 30
}
