package database

import (
	"testing"
)

func TestEstateAdapter_Create(t *testing.T) {
	type args struct {
		estate domain.Estate
	}

	arg := args{
		estate: domain.Estate{
			ID:    domain.EstateID("2BkzYyKOMvdUCnBTAkQ57EFWzS7"),
			Name:  "Dummy Place",
			Point: domain.Point{},
			Price: domain.Price{},
			Area:  0,
		},
	}

	t.Run("Save should return error when creating twice", func(t *testing.T) {
		e := NewEstateRepositoryAdapter()
		e.Save(arg.estate)
		if err := e.Save(arg.estate); err == nil {
			t.Errorf("Save() error = %v, wantErr false", err)
		}
	})

}
