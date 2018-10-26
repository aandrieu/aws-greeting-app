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

coming...