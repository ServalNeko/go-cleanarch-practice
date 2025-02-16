package ports

type ICircleGetInputPort interface {
	Handle(input *CircleGetInputData) (*CircleGetOutputData, error)
}

type CircleGetInputData struct {
	id string
}

func NewCircleGetInputData(id string) *CircleGetInputData {
	return &CircleGetInputData{
		id: id,
	}
}

func (c *CircleGetInputData) ID() string {
	return c.id
}

type CircleGetOutputData struct {
	circle *CircleData
}

func NewCircleGetOutputData(circle *CircleData) *CircleGetOutputData {
	return &CircleGetOutputData{
		circle: circle,
	}
}

func (c *CircleGetOutputData) Circle() *CircleData {
	return c.circle
}
