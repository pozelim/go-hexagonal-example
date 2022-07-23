package model

type Estate struct {
	ID    EstateID
	Name  string
	Point Point
	Price Price
	Area  int
}

type EstateID string
