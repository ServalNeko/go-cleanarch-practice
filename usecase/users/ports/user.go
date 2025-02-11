package ports

type UserData struct {
	id   string
	name string
}

func NewUserData(id, name string) *UserData {
	return &UserData{
		id:   id,
		name: name,
	}
}

func (u *UserData) ID() string {
	return u.id
}

func (u *UserData) Name() string {
	return u.name
}
