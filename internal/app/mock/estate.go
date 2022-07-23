package mock

import (
	"github.com/pozelim/go-hexagonal-example/internal/app"
)

type EstateStorage struct {
	SaveFn   func(app.Estate) error
	DeleteFn func(app.EstateID) error
}

func (e *EstateStorage) Save(estate app.Estate) error {
	if e.SaveFn != nil {
		return e.SaveFn(estate)
	}
	return nil
}

func (e *EstateStorage) Delete(id app.EstateID) error {
	if e.DeleteFn != nil {
		return e.DeleteFn(id)
	}
	return nil
}
