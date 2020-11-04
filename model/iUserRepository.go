package model

import "github.com/JeffMangan/go-ddd-cart/shared"

//IUserRepository is the common interface for all User repo actions
type IUserRepository interface {
	Find(id string) (*User, *shared.CustomError)
	Create(user *User) *shared.CustomError
	Update(user *User) *shared.CustomError
	Delete(id string) *shared.CustomError
}


