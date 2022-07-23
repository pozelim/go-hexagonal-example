package database

import (
	"errors"
	"github.com/pozelim/go-hexagonal-example/internal/app"
)

type EstateRepositoryAdapter struct {
	repository map[app.EstateID]app.Estate
}

func NewEstateRepositoryAdapter() *EstateRepositoryAdapter {
	return &EstateRepositoryAdapter{
		repository: make(map[app.EstateID]app.Estate),
	}
}

func (e *EstateRepositoryAdapter) Save(estate app.Estate) error {
	if _, ok := e.repository[estate.ID]; ok {
		return errors.New("estate with same ID already exists")
	}
	e.repository[estate.ID] = estate
	return nil
}

func (e EstateRepositoryAdapter) Delete(id app.EstateID) error {
	delete(e.repository, id)
	return nil
}
