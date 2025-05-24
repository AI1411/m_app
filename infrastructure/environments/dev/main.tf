terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0" # 適宜バージョンを指定してください
    }
  }
}

provider "aws" {
  region = var.aws_region
  # Assume Roleなどの設定が必要な場合はここに追加
}

# --- VPC Module ---
module "vpc" {
  source = "../../modules/vpc"

  project_name         = var.project_name
  vpc_cidr             = var.vpc_cidr
  public_subnet_cidrs  = var.public_subnet_cidrs
  private_subnet_cidrs = var.private_subnet_cidrs
  availability_zones   = var.availability_zones
  enable_nat_gateway = var.enable_nat_gateway   # modules/vpc/variables.tf で追加した変数を渡します
  single_nat_gateway = var.single_nat_gateway # modules/vpc/variables.tf で追加した変数を渡します
  tags                 = var.vpc_module_tags
}

module "security_group" {
  source = "../../modules/security_group"

  project_name = var.project_name
  vpc_id       = module.vpc.vpc_id
}

# --- RDS Aurora PostgreSQL Module ---
module "rds" {
  source = "../../modules/rds"

  project_name         = var.project_name
  db_instance_class    = var.db_instance_class
  db_name              = var.db_name
  db_username          = var.db_username
  db_password          = var.db_password # Sensitive: tfvarsや環境変数で管理推奨
  db_security_group_id = module.security_group.rds_sg_id
  private_subnet_ids   = module.vpc.private_subnet_ids
  vpc_id               = module.vpc.vpc_id
}

# --- API Service Module (例: ECS Fargate + ALB) ---
module "api_service" {
  source = "../../modules/api_service"

  project_name               = var.project_name
  vpc_id                     = module.vpc.vpc_id
  public_subnet_ids          = module.vpc.public_subnet_ids
  private_subnet_ids         = module.vpc.private_subnet_ids
  alb_security_group_id      = module.security_group.alb_sg_id
  ecs_task_security_group_id = module.security_group.ecs_task_sg_id

  container_image = var.container_image # APIサーバーのDockerイメージ
  container_port  = var.container_port
  desired_count   = var.ecs_desired_count
  cpu             = var.ecs_cpu
  memory          = var.ecs_memory
  aws_region      = var.aws_region # CloudWatch Logs などで使用

  # 環境変数やシークレットなどを渡す
  # environment_variables = {
  #   DB_HOST = module.rds_aurora_postgres.db_cluster_endpoint
  #   DB_PORT = module.rds_aurora_postgres.db_cluster_port
  #   DB_USER = var.db_username
  #   DB_NAME = var.db_name
  # }
  # secrets = {
  #   DB_PASSWORD = {
  #     valueFrom = var.db_password_secret_arn # AWS Secrets ManagerのARNなど
  #   }
  # }
}
