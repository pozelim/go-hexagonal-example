package database

import (
	"errors"
	"github.com/pozelim/go-hexagonal-example/internal/app/model"
)

type EstateRepositoryAdapter struct {
	repository map[model.EstateID]model.Estate
}

func NewEstateRepositoryAdapter() *EstateRepositoryAdapter {
	return &EstateRepositoryAdapter{
		repository: make(map[model.EstateID]model.Estate),
	}
}

func (e *EstateRepositoryAdapter) Save(estate model.Estate) error {
	if _, ok := e.repository[estate.ID]; ok {
		return errors.New("estate with same ID already exists")
	}
	e.repository[estate.ID] = estate
	return nil
}

func (e EstateRepositoryAdapter) Delete(id model.EstateID) error {
	delete(e.repository, id)
	return nil
}
