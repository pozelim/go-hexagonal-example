package service

import (
	"errors"
	"github.com/pozelim/go-hexagonal-example/internal/app"
	"github.com/pozelim/go-hexagonal-example/internal/app/mock"
	"testing"
)

func TestEstateService_Create(t *testing.T) {
	type fields struct {
		estateRepository app.EstateStorage
	}
	type args struct {
		estate app.Estate
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "Should create a new estate",
			fields: fields{&mock.EstateRepository{}},
			args: args{
				app.Estate{},
			},
			wantErr: false,
		},
		{
			name: "Should return error if repository return error",
			fields: fields{&mock.EstateRepository{
				SaveFn: func(estate app.Estate) error {
					return errors.New("mock error")
				},
			}},
			args: args{
				app.Estate{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Estate{
				estateStorage: tt.fields.estateRepository,
			}
			got, err := e.Create(tt.args.estate)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.ID == "" && !tt.wantErr {
				t.Errorf("Create() got = %v", got)
			}
		})
	}
}
