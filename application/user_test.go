package application

import (
	"github.com/JeffMangan/go-ddd-cart/model"
	"github.com/JeffMangan/go-ddd-cart/shared"
	"reflect"
	"testing"
)

//http://cs-guy.com/blog/2015/01/test-main/

func TestNewUserApplication(t *testing.T) {
	type args struct {
		r model.IUserRepository
		l shared.ILogger
		c *shared.Config
	}
	tests := []struct {
		name  string
		args  args
		want  *User
		want1 *shared.CustomError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := NewUserApplication(tt.args.r, tt.args.l, tt.args.c)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserApplication() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NewUserApplication() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUser_Register(t *testing.T) {
	type fields struct {
		repo   model.IUserRepository
		logger shared.ILogger
		config *shared.Config
	}
	type args struct {
		user map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]interface{}
		want1  *shared.CustomError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				repo:   tt.fields.repo,
				logger: tt.fields.logger,
				config: tt.fields.config,
			}
			got, got1 := u.Register(tt.args.user)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Register() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Register() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
