package inmemory

import (
	"errors"
	"github.com/pozelim/go-hexagonal-example/internal/app"
)

type EstateStorage struct {
	repository map[app.EstateID]app.Estate
}

func NewEstateStorage() *EstateStorage {
	return &EstateStorage{
		repository: make(map[app.EstateID]app.Estate),
	}
}

func (e *EstateStorage) Save(estate app.Estate) error {
	if _, ok := e.repository[estate.ID]; ok {
		return errors.New("estate with same ID already exists")
	}
	e.repository[estate.ID] = estate
	return nil
}

func (e EstateStorage) Delete(id app.EstateID) error {
	delete(e.repository, id)
	return nil
}
