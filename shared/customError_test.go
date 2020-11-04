package shared
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ExampleError simple example to show the NewCustomError constructor
func Test_Load_Custom_Error_Works_OK(t *testing.T) {
	const message = "Example string for error"
	//const where = "this is where it happened"
	err := NewCustomError(message, /*where,*/ ErrorTypeInvalidArgument)
	assert.NotNil(t, err, "Expected empty error")
	assert.Equal(t, ErrorTypeInvalidArgument, err.ErrorType(), "Expected an error type of ErrorTypeInvalidArgument")
	assert.Equal(t, message, err.Error(), "Expected error message to be" + message)
	//assert.Equal(t, where, err.Where())
	assert.NotNil(t, err.When(), "Expected when to have a valid date")
}
