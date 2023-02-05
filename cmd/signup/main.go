package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-playground/validator"
	"github.com/olafszymanski/repostli/pkg/response"
)

type input struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

var Headers = map[string]string{
	"Content-Type":                 "application/json",
	"Access-Control-Allow-Headers": "Content-Type",
	"Access-Control-Allow-Methods": "POST",
	"Access-Control-Allow-Origin":  "*",
}

var ErrInvalidRequest = fmt.Errorf("invalid request")

type handler func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

func NewHandler(validator *validator.Validate) handler {
	return func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		var input input
		if err := json.Unmarshal([]byte(request.Body), &input); err != nil {
			return response.New(
				response.WithStatusCode(400),
				response.WithHeaders(Headers),
				response.WithError(ErrInvalidRequest),
			)
		}

		if err := validator.Struct(input); err != nil {
			return response.New(
				response.WithStatusCode(400),
				response.WithHeaders(Headers),
				response.WithError(ErrInvalidRequest),
			)
		}

		// TODO: Create user in database

		return response.New(
			response.WithStatusCode(200),
			response.WithHeaders(Headers),
		)
	}
}

func main() {
	lambda.Start(NewHandler(validator.New()))
}
