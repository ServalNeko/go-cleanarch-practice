package ports

type ICircleUpdateInputPort interface {
	Handle(input *CircleUpdateInputData) (*CircleUpdateOutputData, error)
}

type CircleUpdateInputData struct {
	id   string
	name string
}

func NewCircleUpdateInputData(id, name string) *CircleUpdateInputData {
	return &CircleUpdateInputData{
		id:   id,
		name: name,
	}
}

func (c *CircleUpdateInputData) ID() string {
	return c.id
}

func (c *CircleUpdateInputData) Name() string {
	return c.name
}

type CircleUpdateOutputData struct {
}

func NewCircleUpdateOutputData() *CircleUpdateOutputData {
	return &CircleUpdateOutputData{}
}
