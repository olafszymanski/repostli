package response

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type Response interface {
	Empty() (events.APIGatewayProxyResponse, error)
	WithBody(body map[string]string) (events.APIGatewayProxyResponse, error)
	WithError(err error) (events.APIGatewayProxyResponse, error)
}

type response struct {
	statusCode int
	headers    map[string]string
	body       string
}

func New(statusCode int, headers map[string]string) Response {
	return &response{
		statusCode: statusCode,
		headers:    headers,
		body:       "",
	}
}

func (r *response) Empty() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: r.statusCode,
		Headers:    r.headers,
		Body:       r.body,
	}, nil
}

func (r *response) WithBody(body map[string]string) (events.APIGatewayProxyResponse, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	r.body = string(b)

	return events.APIGatewayProxyResponse{
		StatusCode: r.statusCode,
		Headers:    r.headers,
		Body:       r.body,
	}, nil
}

func (r *response) WithError(err error) (events.APIGatewayProxyResponse, error) {
	e, merr := json.Marshal(map[string]string{
		"error": err.Error(),
	})
	if merr != nil {
		return events.APIGatewayProxyResponse{}, merr
	}

	r.body = string(e)

	return events.APIGatewayProxyResponse{
		StatusCode: r.statusCode,
		Headers:    r.headers,
		Body:       r.body,
	}, nil
}
