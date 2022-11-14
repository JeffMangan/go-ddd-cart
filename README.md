## Table of contents
* [General info](#general-info)
* [Technologies](#technologies)
* [Project Structure](#project-structure)
* [Setup](#setup)

## General info
This is a Go app with some poc quality code in certain areas which will need more attention.  Essentially if you see anything that makes you think "WTF, this is stupid" it is probably one of those areas, given that the goal is to impliment things like SOLID, KISS, and YAGNI.
* This is a starter project of a Domain Driven Design implementation using Go. The example emulates a typical shopping cart project, however it is in the very beginnings and as of now, aside from the infra and shared folders, the rest has only been started for the user.  
* Domain-driven design (DDD) is the concept that the structure and language of your code (class names, class methods, class variables) should match the business domain and behavior. For example, if your software processes loan applications, it might have classes such as LoanApplication and Customer, and methods such as AcceptOffer and Withdraw instead of plain old CRUD statements.
	
## Technologies
* Go
* Docker 
  * DynamoDB Local (for testing), this should not be required, next phase should use mocks.
* Dynamo Local Admin Web Interface (not required) https://www.npmjs.com/package/dynamodb-admin

## Project Structure
The building blocks of this project include:
* client -- This is the top most layer, essentially for example a rest service wrapper.  It is meant to wire up the dependencies for the application layer.  Handler is executed by AWS Lambda in the main function. Once the request is processed, it returns an Amazon API Gateway response object to AWS Lambda this is the aggregate for DynamoDBUserRepo entity.
* application -- Defines the use cases the software is supposed to do and coordinates the domain objects to work out problems.  This layer is kept thin. It does not contain business rules or knowledge, but only coordinates tasks and delegates work to collaborations of domain objects in the next layer down.  It does not have state reflecting the business situation, but it can have state that reflects the progress of a task for the user or the program.
* domain -- Responsible for representing concepts of the business, information about the business situation, and business rules.  State that reflects the business situation is controlled and used here, even though the technical details of storing it are delegated to the infrastructure.  This layer is the heart of business software.
* infrastructure -- Implementation of tech stack.
* model -- Domain Entities
* shared -- Shared logic

## Setup (for DynamoDB Local Testing)
* docker run --name ddblocal -p 8001:8001 amazon/dynamodb-local -jar DynamoDBLocal.jar -sharedDb
  * to understand the cmd line params go here https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DynamoDBLocal.UsageNotes.html#DynamoDBLocal.CommandLineOptions 
  * if you are unable to connect, make sure there are no existing processes running on port 8000.
* This next step is needed each time you run a new container instance, the user table will need to be created.  The endpoint-url is specific to ddb local only.
  * aws dynamodb create-table \
--table-name User \
--attribute-definitions AttributeName=id,AttributeType=S \
--key-schema AttributeName=id,KeyType=HASH \
--provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 \
--endpoint-url http://localhost:8000
  * aws dynamodb list-tables --endpoint-url http://localhost:8000
    * Only required if you want to see that the user table was created.
  * web based ddb local amin console (this is not required but can be helpful to easily see your data)
    * https://www.npmjs.com/package/dynamodb-admin
 * The following environment variables need to be created
   * export AWS_DDBLOCAL_ENDPOINT=http://localhost:8000
   * export AWS_DDBLOCAL_REGION=us-east-1
   * export LOG_LEVEL=Error
   * export LOG_FORMAT_TYPE=json
   * export CONFIG_FILE_PATH=/whever_you_put_it/go-ddd-cart/
     * Update the path to the actual path of your root folder where app.env file is located
     * You may have to configure your ide to use configuration setup for env variables.


## You should now be able to run all tests from the project root using "go test ./..."
* If you get failing tests, make sure your env variables are working, and that you are able to connet to dynamodb inside of docker.