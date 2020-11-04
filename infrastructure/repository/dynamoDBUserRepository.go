package repository

import (
	"github.com/JeffMangan/go-ddd-cart/model"
	"github.com/JeffMangan/go-ddd-cart/shared"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type DynamoDBUserRepo struct{

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
	cfg.Region = aws.String("eu-west-2")
	cfg.Endpoint = aws.String("http://localhost:8000")
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
func NewDynamoDBUserRepository() (*DynamoDBUserRepo, *shared.CustomError) {
	return new(DynamoDBUserRepo), nil
	//todo:add logic here
}