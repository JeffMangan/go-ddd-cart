package repository

import (
	"github.com/JeffMangan/go-ddd-cart/model"
	"github.com/JeffMangan/go-ddd-cart/shared"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type DynamoDBUserRepo struct {
	logger shared.ILogger
	config *shared.Config
}

func (u *DynamoDBUserRepo) Find(id string) (*model.User, *shared.CustomError) {
	panic("implement me")
}

func (u *DynamoDBUserRepo) Create(user *model.User) *shared.CustomError {
	av, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return shared.NewCustomError(err.Error(), shared.ErrorTypeSystem)
	}

	tableName := "User"
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	cfg := aws.Config{}
	cfg.Region = aws.String(u.config.Region)
	cfg.Endpoint = aws.String(u.config.EndPoint)
	sess := session.Must(session.NewSession())
	db := dynamodb.New(sess, &cfg)

	_, err = db.PutItem(input)
	if err != nil {
		return shared.NewCustomError(err.Error(), shared.ErrorTypeSystem)
	}
	return nil
}

func (u *DynamoDBUserRepo) Update(user *model.User) *shared.CustomError {
	panic("implement me")
}

func (u *DynamoDBUserRepo) Delete(id string) *shared.CustomError {
	panic("implement me")
}

//NewDynamoDBUserRepository
func NewDynamoDBUserRepository(l shared.ILogger, c shared.Config) (*DynamoDBUserRepo, *shared.CustomError) {
	ur := new(DynamoDBUserRepo)
	ur.logger = l
	ur.config = &c
	return ur, nil
	//todo:add logic here
}
