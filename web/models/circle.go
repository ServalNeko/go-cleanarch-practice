package models

// レスポンスモデル
type CircleResponseModel struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	MemberIds []string `json:"member_ids"`
}

type CircleGetResponseModel struct {
	Circle CircleResponseModel `json:"circle"`
	Owner  UserResponseModel   `json:"owner"`
}

type CircleIndexResponseModel struct {
	Circles []CircleResponseModel `json:"circles"`
}

type CircleJoinResponseModel struct {
	AddUserId string `json:"add_user_id"`
}

type CirclePostResponseModel struct {
	CreatedId string `json:"created_id"`
}

// リクエストモデル
type CirclePostRequestModel struct {
	CircleName string `json:"circle_name"`
	OwnerID    string `json:"owner_id"`
}

type CirclePutRequestModel struct {
	CircleName string `json:"circle_name"`
}
