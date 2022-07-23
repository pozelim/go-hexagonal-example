package database

import (
	"errors"
)

type EstateRepositoryAdapter struct {
	repository map[domain.EstateID]domain.Estate
}

func NewEstateRepositoryAdapter() *EstateRepositoryAdapter {
	return &EstateRepositoryAdapter{
		repository: make(map[domain.EstateID]domain.Estate),
	}
}

func (e *EstateRepositoryAdapter) Save(estate domain.Estate) error {
	if _, ok := e.repository[estate.ID]; ok {
		return errors.New("estate with same ID already exists")
	}
	e.repository[estate.ID] = estate
	return nil
}

func (e EstateRepositoryAdapter) Delete(id domain.EstateID) error {
	delete(e.repository, id)
	return nil
}
