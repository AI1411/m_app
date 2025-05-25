# Database password secret
resource "aws_secretsmanager_secret" "db_password" {
  name                    = "${var.project_name}-${var.env}-db-password"
  description             = "Database password for ${var.project_name} ${var.env} environment"
  recovery_window_in_days = var.recovery_window_in_days

  tags = merge(
    {
      Name        = "${var.project_name}-${var.env}-db-password"
      Environment = var.env
    },
    var.tags
  )
}

# Database connection details with password
resource "aws_secretsmanager_secret_version" "db_password" {
  secret_id = aws_secretsmanager_secret.db_password.id
  secret_string = jsonencode({
    username            = var.db_username
    password            = var.db_password
    engine              = "postgres"
    host                = var.db_host
    port                = var.db_port
    dbname              = var.db_name
    dbInstanceIdentifier = var.db_instance_identifier
  })

  lifecycle {
    ignore_changes = [secret_string]
  }
}

# API keys and other application secrets (optional)
resource "aws_secretsmanager_secret" "app_secrets" {
  count                   = length(var.app_secrets) > 0 ? 1 : 0
  name                    = "${var.project_name}-${var.env}-app-secrets"
  description             = "Application secrets for ${var.project_name} ${var.env} environment"
  recovery_window_in_days = var.recovery_window_in_days

  tags = merge(
    {
      Name        = "${var.project_name}-${var.env}-app-secrets"
      Environment = var.env
    },
    var.tags
  )
}

resource "aws_secretsmanager_secret_version" "app_secrets" {
  count     = length(var.app_secrets) > 0 ? 1 : 0
  secret_id = aws_secretsmanager_secret.app_secrets[0].id
  secret_string = jsonencode(var.app_secrets)

  lifecycle {
    ignore_changes = [secret_string]
  }
}

# JWT secret for authentication (optional)
resource "aws_secretsmanager_secret" "jwt_secret" {
  count                   = var.create_jwt_secret ? 1 : 0
  name                    = "${var.project_name}-${var.env}-jwt-secret"
  description             = "JWT secret for ${var.project_name} ${var.env} environment"
  recovery_window_in_days = var.recovery_window_in_days

  tags = merge(
    {
      Name        = "${var.project_name}-${var.env}-jwt-secret"
      Environment = var.env
    },
    var.tags
  )
}

resource "aws_secretsmanager_secret_version" "jwt_secret" {
  count     = var.create_jwt_secret ? 1 : 0
  secret_id = aws_secretsmanager_secret.jwt_secret[0].id
  secret_string = jsonencode({
    jwt_secret = var.jwt_secret_value != "" ? var.jwt_secret_value : random_password.jwt_secret[0].result
  })

  lifecycle {
    ignore_changes = [secret_string]
  }
}

# Random password generation for JWT secret
resource "random_password" "jwt_secret" {
  count   = var.create_jwt_secret && var.jwt_secret_value == "" ? 1 : 0
  length  = 64
  special = true
}