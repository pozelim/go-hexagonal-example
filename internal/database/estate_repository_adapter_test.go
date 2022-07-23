package database

import (
	"github.com/pozelim/go-hexagonal-example/internal/app"
	"testing"
)

func TestEstateAdapter_Create(t *testing.T) {
	type args struct {
		estate app.Estate
	}

	arg := args{
		estate: app.Estate{
			ID:    app.EstateID("2BkzYyKOMvdUCnBTAkQ57EFWzS7"),
			Name:  "Dummy Place",
			Point: app.Point{},
			Price: app.Price{},
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
