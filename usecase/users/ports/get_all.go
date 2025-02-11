package ports

// ユーザ全取得のInputPortインターフェース
type IUserGetAllInputPort interface {
	Handle() (*UserGetAllOutputData, error)
}

type UserGetAllOutputData struct {
	users []UserData
}

func NewUserGetAllOutputData(users []UserData) *UserGetAllOutputData {
	return &UserGetAllOutputData{
		users: users,
	}
}

func (u *UserGetAllOutputData) Users() []UserData {
	return u.users
}
