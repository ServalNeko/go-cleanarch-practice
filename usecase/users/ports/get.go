package ports

// ユーザ取得のInputPortインターフェース
type IUserGetInputPort interface {
	Handle(input *UserGetInputData) (*UserGetOutputData, error)
}

// InputPortのIn Outのデータ構造を定義
type UserGetInputData struct {
	id string
}

func NewUserGetInputData(id string) *UserGetInputData {
	return &UserGetInputData{
		id: id,
	}
}

func (u *UserGetInputData) ID() string {
	return u.id
}

type UserGetOutputData struct {
	user UserData
}

func NewUserGetOutputData(user UserData) *UserGetOutputData {
	return &UserGetOutputData{
		user: user,
	}
}

func (u *UserGetOutputData) User() UserData {
	return u.user
}
