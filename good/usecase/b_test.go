package usecase_test

import (
	"accept-interfaces-return-structs/good/repository"
	"accept-interfaces-return-structs/good/usecase"
	"testing"
)

type BMock struct{}

func NewBMock() *BMock {
	return &BMock{}
}

// Unnecessary
// func (m *BMock) Save(item any) error {
// 	return nil
// }

func (m *BMock) Find(item any) (any, error) {
	return nil, nil
}

func TestB_Execute(t *testing.T) {
	t.Parallel()

	type fields struct {
		repository repository.B
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
				t.Errorf("A.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
