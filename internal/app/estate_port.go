package app

import "github.com/pozelim/go-hexagonal-example/internal/app/model"

type EstatePort interface {
	Create(model.Estate) (model.Estate, error)
	Remove(id model.EstateID) error
}
