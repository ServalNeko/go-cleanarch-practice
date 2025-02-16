package users

type IUserRepository interface {
	FindById(id string) (*User, error)
	FindByName(name string) (*User, error)
	FindByIds(ids []string) ([]User, error)
	FindAll() ([]User, error)
	Save(user *User) error
	Update(user *User) error
	Delete(user *User) error
}
