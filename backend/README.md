# Vugu/Todo/Async Demo - Backend

The logic for the backend comprises of the following:
- a set of lambda functions implemented in Go
- a set of Makefiles for building the lambda functions and creating a zip file
  which can be uploaded to AWS
- a Serverless Application Model definition which can be used to deploy the
  lambda functions and configure the API Gateway appropriately and the associated
  DynamoDB tables

# Building the Lambda functions

Assuming a sensible development environment and, in particular the installation
of the Go toolchain, it should be possible to simply run the Makefile to create
zipped executables for each of the lambda functions. (Note that there is a `sam build`
command which can build the application, but this did not work well for Go binaries
here - hence a solution which uses `make` is provided).

```
cd lambda
make
```

# Deploying the application to AWS

## Prerequisites

It is assumed that the AWS CLI (v2) is installed and there is a valid default
configuration. Also, the SAM CLI needs to be installed.

# Deployment

The application definition is provided in the `app.yaml` file as an AWS Serverless
Application Model. This contains all the information necessary to configure the
AWS API Gateway appropriately and to set up DynamoDB tables necessary to run the
application.

To deploy the application simply run

```
sam deploy -f app.yaml
```

On success, the Websockets endpoint is displayed - the FE needs to be modified to
include this parameter.



