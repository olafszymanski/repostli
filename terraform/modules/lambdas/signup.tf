resource "aws_iam_role" "signup" {
  name = "signup-role-${title(var.environment)}"
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          "Service" = "lambda.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_lambda_function" "signup" {
  function_name = "signup-${var.environment}"
  role          = aws_iam_role.signup.arn

  source_code_hash = filebase64sha256("../../../build/signup.zip")
  filename         = "../../../build/signup.zip"
  runtime          = "go1.2x"

  timeout = 10
}
