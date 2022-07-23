package app

import (
	"github.com/pozelim/go-hexagonal-example/internal/app/model"
)

type EstateRepositoryPort interface {
	Save(model.Estate) error
	Delete(model.EstateID) error
}
