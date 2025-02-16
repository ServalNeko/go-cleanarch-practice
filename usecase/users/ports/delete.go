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

type UserDeleteOutputData struct{}

func NewUserDeleteOutputData() *UserDeleteOutputData {
	return &UserDeleteOutputData{}
}
