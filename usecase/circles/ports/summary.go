package ports

type ICircleSummaryInputPort interface {
	Handle(input *CircleSummaryInputData) (*CircleSummaryOutputData, error)
}

type CircleSummaryData struct {
	circleId  string
	ownerName string
}

func NewCircleSummaryData(circleId, ownerName string) *CircleSummaryData {
	return &CircleSummaryData{
		circleId:  circleId,
		ownerName: ownerName,
	}
}

func (c *CircleSummaryData) CircleID() string {
	return c.circleId
}

func (c *CircleSummaryData) OwnerName() string {
	return c.ownerName
}

type CircleSummaryInputData struct {
	page int
	size int
}

func NewCircleSummaryInputData(page, size int) *CircleSummaryInputData {
	return &CircleSummaryInputData{
		page: page,
		size: size,
	}
}

func (c *CircleSummaryInputData) Page() int {
	return c.page
}

func (c *CircleSummaryInputData) Size() int {
	return c.size
}

type CircleSummaryOutputData struct {
	summaries []CircleSummaryData
}

func NewCircleSummaryOutputData(summaries []CircleSummaryData) *CircleSummaryOutputData {
	return &CircleSummaryOutputData{
		summaries: summaries,
	}
}

func (c *CircleSummaryOutputData) Summaries() []CircleSummaryData {
	return c.summaries
}
