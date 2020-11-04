package shared

import (
	"fmt"
	"github.com/satori/go.uuid"
)
//NewUUID returns a GUID using version v4 from https://tools.ietf.org/html/rfc4122
func NewUUID() string {
	u := uuid.Must(uuid.NewV4(), nil)
	return u.String()
}

func ParseUUID(id string) *CustomError {
	_, err := uuid.FromString(id)
	if err != nil {
		//fmt.Printf("Something went wrong: %s", err)
		return NewCustomError(fmt.Sprintf("%s %s",err.Error(), id), ErrorTypeInvalidArgument)
	}
	return nil
}
