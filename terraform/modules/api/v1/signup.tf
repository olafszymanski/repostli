resource "aws_api_gateway_resource" "signup" {
  rest_api_id = aws_api_gateway_rest_api.v1.id
  parent_id   = aws_api_gateway_rest_api.v1.root_resource_id
  path_part   = "signup"
}

resource "aws_api_gateway_method" "signup_post" {
  rest_api_id      = aws_api_gateway_rest_api.v1.id
  resource_id      = aws_api_gateway_resource.signup.id
  http_method      = "POST"
  authorization    = "NONE"
  api_key_required = false
}

resource "aws_api_gateway_integration" "signup" {
  rest_api_id             = aws_api_gateway_rest_api.v1.id
  resource_id             = aws_api_gateway_resource.signup.id
  http_method             = aws_api_gateway_method.signup_post.http_method
  type                    = "AWS_PROXY"
  integration_http_method = "POST"
  uri                     = var.signup_lambda_invoke_arn
}
