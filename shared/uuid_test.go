package shared

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"
)

// ExampleError simple example to show the NewCustomError constructor
func Test_Creat_New_UUIDV4(t *testing.T) {
	assert.NotEmpty(t, NewUUID())
	assert.True(t, govalidator.IsUUIDv4(NewUUID()))
	assert.NotEqual(t, NewUUID(), uuid.Nil)
}

func Test_ParseUUID(t *testing.T) {

	uuid := NewUUID()
	assert.Empty(t, ParseUUID(uuid))
	assert.NotEmpty(t, ParseUUID("xxx"))
}