## Table of contents
* [General info](#general-info)
* [Technologies](#technologies)
* [Project Structure](#project-structure)
* [Setup](#setup)

## General info
This project needs additional contributors to build Aggregate Roots following a shopping cart domain. There  is also a lot of code that  could be improved, some of it is  marked with //ToDo.
* This is a starter project of a Domain Driven Design implementation using Go. The example emulates a typical shopping cart project, however it is in the very beginnings and as of now only has the user aggregate root started.  
* Domain-driven design (DDD) is the concept that the structure and language of your code (class names, class methods, class variables) should match the business domain and behavior. For example, if your software processes loan applications, it might have classes such as LoanApplication and Customer, and methods such as AcceptOffer and Withdraw instead of plain old CRUD statements.
	
## Technologies
This solution is written in Go, and although I am targeting AWS Serverless services, since it uses custom interfaces and inversion of control, it can run on any environment that supports the go runtime by swapping out the underlying infrastructure layer (where the technical implementation is located).
	
## Project Structure
The building blocks of this project include:
* client -- This is the top most layer, essentially for example a rest service wrapper.  It is meant to wire up the dependencies for the application layer.  Handler is executed by AWS Lambda in the main function. Once the request is processed, it returns an Amazon API Gateway response object to AWS Lambda this is the aggregate for DynamoDBUserRepo entity.
* application -- Defines the use cases the software is supposed to do and coordinates the domain objects to work out problems.  This layer is kept thin. It does not contain business rules or knowledge, but only coordinates tasks and delegates work to collaborations of domain objects in the next layer down.  It does not have state reflecting the business situation, but it can have state that reflects the progress of a task for the user or the program.
* domain -- Responsible for representing concepts of the business, information about the business situation, and business rules.  State that reflects the business situation is controlled and used here, even though the technical details of storing it are delegated to the infrastructure.  This layer is the heart of business software.
* infrastructure -- Implementation of tech stack.
* model -- Domain Entities
* shared -- Shared logic
* db --  This is where the local data store is implemented and uses DynamoDB Local for the automated testing.  This should be updated to not depend on a local db, which is possible by simply creating mock data that adheres to the repo interface and can easily be switched out.

## Setup (this needs to be automated, most all of this is for using ddblocal for testing)

-----------------new steps to use docker and not install ddblocal locally----------------
docker run -p 8000:8000 amazon/dynamodb-local

aws dynamodb create-table \
--table-name User \
--attribute-definitions AttributeName=id,AttributeType=S \
--key-schema AttributeName=id,KeyType=HASH \
--provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 \
--endpoint-url http://localhost:8000

aws dynamodb list-tables --endpoint-url http://localhost:8000

web based ddb local amin console
https://www.npmjs.com/package/dynamodb-admin

---------------------------------------------
For the tests, the current  infra repo uses dynamodb local, it should use a mock (change soon).
1. You must have the Java runtime and awscli installed and setup config using "aws configure", although ddb local does not use the keys, you can put whatever you want for local testing.  Here is the link to install the jre https://www.java.com/en/download/manual.jsp and awscli https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html
1. DynamoDB local is used https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DynamoDBLocal.html therefore you must download and extract this file: https://s3.us-west-2.amazonaws.com/dynamodb-local/dynamodb_local_latest.zip
1. Install the nosql workbench to view the structure of the ddb database https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/workbench.html
1. Copy and past only the DynamoDBLocal_lib & third_party_licenses folders & the DynamoDBLocal.jar file you just extracted and paste that folder into the db/dynamodb_local_latest folder.  This will keep from overwriting the needed project files.
1. Go to the db/dynamodb_local_latest folder and run the following java command to start the in memory db (install java if you have not)
1. java -Djava.library.path=./DynamoDBLocal_lib -jar DynamoDBLocal.jar -sharedDb project
1.  In the shared/config.go file change the ConfigPath constant to the path on your local machine
1. You should now be able to run all tests from the project root using "go test ./..."

