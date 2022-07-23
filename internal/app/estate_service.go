package app

import (
	"github.com/segmentio/ksuid"
)

type EstatePort interface {
	Create(Estate) (Estate, error)
	Remove(id EstateID) error
}

type EstateRepositoryPort interface {
	Save(Estate) error
	Delete(EstateID) error
}

type EstateService struct {
	estateRepository EstateRepositoryPort
}

func (e *EstateService) Create(estate Estate) (Estate, error) {
	estate.ID = EstateID(ksuid.New().String())
	if err := e.estateRepository.Save(estate); err != nil {
		return Estate{}, err
	}
	return estate, nil
}

func (e *EstateService) Remove(id EstateID) error {
	//TODO implement me
	panic("implement me")
}

func NewEstateService(estateRepository EstateRepositoryPort) *EstateService {
	return &EstateService{estateRepository: estateRepository}
}
