output "signup_lambda_invoke_arn" {
  description = "Arn of the signup lambda function"
  value       = aws_lambda_function.signup.invoke_arn
}