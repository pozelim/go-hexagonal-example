package app

import (
	"github.com/pozelim/go-hexagonal-example/internal/app/domain"
	"github.com/segmentio/ksuid"
)

type EstatePort interface {
	Create(domain.Estate) (domain.Estate, error)
	Remove(id domain.EstateID) error
}

type EstateRepositoryPort interface {
	Save(domain.Estate) error
	Delete(domain.EstateID) error
}

type EstateService struct {
	estateRepository EstateRepositoryPort
}

func (e *EstateService) Create(estate domain.Estate) (domain.Estate, error) {
	estate.ID = domain.EstateID(ksuid.New().String())
	if err := e.estateRepository.Save(estate); err != nil {
		return domain.Estate{}, err
	}
	return estate, nil
}

func (e *EstateService) Remove(id domain.EstateID) error {
	//TODO implement me
	panic("implement me")
}

func NewEstateService(estateRepository EstateRepositoryPort) *EstateService {
	return &EstateService{estateRepository: estateRepository}
}
