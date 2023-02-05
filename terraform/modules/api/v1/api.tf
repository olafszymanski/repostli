resource "aws_api_gateway_rest_api" "v1" {
  name = "rest-api-v1-${var.environment}"

  endpoint_configuration {
    types = ["REGIONAL"]
  }
}
