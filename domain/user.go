package domain

import "github.com/JeffMangan/go-ddd-cart/shared"

/*
Responsible for representing concepts of the business, information about the business situation, and business rules.
State that reflects the business situation is controlled and used here, even though the technical details of storing it
//are delegated to the infrastructure.
This layer is the heart of business software.
*/

type User struct {
}

func (u *User) DoStuff() *shared.CustomError {
	return nil
}

func NewUser() User {
	return User{}
}
