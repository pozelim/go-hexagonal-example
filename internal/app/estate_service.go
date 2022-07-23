package app

import (
	"github.com/pozelim/go-hexagonal-example/internal/app/model"
	"github.com/segmentio/ksuid"
)

type EstateService struct {
	estateRepository EstateRepositoryPort
}

func (e *EstateService) Create(estate model.Estate) (model.Estate, error) {
	estate.ID = model.EstateID(ksuid.New().String())
	if err := e.estateRepository.Save(estate); err != nil {
		return model.Estate{}, err
	}
	return estate, nil
}

func (e *EstateService) Remove(id model.EstateID) error {
	//TODO implement me
	panic("implement me")
}

func NewEstateService(estateRepository EstateRepositoryPort) *EstateService {
	return &EstateService{estateRepository: estateRepository}
}
