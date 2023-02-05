package main_test

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/go-playground/validator"
	main "github.com/olafszymanski/repostli/cmd/signup"
	"github.com/olafszymanski/repostli/pkg/response"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	t.Parallel()

	validRes, err := response.New(
		response.WithStatusCode(200),
		response.WithHeaders(main.Headers),
	)
	assert.NoError(t, err)

	errRes, err := response.New(
		response.WithStatusCode(400),
		response.WithHeaders(main.Headers),
		response.WithError(main.ErrInvalidRequest),
	)
	assert.NoError(t, err)

	tests := []struct {
		name     string
		request  events.APIGatewayProxyRequest
		response events.APIGatewayProxyResponse
	}{
		{
			name: "valid request",
			request: events.APIGatewayProxyRequest{
				Body: `{"email": "email", "password": "password"}`,
			},
			response: validRes,
		},
		{
			name: "no email",
			request: events.APIGatewayProxyRequest{
				Body: `{"password": "password"}`,
			},
			response: errRes,
		},
		{
			name: "no password",
			request: events.APIGatewayProxyRequest{
				Body: `{"email": "email"}`,
			},
			response: errRes,
		},
	}

	h := main.NewHandler(validator.New())

	for _, test := range tests {
		res, _ := h(context.TODO(), test.request)
		if res.Body != test.response.Body {
			t.Errorf("expected '%s', got '%s'", test.response.Body, res.Body)
		}
	}
}
