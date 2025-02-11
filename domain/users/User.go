package users

import "errors"

type UserType int

const (
	NomalUser UserType = iota
	PremiumUser
)

type User struct {
	id       string
	name     string
	userType UserType
}

func NewUser(id, name string, userType UserType) (*User, error) {

	if id == "" || name == "" {
		return nil, errors.New("id or name is empty")
	}

	return &User{
		id:       id,
		name:     name,
		userType: userType,
	}, nil
}

func (u *User) ID() string {
	return u.id
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
