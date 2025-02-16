package ports

import user_ports "go-arch-practice/usecase/users/ports"

type ICircleCandidatesInputPort interface {
	Candidates(input *CircleCandidatesInputData) (*CircleCandidatesOutputData, error)
}

type CircleCandidatesInputData struct {
	circleId string
	page     int
	size     int
}

func NewCircleCandidatesInputData(circleId string, page, size int) *CircleCandidatesInputData {
	return &CircleCandidatesInputData{
		circleId: circleId,
		page:     page,
		size:     size,
	}
}

func (c *CircleCandidatesInputData) CircleID() string {
	return c.circleId
}

func (c *CircleCandidatesInputData) Page() int {
	return c.page
}

func (c *CircleCandidatesInputData) Size() int {
	return c.size
}

type CircleCandidatesOutputData struct {
	users []user_ports.UserData
}

func NewCircleCandidatesOutputData(users []user_ports.UserData) *CircleCandidatesOutputData {
	return &CircleCandidatesOutputData{
		users: users,
	}
}

func (c *CircleCandidatesOutputData) Users() []user_ports.UserData {
	return c.users
}
