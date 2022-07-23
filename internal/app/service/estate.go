package service

import (
	"github.com/pozelim/go-hexagonal-example/internal/app"
	"github.com/segmentio/ksuid"
)

type Estate struct {
	estateStorage app.EstateStorage
}

func NewEstateService(estateStorage app.EstateStorage) *Estate {
	return &Estate{estateStorage: estateStorage}
}

func (e *Estate) Create(estate app.Estate) (app.Estate, error) {
	estate.ID = app.EstateID(ksuid.New().String())
	if err := e.estateStorage.Save(estate); err != nil {
		return app.Estate{}, err
	}
	return estate, nil
}

func (e *Estate) Remove(id app.EstateID) error {
	return e.estateStorage.Delete(id)
}
