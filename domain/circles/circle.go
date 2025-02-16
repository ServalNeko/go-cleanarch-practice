package circles

import (
	"errors"
	"go-arch-practice/domain/lib"
	"go-arch-practice/domain/users"
)

type CircleID struct {
	lib.ValueObject[string]
}

func NewCircleID(id string) (*CircleID, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}

	return &CircleID{
		lib.NewValueObject(id),
	}, nil
}

type Circle struct {
	id      *CircleID
	name    string
	owner   *users.UserID
	members []users.UserID
}

func NewCircle(id *CircleID, name string, owner *users.UserID, members []users.UserID) (*Circle, error) {

	if id == nil || name == "" {
		return nil, errors.New("id or name is empty")
	}

	return &Circle{
		id:      id,
		name:    name,
		owner:   owner,
		members: members,
	}, nil
}

func (c *Circle) ID() *CircleID {
	return c.id
}

func (c *Circle) Name() string {
	return c.name
}

func (c *Circle) Owner() *users.UserID {
	return c.owner
}

func (c *Circle) Members() []users.UserID {

	members := make([]users.UserID, 0, c.CountMembers(false))
	members = append(members, c.members...)
	return members
}

func (c *Circle) MemberIDs() []string {

	memberIds := make([]string, 0, c.CountMembers(false))
	for _, m := range c.members {
		memberIds = append(memberIds, m.Value())
	}

	return memberIds
}

func (c *Circle) CountMembers(containOwner bool) int {
	if containOwner && c.owner != nil {
		return len(c.members) + 1
	}
	return len(c.members)
}

func (c *Circle) Join(new_member users.User, fullSpecification ICircleFullSpecification) error {

	for _, m := range c.members {
		if m.Equals(new_member.ID()) {
			return errors.New("既に所属済みのメンバーです")
		}
	}

	if fullSpecification.IsSatisfiedBy(c) {
		return errors.New("サークルに所属しているメンバーが上限に達しています。")
	}

	c.members = append(c.members, new_member.ID())

	return nil
}

func (c *Circle) ChangeName(name string) {
	c.name = name
}

func (c *Circle) existOwner() bool {
	return c.owner != nil
}
