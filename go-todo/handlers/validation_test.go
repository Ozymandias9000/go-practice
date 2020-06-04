package handlers

import (
	"go-todo/domain"
	"testing"
)

func Test_validateRegisterUser(t *testing.T) {
	type args struct {
		payload *domain.RegisterPayload
	}

	args1 := args{payload: &domain.RegisterPayload{Username: "joe", Password: "test123", ConfirmPassword: "test123", Email: "joe@joe.com"}}
	args2 := args{payload: &domain.RegisterPayload{Username: "joe", Password: "test123", ConfirmPassword: "ohnooo", Email: "joe@joe.com"}}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "shouldNotFailIfValid",
			args:    args1,
			wantErr: false,
		},
		{
			name:    "shouldFailIfPasswordAndConfirmNotEqual",
			args:    args2,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateRegisterUser(tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("validateRegisterUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
