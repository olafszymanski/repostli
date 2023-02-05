module "lambdas" {
  source      = "../modules/lambdas"

  environment = var.environment
}

module "api" {
  source      = "../modules/api/v1"

  environment = var.environment

  signup_lambda_invoke_arn = module.lambdas.signup_lambda_invoke_arn
}
