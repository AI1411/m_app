variable "project_name" {
  description = "Name of the project"
  type        = string
}

variable "env" {
  description = "Environment (e.g., dev, staging, prod)"
  type        = string
}

variable "vpc_id" {
  description = "ID of the VPC"
  type        = string
}

variable "public_subnet_ids" {
  description = "List of public subnet IDs for ALB"
  type = list(string)
}

variable "private_subnet_ids" {
  description = "List of private subnet IDs for ECS tasks"
  type = list(string)
}

variable "alb_security_group_id" {
  description = "Security group ID for the ALB"
  type        = string
}

variable "ecs_task_security_group_id" {
  description = "Security group ID for the ECS tasks"
  type        = string
}

variable "container_image" {
  description = "Docker image for the API service (e.g., from ECR)"
  type        = string
}

variable "container_port" {
  description = "Port the container listens on"
  type        = number
}

variable "desired_count" {
  description = "Desired number of ECS tasks"
  type        = number
  default     = 1
}

variable "cpu" {
  description = "CPU units for ECS task (e.g., 256 for 0.25 vCPU)"
  type        = number
}

variable "memory" {
  description = "Memory in MiB for ECS task (e.g., 512 for 0.5GB)"
  type        = number
}

variable "health_check_path" {
  description = "Path for ALB health check"
  type        = string
  default     = "/"
}

variable "aws_region" {
  description = "AWS region for CloudWatch Logs"
  type        = string
}

# variable "acm_certificate_arn" {
#   description = "ARN of the ACM certificate for HTTPS listener (optional)"
#   type        = string
#   default     = ""
# }

# 環境変数やシークレット関連の変数
# variable "db_host" {
#   description = "Database host"
#   type        = string
# }
# variable "db_port" {
#   description = "Database port"
#   type        = number
# }
# variable "db_user" {
#   description = "Database user"
#   type        = string
# }
# variable "db_name" {
#   description = "Database name"
#   type        = string
# }
# variable "db_password_secret_arn" {
#   description = "ARN of the Secrets Manager secret for DB password"
#   type        = string
#   sensitive   = true
# }