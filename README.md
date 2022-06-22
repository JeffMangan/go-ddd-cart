## Table of contents
* [General info](#general-info)
* [Technologies](#technologies)
* [Project Structure](#project-structure)
* [Setup](#setup)

## General info
This project needs additional contributers to buld Aggregate Roots following a shopping cart domain. There  is also a lot of code that  could be improved, some of it is  marked with //ToDo.
* This is a starter project of a Domain Driven Design implimentation using Go. The example emulates a typical shopping cart project, however it is in the very beginnings and as of now only has the user aggregate root started.  
* Domain-driven design (DDD) is the concept that the structure and language of your code (class names, class methods, class variables) should match the business domain. For example, if your software processes loan applications, it might have classes such as LoanApplication and Customer, and methods such as AcceptOffer and Withdraw.
	
## Technologies
Go is the programming language and it primarly targeted to use AWS Serverless services including Lambda, API Gateway, Cognito, IAM, DynamoDB, Aurora, and more.
	
## Project Structure
The building blocks of this project include:
* application -- Defines the use cases the software is supposed to do and coordinates the domain objects to work out problems.  This layer is kept thin. It does not contain business rules or knowledge, but only coordinates tasks and delegates work to collaborations of domain objects in the next layer down.  It does not have state reflecting the business situation, but it can have state that reflects the progress of a task for the user or the program.
* client -- This is the top most layer, essentailly for example a rest service wrapper.  It is meant to wire up all of the depedencies for the application layer.  Handler is executed by AWS Lambda in the main function. Once the request is processed, it returns an Amazon API Gateway response object to AWS Lambda this is the aggregate for DynamoDBUserRepo entity.
* db --  This is where the data store is implimented. It uses DynamoDB Local and is currently used for the autommated testing.
* domain -- Responsible for representing concepts of the business, information about the business situation, and business rules.  State that reflects the business situation is controlled and used here, even though the technical details of storing it are delegated to the infrastructure.  This layer is the heart of business software.
* infrastructure -- Implimention of tech stack.
* model -- Domain Entities
* shared -- Shared logic


## Setup
For the tests, the current  infra repo uses dynamodb local, it should use a mock (change soon).
1. Download and extract this file: https://s3.us-west-2.amazonaws.com/dynamodb-local/dynamodb_local_latest.zip
1. Copy and past only the DynamoDBLocal_lib & third_party_licenses folders & the DynamoDBLocal.jar file you just extracted and paste that folder into the db/dynamodb_local_latest folder.  This will keep from overwriting the needed project files.
1. Go to the db/dynamodb_local_lastest folder and run the following java command to start the in memory db (install java if you have not)
1. java -Djava.library.path=./DynamoDBLocal_lib -jar DynamoDBLocal.jar -sharedDb project
1.  You should now be able to run all tests from the project root using "go test ./..."

