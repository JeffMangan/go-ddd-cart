package main

import (
	"os"
	"testing"

	"github.com/JeffMangan/go-ddd-cart/shared"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

//var test string = "hi"

func setup()    {}
func shutdown() {}

func TestMain(m *testing.M) {
	setup()
	//test = "no"
	//fmt.Println(test)
	code := m.Run()
	shutdown()
	os.Exit(code)
}

//func TestHandler(t *testing.T) {
func TestSomething(t *testing.T) {
	request := events.APIGatewayProxyRequest{}
	request.Body = "{\"first_name\":\"Jeff\",\"last_name\":\"Mangan\",\"display_name\":\"JeffM\",\"email\":\"Jeff@Magentatech.com\"}"

	expectedResponse := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
		Body: "",
	}
	//fmt.Println("handler test")
	if response, err := Handler(request); err != nil {
		//todo:fix this
		//fmt.Println("handler test failed")
		//fmt.Println(response)
		//fmt.Println(err)
	} else {
		assert.Equal(t, expectedResponse.Headers, response.Headers)
		assert.Nil(t, shared.ParseUUID(response.Body))
		assert.Equal(t, expectedResponse.StatusCode, response.StatusCode)
	}

}
