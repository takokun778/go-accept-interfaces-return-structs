package usecase_test

import (
	"accept-interfaces-return-structs/good/repository"
	"accept-interfaces-return-structs/good/usecase"
	"testing"
)

type AMock struct{}

func NewAMock() *AMock {
	return &AMock{}
}

func (m *AMock) Save(item any) error {
	return nil
}

// Unnecessary
// func (m *Mock) Find(item any) (any, error) {
// 	return nil, nil
// }

func TestA_Execute(t *testing.T) {
	t.Parallel()

	type fields struct {
		repository repository.A
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
