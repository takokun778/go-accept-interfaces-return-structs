package usecase_test

import (
	"accept-interfaces-return-structs/bad/repository"
	"accept-interfaces-return-structs/bad/usecase"
	"testing"
)

type BMock struct{}

func NewBMock() repository.Repository {
	return &BMock{}
}

func (m *BMock) Save(item any) error {
	return nil
}

func (m *BMock) Find(item any) (any, error) {
	return nil, nil
}

func TestB_Execute(t *testing.T) {
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
				repository: NewBMock(),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			b := usecase.NewB(tt.fields.repository)
			if err := b.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("B.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
