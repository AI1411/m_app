resource "aws_rds_cluster_parameter_group" "aurora_pg_cluster_params" {
  name        = "${var.project_name}-aurora-pg15-cluster-params" # PostgreSQLのバージョンに合わせて変更
  family      = "aurora-postgresql15" # PostgreSQLのバージョンに合わせて変更
  description = "Aurora PostgreSQL 15 cluster parameter group for ${var.project_name}"

  # 必要に応じてパラメータを設定
  # parameter {
  #   name  = "log_statement"
  #   value = "ddl"
  # }
}

resource "aws_db_subnet_group" "default" {
  name       = "${var.project_name}-rds-subnet-group"
  subnet_ids = var.private_subnet_ids # RDSはプライベートサブネットに配置

  tags = {
    Name = "${var.project_name}-rds-subnet-group"
  }
}

resource "aws_rds_cluster" "default" {
  cluster_identifier              = "${var.project_name}-aurora-cluster"
  engine                          = "aurora-postgresql"
  engine_version                  = "15.5" # 利用可能な最新バージョンを確認・指定
  database_name                   = var.db_name
  master_username                 = var.db_username
  master_password                 = var.db_password
  db_subnet_group_name            = aws_db_subnet_group.default.name
  vpc_security_group_ids          = [var.db_security_group_id]
  skip_final_snapshot             = var.skip_final_snapshot
  backup_retention_period         = var.backup_retention_period
  preferred_backup_window         = var.preferred_backup_window
  preferred_maintenance_window    = var.preferred_maintenance_window
  storage_encrypted               = true # 暗号化を有効化
  # db_cluster_parameter_group_name = aws_rds_cluster_parameter_group.aurora_pg_cluster_params.name # 必要であれば指定

  # 削除保護 (本番環境では有効化を推奨)
  # deletion_protection = false

  tags = {
    Name = "${var.project_name}-aurora-cluster"
  }
}

resource "aws_rds_cluster_instance" "default" {
  count              = 1 # 必要なインスタンス数に応じて変更 (通常はリーダー1台 + ライター1台など)
  identifier         = "${var.project_name}-aurora-instance-${count.index}"
  cluster_identifier = aws_rds_cluster.default.id
  instance_class     = var.db_instance_class
  engine             = "aurora-postgresql"
  engine_version     = aws_rds_cluster.default.engine_version
  # db_parameter_group_name = aws_db_parameter_group.aurora_pg_instance_params.name # 必要であればインスタンスパラメータグループも作成・指定

  tags = {
    Name = "${var.project_name}-aurora-instance-${count.index}"
  }
}