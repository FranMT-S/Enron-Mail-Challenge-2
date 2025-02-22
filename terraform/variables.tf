variable "region" {
  description = "region aws ec2"
  type        = string
}

variable "availability_zone" {
  description = "available zone to volumen of aws ec2"
  type        = string
}

variable "access_key" {
  description = "IAM acces key"
  type        = string
  sensitive   = true
}

variable "secret_key" {
  description = "IAM secret key"
  type        = string
  sensitive   = true
}

variable "ami" {
  description = "ami id of instance must be valid in the region"
  type = string
}

variable "instance_type" {
  description = "type of the instance server"
  type = string
}

variable "key_name" {
  description = "key ssh to connect instance"
  type = string
}