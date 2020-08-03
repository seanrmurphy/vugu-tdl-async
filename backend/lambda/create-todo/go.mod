module github.com/seanrmurphy/vugu-tdl-async/backend/lambda/create-todo

go 1.14

replace github.com/seanrmurphy/vugu-tdl-async/backend/lambda/types => ../types

replace github.com/seanrmurphy/vugu-tdl-async/backend/lambda/util => ../util

require (
	github.com/aws/aws-lambda-go v1.18.0
	github.com/aws/aws-sdk-go v1.33.14
	github.com/google/uuid v1.1.1
	github.com/seanrmurphy/go-fullstack v0.0.0-20200629071412-e58ac132051d
	github.com/seanrmurphy/vugu-tdl-async/backend/lambda/types v0.0.0-00010101000000-000000000000 // indirect
	github.com/seanrmurphy/vugu-tdl-async/backend/lambda/util v0.0.0-00010101000000-000000000000 // indirect
)
