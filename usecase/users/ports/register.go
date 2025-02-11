package ports

// ユーザ登録のInputPortインターフェース
type IUserRegisterInputPort interface {
	Handle(input *UserRegisterInputData) (*UserRegisterOutputData, error)
}

// InputPortのIn Outのデータ構造を定義
type UserRegisterInputData struct {
	name string
}

func NewUserRegisterInputData(name string) *UserRegisterInputData {
	return &UserRegisterInputData{
		name: name,
	}
}

func (u *UserRegisterInputData) Name() string {
	return u.name
}

type UserRegisterOutputData struct {
	createdUserID string
}

func NewUserRegisterOutputData(createdUserID string) *UserRegisterOutputData {
	return &UserRegisterOutputData{
		createdUserID: createdUserID,
	}
}

func (u *UserRegisterOutputData) CreatedUserID() string {
	return u.createdUserID
}
