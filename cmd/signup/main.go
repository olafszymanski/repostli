package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-playground/validator"
)

type input struct {
	Email    string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

var headers = map[string]string{
	"Content-Type":                 "application/json",
	"Access-Control-Allow-Headers": "Content-Type",
	"Access-Control-Allow-Methods": "POST",
	"Access-Control-Allow-Origin":  "*",
}

var (
	ErrUnmarshalRequest = fmt.Errorf("failed to unmarshal request")
	ErrInvalidRequest   = fmt.Errorf("invalid request")
)

type RequestHandler func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

func NewRequestHandler(validator *validator.Validate) RequestHandler {
	return func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		var input input
		if err := json.Unmarshal([]byte(request.Body), &input); err != nil {
			err = fmt.Errorf("%w: %s", ErrUnmarshalRequest, err)
			return events.APIGatewayProxyResponse{
				StatusCode: 400,
				Headers:    headers,
				Body:       err.Error(),
			}, err
		}

		if err := validator.Struct(input); err != nil {
			err = fmt.Errorf("%w: %s", ErrInvalidRequest, err)
			return events.APIGatewayProxyResponse{
				StatusCode: 400,
				Headers:    headers,
				Body:       err.Error(),
			}, err
		}

		// TODO: Create user in database

		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Headers:    headers,
			Body:       fmt.Sprintf("Hello %s!", input.Email),
		}, nil
	}
}

func main() {
	lambda.Start(NewRequestHandler(validator.New()))
}
