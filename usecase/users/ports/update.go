package ports

// ユーザ更新のInputPortインターフェース
type IUserUpdateInputPort interface {
	Handle(input *UserUpdateInputData) (*UserUpdateOutputData, error)
}

// InputPortのIn Outのデータ構造を定義
type UserUpdateInputData struct {
	id   string
	name string
}

func NewUserUpdateInputData(id, name string) *UserUpdateInputData {
	return &UserUpdateInputData{
		id:   id,
		name: name,
	}
}

func (u *UserUpdateInputData) ID() string {
	return u.id
}

func (u *UserUpdateInputData) Name() string {
	return u.name
}

type UserUpdateOutputData struct {
}

func NewUserUpdateOutputData() *UserUpdateOutputData {
	return &UserUpdateOutputData{}
}
