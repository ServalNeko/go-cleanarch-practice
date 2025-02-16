package ports

type ICircleDeleteInputPort interface {
	Handle(input *CircleDeleteInputData) (*CircleDeleteOutputData, error)
}

type CircleDeleteInputData struct {
	circleId string
}

func NewCircleDeleteInputData(circleId string) *CircleDeleteInputData {
	return &CircleDeleteInputData{
		circleId: circleId,
	}
}

func (c *CircleDeleteInputData) CircleID() string {
	return c.circleId
}

type CircleDeleteOutputData struct {
}

func NewCircleDeleteOutputData() *CircleDeleteOutputData {
	return &CircleDeleteOutputData{}
}
