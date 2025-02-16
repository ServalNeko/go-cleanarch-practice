package ports

type CircleData struct {
	id        string
	name      string
	ownerId   string
	memberIds []string
}

func NewCircleData(id, name, ownerId string, memberIds []string) *CircleData {
	return &CircleData{
		id:        id,
		name:      name,
		ownerId:   ownerId,
		memberIds: memberIds,
	}
}

func (c *CircleData) ID() string {
	return c.id
}

func (c *CircleData) Name() string {
	return c.name
}

func (c *CircleData) OwnerID() string {
	return c.ownerId
}

func (c *CircleData) MemberIDs() []string {
	return c.memberIds
}
