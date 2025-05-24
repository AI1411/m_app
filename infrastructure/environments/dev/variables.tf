variable "aws_region" {
  description = "AWS region"
  type        = string
  default     = "ap-northeast-1" # 例: 東京リージョン
}

variable "project_name" {
  description = "Name of the project"
  type        = string
}

# VPC Variables
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

# RDS Variables
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
  sensitive   = true # パスワードは機密情報として扱う
}

variable "rds_skip_final_snapshot" {
  description = "Skip final snapshot on DB deletion"
  type        = bool
  default     = true # 開発環境などではtrue、本番ではfalse推奨
}

variable "rds_backup_retention_period" {
  description = "Backup retention period in days"
  type        = number
  default     = 7
}

variable "rds_preferred_backup_window" {
  description = "Preferred backup window"
  type        = string
  default     = "03:00-04:00" # JST 12:00-13:00
}

variable "rds_preferred_maintenance_window" {
  description = "Preferred maintenance window"
  type        = string
  default     = "sun:04:30-sun:05:30" # JST Sun 13:30-14:30
}

# API Service Variables (ECS Fargate例)
variable "container_image" {
  description = "Docker image for the API service"
  type        = string
}

variable "container_port" {
  description = "Port the container listens on"
  type        = number
}

variable "ecs_desired_count" {
  description = "Desired number of ECS tasks"
  type        = number
  default     = 1
}

variable "ecs_cpu" {
  description = "CPU units for ECS task"
  type        = number
  default     = 256 # 0.25 vCPU
}

variable "ecs_memory" {
  description = "Memory in MiB for ECS task"
  type        = number
  default     = 512 # 0.5 GB
}

variable "enable_nat_gateway" {
  description = "Whether to enable NAT Gateway"
  type        = bool
  default     = true
}

variable "single_nat_gateway" {
  description = "Whether to use a single NAT Gateway for all private subnets"
  type        = bool
  default     = false
}

variable "vpc_module_tags" {
  description = "Tags to apply to VPC resources"
  type        = map(string)
  default     = {}
}