package ports

// ユーザ全取得のInputPortインターフェース
type IUserGetAllInputPort interface {
	Handle(input *UserGetAllInputData) (*UserGetAllOutputData, error)
}

// InputPortのIn Outのデータ構造を定義
type UserGetAllInputData struct {
	users []UserData
}

func NewUserGetAllInputData(users []UserData) *UserGetAllInputData {
	return &UserGetAllInputData{
		users: users,
	}
}

func (u *UserGetAllInputData) Users() []UserData {
	return u.users
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
