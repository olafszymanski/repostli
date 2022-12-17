variable "prefix" {
  type        = string
  default     = "repostli"
  description = "Prefix for all resources"
}

variable "environment" {
  type        = string
  default     = "production"
  description = "Environment type"
}
