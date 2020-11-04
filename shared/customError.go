package shared

import (
	"time"
)

//type (
//	ErrorInterface interface {
//		CustomError() string
//	}
//)

//ErrorType is the type of error
type ErrorType string

const (
	//ErrorTypeInvalidUserArgument represents invalid input from a user
	ErrorTypeInvalidUserArgument ErrorType = "InvalidUserInput"
	//ErrorTypeInvalidArgument represents an invalid argument
	ErrorTypeInvalidArgument ErrorType = "InvalidArgument"
	ErrorTypeNotUniqueArgument ErrorType = "NotUniqueArgument"
	//ErrorTypeNilArgument represents null argument
	ErrorTypeNilArgument ErrorType = "NilArgument"
	//ErrorTypeInvalidResponse represents an invalid response
	ErrorTypeInvalidResponse ErrorType = "InvalidResponse"
	//ErrorTypeSystem represents a system error
	ErrorTypeSystem ErrorType = "SystemError"
)

// CustomError is an error implementation that includes a time and message.
type CustomError struct {
	when      time.Time
	what      string
	//where 	string
	errorType ErrorType
}

// CustomError returns the error  message.
func (e *CustomError) Error() string {
	//fmt.Println(e.what)
	return e.what
}

//When returns the time the error occured.
func (e *CustomError) When() time.Time {
	return e.when
}

//When returns the time when the error occured.
/*
func (e *CustomError) Where() string {
	return e.where
}
*/

//ErrorType returns the ErrorType of this error
func (e *CustomError) ErrorType() ErrorType {
	return e.errorType
}

//NewCustomError returns a shared error message with a time and message
func NewCustomError(what string, /*where string,*/ errorType ErrorType) *CustomError {
	if what == "" {
		panic("Missing error message")
	}
	//fmt.Println(what,errorType)
	return &CustomError{
		time.Now().UTC(),
		what,
		//where,
		errorType,
	}
}
