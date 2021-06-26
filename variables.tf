variable "access_key" {
  type = string
}

variable "secret_key" {
  type = string
}

variable "region" {
  description = "Region where we will create the bucket"
  default     = "us-west-2"
}
