package application

import (
	"github.com/JeffMangan/go-ddd-cart/model"
	"github.com/JeffMangan/go-ddd-cart/shared"
)

/*
Defines the jobs (use cases) the software is supposed to do and coordinates the domain objects to work out problems.
This layer is kept thin. It does not contain business rules or knowledge, but only coordinates tasks and delegates work
to collaborations of domain objects in the next layer down.
It does not have state reflecting the business situation, but it can have state that reflects the progress of a task
for the user or the program.
*/

//user is the coordinator for all user domain behaviors
type User struct {
	repo   model.IUserRepository
	logger shared.ILogger
	config *shared.Config
}

func (u *User) Register(user map[string]interface{}) (map[string]interface{}, *shared.CustomError) {

	//todo: do stuff with domain if needed
	//domain := domain.NewUser()
	//domain.DoStuff()

	var uu *model.User
	var err *shared.CustomError
	if uu, err = model.NewUser(user); err != nil {
		return nil, err
	}

	if err := u.repo.Create(uu); err != nil {
		return nil, err
	}

	response := make(map[string]interface{})
	response["id"] = uu.ID()

	u.logger.AddFieldToLog(shared.FieldTypeUserInput, user)
	u.logger.AddFieldToLog(shared.FieldTypeFunctionName, "application.User.Register")
	u.logger.AddFieldToLog(shared.FieldTypeStructName, "map[string]interface{}")
	u.logger.AddFieldToLog(shared.FieldTypeOutput, response)
	u.logger.LogInfo(shared.LogTypeInfo)

	return response, nil
}

func NewUserApplication(r model.IUserRepository, l shared.ILogger, c *shared.Config) (*User, *shared.CustomError) {
	return &User{
		r,
		l,
		c,
	}, nil
}
