package domain

import (
	"github.com/JeffMangan/go-ddd-cart/shared"
	"reflect"
	"testing"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		name string
		want User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_DoStuff(t *testing.T) {
	tests := []struct {
		name string
		want *shared.CustomError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{}
			if got := u.DoStuff(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DoStuff() = %v, want %v", got, tt.want)
			}
		})
	}
}
