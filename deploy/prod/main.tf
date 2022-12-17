terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.47.0"
    }
  }
}

provider "aws" {
  region = "us-east-1"
}

module "api" {
  source      = "../modules/api"
  prefix      = var.prefix
  environment = var.environment
}
