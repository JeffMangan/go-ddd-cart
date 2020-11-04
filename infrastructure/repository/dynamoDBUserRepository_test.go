package repository

import (
	"github.com/JeffMangan/go-ddd-cart/model"
	"github.com/JeffMangan/go-ddd-cart/shared"
	"reflect"
	"testing"
)

func TestDynamoDBUserRepo_Create(t *testing.T) {
	type args struct {
		user *model.User
	}
	tests := []struct {
		name string
		args args
		want *shared.CustomError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &DynamoDBUserRepo{}
			if got := u.Create(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDynamoDBUserRepo_Delete(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want *shared.CustomError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &DynamoDBUserRepo{}
			if got := u.Delete(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDynamoDBUserRepo_Find(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name  string
		args  args
		want  *model.User
		want1 *shared.CustomError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &DynamoDBUserRepo{}
			got, got1 := u.Find(tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Find() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDynamoDBUserRepo_Update(t *testing.T) {
	type args struct {
		user *model.User
	}
	tests := []struct {
		name string
		args args
		want *shared.CustomError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &DynamoDBUserRepo{}
			if got := u.Update(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDynamoDBUserRepository(t *testing.T) {
	tests := []struct {
		name  string
		want  *DynamoDBUserRepo
		want1 *shared.CustomError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := NewDynamoDBUserRepository()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDynamoDBUserRepository() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NewDynamoDBUserRepository() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}