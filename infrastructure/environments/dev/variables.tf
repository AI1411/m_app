variable "aws_region" {
  description = "AWS region"
  type        = string
  default     = "ap-northeast-1"
}

variable "project_name" {
  description = "Name of the project"
  type        = string
}

variable "env" {
  description = "Environment (e.g., dev, staging, prod)"
  type        = string
}

variable "vpc_cidr" {
  description = "CIDR block for the VPC"
  type        = string
}

variable "public_subnet_cidrs" {
  description = "List of CIDR blocks for public subnets"
  type        = list(string)
}

variable "private_subnet_cidrs" {
  description = "List of CIDR blocks for private subnets"
  type        = list(string)
}

variable "availability_zones" {
  description = "List of availability zones"
  type        = list(string)
}

# RDS関連変数
variable "db_instance_class" {
  description = "RDS instance class"
  type        = string
}

variable "db_name" {
  description = "Database name"
  type        = string
}

variable "db_username" {
  description = "Database master username"
  type        = string
}

variable "db_password" {
  description = "Database master password"
  type        = string
  sensitive   = true
}

variable "rds_skip_final_snapshot" {
  description = "Skip final snapshot on DB deletion"
  type        = bool
  default     = true
}

variable "rds_backup_retention_period" {
  description = "Backup retention period in days"
  type        = number
  default     = 7
}

variable "rds_preferred_backup_window" {
  description = "Preferred backup window"
  type        = string
  default     = "03:00-04:00"
}

variable "rds_preferred_maintenance_window" {
  description = "Preferred maintenance window"
  type        = string
  default     = "sun:04:30-sun:05:30"
}

# ECS関連変数
variable "container_image_tag" {
  description = "Tag for the Docker image in ECR (e.g., latest, v1.0.0)"
  type        = string
  default     = "latest"
}

variable "container_port" {
  description = "Port the container listens on"
  type        = number
  default     = 8080
}

variable "ecs_desired_count" {
  description = "Desired number of ECS tasks"
  type        = number
  default     = 1
}

variable "ecs_cpu" {
  description = "CPU units for ECS task"
  type        = number
  default     = 256
}

variable "ecs_memory" {
  description = "Memory in MiB for ECS task"
  type        = number
  default     = 512
}

variable "health_check_path" {
  description = "Path for ALB health check"
  type        = string
  default     = "/health"
}

# VPC関連変数
variable "enable_nat_gateway" {
  description = "Whether to enable NAT Gateway"
  type        = bool
  default     = true
}

variable "single_nat_gateway" {
  description = "Whether to use a single NAT Gateway for all private subnets"
  type        = bool
  default     = true
}

variable "vpc_module_tags" {
  description = "Tags to apply to VPC resources"
  type        = map(string)
  default     = {}
}

# Secrets Manager関連変数
variable "enable_secrets_manager" {
  description = "Whether to enable Secrets Manager"
  type        = bool
  default     = true
}

variable "app_secrets" {
  description = "Map of application secrets (key-value pairs)"
  type        = map(string)
  default     = {}
  sensitive   = true
}

variable "create_jwt_secret" {
  description = "Whether to create a JWT secret"
  type        = bool
  default     = false
}

variable "jwt_secret_value" {
  description = "JWT secret value (if empty, a random value will be generated)"
  type        = string
  default     = ""
  sensitive   = true
}

# Bastion Host関連変数
variable "enable_bastion" {
  description = "Whether to create a bastion host for database access"
  type        = bool
  default     = false
}

variable "bastion_public_key" {
  description = "Public key for SSH access to bastion host"
  type        = string
  default     = ""
}

variable "bastion_instance_type" {
  description = "EC2 instance type for bastion host"
  type        = string
  default     = "t3.micro"
}

variable "bastion_allowed_ssh_cidrs" {
  description = "CIDR blocks allowed to SSH to bastion host"
  type        = list(string)
  default     = ["0.0.0.0/0"]
}

variable "secrets_recovery_window_days" {
  description = "Number of days that AWS Secrets Manager waits before it can delete the secret"
  type        = number
  default     = 7
}