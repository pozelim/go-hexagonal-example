package app

type Estate struct {
	ID    EstateID
	Name  string
	Point Point
	Price Price
	Area  int
}

type EstateID string

type EstatePort interface {
	Create(Estate) (Estate, error)
	Remove(EstateID) error
}

type EstateStorage interface {
	Save(Estate) error
	Delete(EstateID) error
}
