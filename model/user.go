package model

import (
	"fmt"
	"github.com/JeffMangan/go-ddd-cart/shared"
	"github.com/asaskevich/govalidator"
	"regexp"
	"time"
)
//https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/dynamodbattribute/
//User stores info for authentication and other basic User info
type User struct {
	base
	Emailx       string `json:"email,omitempty"`
	FirstNamex   string `json:"first_name,omitempty"`
	LastNamex    string `json:"last_name,omitempty"`
	DisplayNamex string `json:"display_name,omitempty"`
}


var (
	firstNameMaxLength int = 50
	lastNameMaxLength = 50
	displayNameSize = map[string]int{
		"min": 5,
		"max" : 15,
	}
)

//SetUpdateDate sets the updated data to the current utc time //////////////////////////////////
func (u *User) SetUpdateDate() {
	u.DateUpdatedUTCx = time.Now().UTC()
}

//Email gets the users Emailx address //////////////////////////////////
func (u *User) Email() string { return u.Emailx }
//SetEmail sets the users Emailx address
func (u *User) SetEmail(email string) *shared.CustomError {
	if govalidator.IsEmail(email) {
		u.Emailx = email
		return nil
	} else if email == "" {
		return  shared.NewCustomError("Email is required", shared.ErrorTypeNilArgument)
	}
	return shared.NewCustomError("Email is not valid", shared.ErrorTypeInvalidArgument)
}

//FirstName gets the users first name //////////////////////////////////
func (u *User) FirstName() string { return u.FirstNamex }
//SetFirstName sets the users first name
func (u *User) SetFirstName(firstName string) *shared.CustomError {
	if firstName == "" {
		return shared.NewCustomError("First name is required.", shared.ErrorTypeInvalidUserArgument)
	} else if len(firstName) > firstNameMaxLength {
		return shared.NewCustomError(fmt.Sprintf("First name max length is %d characters", firstNameMaxLength), shared.ErrorTypeInvalidUserArgument)
	} else if u.FirstNamex == firstName {
		return shared.NewCustomError("First name has not changed.", shared.ErrorTypeInvalidUserArgument)
	}
	u.FirstNamex = firstName
	return nil
}

//LastName gets the users last name //////////////////////////////////
func (u *User) LastName() string { return u.LastNamex }
//SetLastName sets the users last name
func (u *User) SetLastName(lastName string) *shared.CustomError {
	if lastName == "" {
		return shared.NewCustomError("Last name is required.", shared.ErrorTypeInvalidUserArgument)
	} else if len(lastName) > lastNameMaxLength {
		return shared.NewCustomError(fmt.Sprintf("Last name max length is %d characters", lastNameMaxLength), shared.ErrorTypeInvalidUserArgument)
	} else if u.LastNamex == lastName {
		return shared.NewCustomError("Last name has not changed.", shared.ErrorTypeInvalidUserArgument)
	}
	u.LastNamex = lastName
	return nil
}

//DisplayName gets the users Display name //////////////////////////////////
func (u *User) DisplayName() string { return u.DisplayNamex }
//SetDisplayName sets the users Display name
func (u *User) SetDisplayName(displayName string) *shared.CustomError {

	if displayName == "" {
		return shared.NewCustomError("Display name is required.", shared.ErrorTypeInvalidUserArgument)
	} else if len(displayName) < displayNameSize["min"] || len(displayName) > displayNameSize["max"] {
		return shared.NewCustomError(fmt.Sprintf("Display name must be bewtween %d and %d characters in length", displayNameSize["min"], displayNameSize["max"]), shared.ErrorTypeInvalidUserArgument)
	} else if u.DisplayNamex == displayName {
		return shared.NewCustomError("Display name has not changed.", shared.ErrorTypeInvalidUserArgument)
	} else {
		Re := regexp.MustCompile("^[a-zA-Z0-9_-]*$")
		if !Re.MatchString(displayName) {
			return shared.NewCustomError("Display name may only contain alpha numeric characters",shared.ErrorTypeInvalidArgument)
		}
	}
	u.DisplayNamex = displayName
	return nil
}

// NewUser returns a new User entity
func NewUser(user map[string]interface{}) (*User, *shared.CustomError) {
	//todo: is this janky code can't do a type cast if it's nil so ?
	var err *shared.CustomError
	u := &User{}
	u.base = *newBase()
	if err = u.SetFirstName(checkNil(user["first_name"]).(string)); err != nil {
		return nil, err
	}
	if err = u.SetLastName(checkNil(user["last_name"]).(string)); err != nil {
		return nil, err
	}
	if err = u.SetDisplayName(checkNil(user["display_name"]).(string)); err != nil {
		return nil, err
	}
	if err = u.SetEmail(checkNil(user["email"]).(string)); err != nil {
		return nil, err
	}
	return u, nil
}