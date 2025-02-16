package users

import (
	"errors"
	"go-arch-practice/domain/lib"
)

type UserType int

const (
	NomalUser UserType = iota
	PremiumUser
)

type UserID struct {
	lib.ValueObject[string]
}

func NewUserID(id string) (*UserID, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}

	return &UserID{
		lib.NewValueObject(id),
	}, nil
}

type User struct {
	id       *UserID
	name     string
	userType UserType
}

func NewUser(id *UserID, name string, userType UserType) (*User, error) {

	if id == nil || name == "" {
		return nil, errors.New("id or name is empty")
	}

	return &User{
		id:       id,
		name:     name,
		userType: userType,
	}, nil
}

func (u *User) ID() UserID {
	return *u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) UserType() UserType {
	return u.userType
}

func (u *User) ChangeName(name string) {
	u.name = name
}

func (u *User) Upgrade() {
	u.userType = PremiumUser
}

func (u *User) Downgrade() {
	u.userType = NomalUser
}
