package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ExampleError simple example to show the NewError constructor
func Test_Creat_New_ClientResponse(t *testing.T) {

	assert.NotEmpty(t, NewClientResponse())
	assert.Equal(t, NewClientResponse().Headers["Access-Control-Allow-Origin"], "*")
}
