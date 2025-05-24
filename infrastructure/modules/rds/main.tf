resource "aws_rds_cluster_parameter_group" "aurora_pg_cluster_params" {
  name        = "${var.project_name}-${var.env}-aurora-pg15-cluster-params"
  family      = "aurora-postgresql15"
  description = "Aurora PostgreSQL 15 cluster parameter group for ${var.project_name}-${var.env}"
}

resource "aws_db_subnet_group" "default" {
  name       = "${var.project_name}-${var.env}-rds-subnet-group"
  subnet_ids = var.private_subnet_ids

  tags = {
    Name = "${var.project_name}-${var.env}-rds-subnet-group"
  }
}

resource "aws_rds_cluster" "default" {
  cluster_identifier           = "${var.project_name}-${var.env}-aurora-cluster"
  engine                       = "aurora-postgresql"
  engine_version               = "15.5"
  database_name                = var.db_name
  master_username              = var.db_username
  master_password              = var.db_password
  db_subnet_group_name         = aws_db_subnet_group.default.name
  vpc_security_group_ids = [var.db_security_group_id]
  skip_final_snapshot          = var.skip_final_snapshot
  backup_retention_period      = var.backup_retention_period
  preferred_backup_window      = var.preferred_backup_window
  preferred_maintenance_window = var.preferred_maintenance_window
  storage_encrypted            = true

  tags = {
    Name = "${var.project_name}-${var.env}-aurora-cluster"
  }
}

resource "aws_rds_cluster_instance" "default" {
  count              = 1
  identifier         = "${var.project_name}-${var.env}-aurora-instance-${count.index}"
  cluster_identifier = aws_rds_cluster.default.id
  instance_class     = var.db_instance_class
  engine             = "aurora-postgresql"
  engine_version     = aws_rds_cluster.default.engine_version

  tags = {
    Name = "${var.project_name}-${var.env}-aurora-instance-${count.index}"
  }
}