package usecase_test

import (
	"accept-interfaces-return-structs/bad/repository"
	"accept-interfaces-return-structs/bad/usecase"
	"testing"
)

type AMock struct{}

func NewAMock() repository.Repository {
	return &AMock{}
}

func (m *AMock) Save(item any) error {
	return nil
}

func (m *AMock) Find(item any) (any, error) {
	return nil, nil
}

func TestA_Execute(t *testing.T) {
	t.Parallel()

	type fields struct {
		repository repository.Repository
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				repository: NewAMock(),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := usecase.NewA(tt.fields.repository)
			if err := a.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("A.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
