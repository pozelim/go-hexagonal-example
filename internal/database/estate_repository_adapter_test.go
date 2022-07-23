package database

import (
	"github.com/pozelim/go-hexagonal-example/internal/app/model"
	"testing"
)

func TestEstateAdapter_Create(t *testing.T) {
	type args struct {
		estate model.Estate
	}

	arg := args{
		estate: model.Estate{
			ID:    model.EstateID("2BkzYyKOMvdUCnBTAkQ57EFWzS7"),
			Name:  "Dummy Place",
			Point: model.Point{},
			Price: model.Price{},
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
