package app

import (
	"errors"
	"github.com/pozelim/go-hexagonal-example/test/mock"
	"testing"
)

func TestEstateService_Create(t *testing.T) {
	type fields struct {
		estateRepository EstateRepositoryPort
	}
	type args struct {
		estate Estate
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
				Estate{},
			},
			wantErr: false,
		},
		{
			name: "Should return error if repository return error",
			fields: fields{&mock.EstateRepository{
				SaveFn: func(estate Estate) error {
					return errors.New("mock error")
				},
			}},
			args: args{
				Estate{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EstateService{
				estateRepository: tt.fields.estateRepository,
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
