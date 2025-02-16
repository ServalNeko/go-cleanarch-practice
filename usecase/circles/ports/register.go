package ports

type ICircleRegisterInputPort interface {
	Handle(input *CircleRegisterInputData) (*CircleRegisterOutputData, error)
}

type CircleRegisterInputData struct {
	name    string
	ownerId string
}

func NewCircleRegisterInputData(name, ownerId string) *CircleRegisterInputData {
	return &CircleRegisterInputData{
		name:    name,
		ownerId: ownerId,
	}
}

func (c *CircleRegisterInputData) Name() string {
	return c.name
}

func (c *CircleRegisterInputData) OwnerID() string {
	return c.ownerId
}

type CircleRegisterOutputData struct {
	createdCircleId string
}

func NewCircleCreateOutputData(createdCircleId string) *CircleRegisterOutputData {
	return &CircleRegisterOutputData{
		createdCircleId: createdCircleId,
	}
}

func (c *CircleRegisterOutputData) CreatedCircleID() string {
	return c.createdCircleId
}
