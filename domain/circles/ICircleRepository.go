package circles

type ICircleRepository interface {
	FindById(id string) (*Circle, error)
	FindByName(name string) (*Circle, error)
	FindAll() ([]Circle, error)
	Save(circle *Circle) error
	Update(circle *Circle) error
	Delete(circle *Circle) error
}
