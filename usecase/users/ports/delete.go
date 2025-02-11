package ports

// ユーザ削除のInputPortインターフェース
type IUserDeleteInputPort interface {
	Handle(input *UserDeleteInputData) (*UserDeleteOutputData, error)
}

// InputPortのIn Outのデータ構造を定義
type UserDeleteInputData struct {
	id string
}

func NewUserDeleteInputData(id string) *UserDeleteInputData {
	return &UserDeleteInputData{
		id: id,
	}
}

func (u *UserDeleteInputData) ID() string {
	return u.id
}

type UserDeleteOutputData struct {
	id string
}

func NewUserDeleteOutputData(id string) *UserDeleteOutputData {
	return &UserDeleteOutputData{
		id: id,
	}
}

func (u *UserDeleteOutputData) ID() string {
	return u.id
}
