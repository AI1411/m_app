variable "project_name" {
  description = "Name of the project"
  type        = string
}

variable "env" {
  description = "Environment (e.g., dev, staging, prod)"
  type        = string
}

variable "vpc_id" {
  description = "VPC ID where bastion host will be created"
  type        = string
}

variable "public_subnet_id" {
  description = "Public subnet ID for bastion host"
  type        = string
}

variable "rds_security_group_id" {
  description = "RDS security group ID to allow access from bastion"
  type        = string
}

variable "db_endpoint" {
  description = "RDS database endpoint"
  type        = string
}

variable "public_key" {
  description = "Public key for SSH access to bastion host"
  type        = string
}

variable "instance_type" {
  description = "EC2 instance type for bastion host"
  type        = string
  default     = "t3.micro"
}

variable "allowed_ssh_cidrs" {
  description = "CIDR blocks allowed to SSH to bastion host"
  type        = list(string)
  default     = ["0.0.0.0/0"]  # 本番環境では制限することを推奨
}

variable "enable_eip" {
  description = "Whether to create an Elastic IP for bastion host"
  type        = bool
  default     = true
}