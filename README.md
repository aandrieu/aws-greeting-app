# Greeting App for AWS

Personal test for an AWS "micro-services" stack with terraform, rails, go and redis.

The app consists in a form where you put your name. It will then fetch a message to display for greeting you.
The app will also records how many time your name has been greeted.

- api: contains the go api that creates the messages
- dev: contains files to run the stack in local
- metadatastore: contains the rails api to communicate with the redis instance that records how many time your name has been entered
- stack: contains the terraform code to deploy the stack on aws
- ui: contains the code for the front-end (vanilla js)

## How to run the stack locally?

Run `docker-compose up` then go to `http://localhost:8080`.

## How to run the stack on AWS?

Rename `terraform.tfvars.example` to `terraform.tfvars` and fill the variables.

Run `terraform plan` to check that everything is ok. If all is ok run `terraform apply` to setup the stack.

The stack is composed of:
- a VPC
- 2 public subnets (used by the alb)
- 2 private subnets (used by fargate)
- a fargate cluster (with services and repositories for the api and the metadatastore)
- 1 bucket s3 (for the ui)
- 1 ALB (it serves the same purpose than nginx in local)
- 1 redis cluster

When the infrastructure is setup, you have to deploy the code.

## How to deploy the code on AWS?

coming...