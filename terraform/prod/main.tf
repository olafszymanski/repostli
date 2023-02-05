module "lambdas" {
  source      = "../modules/lambdas"
  environment = var.environment
}