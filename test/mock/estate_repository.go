package mock

type EstateRepository struct {
	SaveFn   func(domain.Estate) error
	DeleteFn func(domain.EstateID) error
}

func (e *EstateRepository) Save(estate domain.Estate) error {
	if e.SaveFn != nil {
		return e.SaveFn(estate)
	}
	return nil
}

func (e *EstateRepository) Delete(id domain.EstateID) error {
	if e.DeleteFn != nil {
		return e.DeleteFn(id)
	}
	return nil
}
