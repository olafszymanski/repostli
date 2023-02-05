package response

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type response struct {
	statusCode int
	headers    map[string]string
	body       string
}

type Options func(options *response) error

func New(options ...Options) (events.APIGatewayProxyResponse, error) {
	r := &response{}
	for _, opt := range options {
		err := opt(r)
		if err != nil {
			return events.APIGatewayProxyResponse{}, err
		}
	}

	return events.APIGatewayProxyResponse{
		StatusCode: r.statusCode,
		Headers:    r.headers,
		Body:       r.body,
	}, nil
}

func WithStatusCode(statusCode int) Options {
	return func(options *response) error {
		options.statusCode = statusCode
		return nil
	}
}

func WithHeaders(headers map[string]string) Options {
	return func(options *response) error {
		options.headers = headers
		return nil
	}
}

func WithBody(body map[string]string) Options {
	return func(options *response) error {
		b, err := json.Marshal(body)
		if err != nil {
			return err
		}

		options.body = string(b)
		return nil
	}
}

func WithError(err error) Options {
	return func(options *response) error {
		e, err := json.Marshal(map[string]string{
			"error": err.Error(),
		})
		if err != nil {
			return err
		}

		options.body = string(e)
		return nil
	}
}
