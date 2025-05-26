output "alb_dns_name" {
  description = "DNS name of the Application Load Balancer"
  value       = module.api_service.alb_dns_name
}

output "ecr_repository_url" {
  description = "URL of the ECR repository"
  value       = module.ecr.repository_url
}

output "ecr_repository_name" {
  description = "Name of the ECR repository"
  value       = module.ecr.repository_name
}

output "rds_aurora_endpoint" {
  description = "Endpoint of the RDS Aurora cluster"
  value       = module.rds.db_cluster_endpoint
}

output "rds_aurora_reader_endpoint" {
  description = "Reader endpoint of the RDS Aurora cluster"
  value       = module.rds.db_cluster_reader_endpoint
}

output "vpc_id" {
  description = "ID of the created VPC"
  value       = module.vpc.vpc_id
}

# Secrets Manager outputs
output "db_password_secret_name" {
  description = "Name of the database password secret in Secrets Manager"
  value       = var.enable_secrets_manager ? module.secrets_manager[0].db_password_secret_name : ""
}

output "db_password_secret_arn" {
  description = "ARN of the database password secret in Secrets Manager"
  value       = var.enable_secrets_manager ? module.secrets_manager[0].db_password_secret_arn : ""
  sensitive   = true
}

output "app_secrets_name" {
  description = "Name of the application secrets in Secrets Manager"
  value       = var.enable_secrets_manager ? module.secrets_manager[0].app_secrets_name : ""
}

output "jwt_secret_name" {
  description = "Name of the JWT secret in Secrets Manager"
  value       = var.enable_secrets_manager ? module.secrets_manager[0].jwt_secret_name : ""
}