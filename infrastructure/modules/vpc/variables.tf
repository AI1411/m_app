variable "project_name" {
  description = "Name of the project"
  type        = string
}

variable "vpc_cidr" {
  description = "CIDR block for the VPC"
  type        = string
}

variable "public_subnet_cidrs" {
  description = "List of CIDR blocks for public subnets"
  type = list(string)
}

variable "private_subnet_cidrs" {
  description = "List of CIDR blocks for private subnets"
  type = list(string)
}

variable "availability_zones" {
  description = "List of availability zones for the subnets"
  type = list(string)
}

variable "enable_nat_gateway" {
  description = "Set to true to create a NAT Gateway. If false, private subnets will not have internet access."
  type        = bool
  default     = true
}

variable "single_nat_gateway" {
  description = "Set to true to create a single NAT Gateway shared across all AZs. Set to false to create one NAT Gateway per AZ."
  type        = bool
  default     = true
}

variable "tags" {
  description = "A map of tags to add to all resources."
  type = map(string)
  default = {}
}