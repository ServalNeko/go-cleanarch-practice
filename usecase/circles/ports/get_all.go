package ports

type ICircleGetAllInputPort interface {
	Handle() (*CircleGetAllOutputData, error)
}

type CircleGetAllOutputData struct {
	circles []CircleData
}

func NewCircleGetAllOutputData(circles []CircleData) *CircleGetAllOutputData {
	return &CircleGetAllOutputData{
		circles: circles,
	}
}

func (c *CircleGetAllOutputData) Circles() []CircleData {
	return c.circles
}
