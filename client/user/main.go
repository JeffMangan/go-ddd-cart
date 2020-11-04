package main

import (
	"encoding/json"
	"github.com/JeffMangan/go-ddd-cart/application"
	"github.com/JeffMangan/go-ddd-cart/infrastructure/repository"
	"github.com/JeffMangan/go-ddd-cart/shared"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler is executed by AWS Lambda in the main function. Once the request
// is processed, it returns an Amazon API Gateway response object to AWS Lambda
// this is the aggregate for DynamoDBUserRepo entity
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var config *shared.Config
	var err *shared.CustomError
	var logger *shared.Logger
	var repo *repository.DynamoDBUserRepo
	var user *application.User
	//todo:determine if this is the best type to use
	var response map[string]interface{}

	//todo:fix the circular reference here
	if config, err = shared.NewConfigFromPath(shared.ConfigPath); err != nil {
//		logger.AddFieldToLog(shared.FieldTypeFunctionName, "Handler")
//		logger.AddFieldToLog(shared.FieldTypeStructName, "shared.Config")
//		logger.AddFieldToLog(shared.FieldTypeFunctionCalled, "shared.NewConfigFromPath")
//		logger.AddFieldToLog(shared.FieldTypeSystemError, err)
//		logger.LogError(shared.LogTypeError)
		//panic(err.Error())
		shared.StandardErrorLog(err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers: map[string]string{
				"Content-Type": "text/html",
			},
		}, nil
	}

	if logger, err = shared.NewLogger(config.LogLevel, config.LogFormatType); err != nil {
		//logger.AddFieldToLog(shared.FieldTypeFunctionName, "Handler")
		//logger.AddFieldToLog(shared.FieldTypeStructName, "shared.Logger")
		//logger.AddFieldToLog(shared.FieldTypeFunctionCalled, "shared.NewLogger")
		//logger.AddFieldToLog(shared.FieldTypeSystemError, err)
		//logger.LogError(shared.LogTypeError)
		shared.StandardErrorLog(err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers: map[string]string{
				"Content-Type": "text/html",
			},
		}, nil
	}

	if repo, err = repository.NewDynamoDBUserRepository(); err != nil {
		logger.AddFieldToLog(shared.FieldTypeFunctionName, "Handler")
		logger.AddFieldToLog(shared.FieldTypeStructName, "repository.DynamoDBUserRepository")
		logger.AddFieldToLog(shared.FieldTypeFunctionCalled, "repository.NewDynamoDBUserRepository")
		logger.AddFieldToLog(shared.FieldTypeSystemError, err)
		logger.LogError(shared.LogTypeError)

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers: map[string]string{
				"Content-Type": "text/html",
			},
		}, nil
	}

	body := make(map[string]interface{})
	if er := json.Unmarshal([]byte(request.Body), &body); er != nil {
		logger.AddFieldToLog(shared.FieldTypeFunctionName, "Handler")
		logger.AddFieldToLog(shared.FieldTypeStructName, "map[string]interface{}")
		logger.AddFieldToLog(shared.FieldTypeFunctionCalled, "json.Unmarshal")
		logger.AddFieldToLog(shared.FieldTypeUserInput, request.Body)
		logger.AddFieldToLog(shared.FieldTypeSystemError, er.Error())
		logger.LogError(shared.LogTypeError)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers: map[string]string{
				"Content-Type": "text/html",
			},
		}, nil
	}

	if user, err = application.NewUserApplication(repo, logger, config); err != nil {
		logger.AddFieldToLog(shared.FieldTypeFunctionName, "Handler")
		logger.AddFieldToLog(shared.FieldTypeStructName, "application.UserApplication")
		logger.AddFieldToLog(shared.FieldTypeFunctionCalled, "application.NewUserApplication")
		logger.AddFieldToLog(shared.FieldTypeSystemError, err.Error())
		logger.LogError(shared.LogTypeError)

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers: map[string]string{
				"Content-Type": "text/html",
			},
		}, nil
	}
	u := make(map[string]interface{})
	if e := json.Unmarshal([]byte(request.Body), &u); e != nil {
		logger.AddFieldToLog(shared.FieldTypeFunctionName, "Handler")
		logger.AddFieldToLog(shared.FieldTypeFunctionCalled, "json.Unmarshall")
		logger.AddFieldToLog(shared.FieldTypeStructName, "map[string]interface{}")
		logger.AddFieldToLog(shared.FieldTypeErrorInfo, "unable to parse request user params: " + request.Body)
		logger.AddFieldToLog(shared.FieldTypeSystemError, e.Error())
		logger.LogError(shared.LogTypeError)

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers: map[string]string{
				"Content-Type": "text/html",
			},
		}, nil
	}

	if response, err = user.Register(u); err != nil {
		logger.AddFieldToLog(shared.FieldTypeFunctionName, "Handler")
		logger.AddFieldToLog(shared.FieldTypeFunctionCalled, "user.Register")
		logger.AddFieldToLog(shared.FieldTypeStructName, "map[string]interface{}")
		logger.AddFieldToLog(shared.FieldTypeSystemError, err)
		logger.LogError(shared.LogTypeError)

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers: map[string]string{
				"Content-Type": "text/html",
			},
		}, nil
	}

	id := (response["id"]).(string)
	r := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       id,
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
	}

	logger.AddFieldToLog(shared.FieldTypeUserInput, request)
	logger.AddFieldToLog(shared.FieldTypeFunctionName, "Handler")
	logger.AddFieldToLog(shared.FieldTypeStructName, "events.APIGatewayProxyResponse")
	logger.AddFieldToLog(shared.FieldTypeOutput, r)
	logger.LogInfo(shared.LogTypeInfo)

	return r, nil
}

func main() {
	lambda.Start(Handler)
}
