package mock

import "github.com/pozelim/go-hexagonal-example/internal/app/model"

type EstateRepository struct {
	SaveFn   func(model.Estate) error
	DeleteFn func(model.EstateID) error
}

func (e *EstateRepository) Save(estate model.Estate) error {
	if e.SaveFn != nil {
		return e.SaveFn(estate)
	}
	return nil
}

func (e *EstateRepository) Delete(id model.EstateID) error {
	if e.DeleteFn != nil {
		return e.DeleteFn(id)
	}
	return nil
}
