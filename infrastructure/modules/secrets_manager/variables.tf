variable "project_name" {
  description = "Name of the project"
  type        = string
}

variable "env" {
  description = "Environment (e.g., dev, staging, prod)"
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

variable "db_host" {
  description = "Database host endpoint"
  type        = string
}

variable "db_port" {
  description = "Database port"
  type        = number
  default     = 5432
}

variable "db_name" {
  description = "Database name"
  type        = string
}

variable "db_instance_identifier" {
  description = "Database instance identifier"
  type        = string
}

variable "recovery_window_in_days" {
  description = "Number of days that AWS Secrets Manager waits before it can delete the secret"
  type        = number
  default     = 7
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

variable "tags" {
  description = "A map of tags to add to all resources"
  type        = map(string)
  default     = {}
}