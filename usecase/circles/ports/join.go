package ports

type ICircleJoinInputPort interface {
	Handle(input *CircleJoinInputData) (*CircleJoinOutputData, error)
}

type CircleJoinInputData struct {
	circleId string
	memberId string
}

func NewCircleJoinInputData(circleId, memberId string) *CircleJoinInputData {
	return &CircleJoinInputData{
		circleId: circleId,
		memberId: memberId,
	}
}

func (c *CircleJoinInputData) CircleID() string {
	return c.circleId
}

func (c *CircleJoinInputData) MemberID() string {
	return c.memberId
}

type CircleJoinOutputData struct {
}

func NewCircleJoinOutputData() *CircleJoinOutputData {
	return &CircleJoinOutputData{}
}
