package models

// レスポンスモデル
type UserResponseModel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UserGetResponseModel struct {
	UserResponseModel
}

type UserIndexResponseModel struct {
	Users []UserResponseModel `json:"users"`
}

type UserPostResponseModel struct {
	CreatedId string `json:"created_id"`
}

// リクエストモデル
type UserPostRequestModel struct {
	Name string `json:"name"`
}

type UserPutRequestModel struct {
	Name string `json:"name"`
}
