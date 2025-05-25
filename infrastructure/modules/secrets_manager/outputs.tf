output "db_password_secret_arn" {
  description = "ARN of the database password secret"
  value       = aws_secretsmanager_secret.db_password.arn
}

output "db_password_secret_name" {
  description = "Name of the database password secret"
  value       = aws_secretsmanager_secret.db_password.name
}

output "app_secrets_arn" {
  description = "ARN of the application secrets"
  value       = length(aws_secretsmanager_secret.app_secrets) > 0 ? aws_secretsmanager_secret.app_secrets[0].arn : ""
}

output "app_secrets_name" {
  description = "Name of the application secrets"
  value       = length(aws_secretsmanager_secret.app_secrets) > 0 ? aws_secretsmanager_secret.app_secrets[0].name : ""
}

output "jwt_secret_arn" {
  description = "ARN of the JWT secret"
  value       = length(aws_secretsmanager_secret.jwt_secret) > 0 ? aws_secretsmanager_secret.jwt_secret[0].arn : ""
}

output "jwt_secret_name" {
  description = "Name of the JWT secret"
  value       = length(aws_secretsmanager_secret.jwt_secret) > 0 ? aws_secretsmanager_secret.jwt_secret[0].name : ""
}