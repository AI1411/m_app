terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
}

module "vpc" {
  source = "../../modules/vpc"

  project_name         = var.project_name
  vpc_cidr             = var.vpc_cidr
  public_subnet_cidrs  = var.public_subnet_cidrs
  private_subnet_cidrs = var.private_subnet_cidrs
  availability_zones   = var.availability_zones
  enable_nat_gateway   = var.enable_nat_gateway
  single_nat_gateway   = var.single_nat_gateway
  tags                 = var.vpc_module_tags
}

module "security_group" {
  source = "../../modules/security_group"

  project_name = var.project_name
  vpc_id       = module.vpc.vpc_id
}

module "rds" {
  source = "../../modules/rds"

  project_name                 = var.project_name
  env                          = var.env
  db_instance_class            = var.db_instance_class
  db_name                      = var.db_name
  db_username                  = var.db_username
  db_password                  = var.db_password
  db_security_group_id         = module.security_group.rds_sg_id
  private_subnet_ids           = module.vpc.private_subnet_ids
  vpc_id                       = module.vpc.vpc_id
  skip_final_snapshot          = var.rds_skip_final_snapshot
  backup_retention_period      = var.rds_backup_retention_period
  preferred_backup_window      = var.rds_preferred_backup_window
  preferred_maintenance_window = var.rds_preferred_maintenance_window
}

module "ecr" {
  source = "../../modules/ecr"

  project_name = var.project_name
  env          = var.env
}

module "api_service" {
  source = "../../modules/api_service"

  project_name          = var.project_name
  env                   = var.env
  vpc_id                = module.vpc.vpc_id
  public_subnet_ids     = module.vpc.public_subnet_ids
  private_subnet_ids    = module.vpc.private_subnet_ids
  alb_security_group_id = module.security_group.alb_sg_id
  ecs_task_security_group_id = module.security_group.ecs_task_sg_id

  # Use the ECR repository URL with the image tag
  container_image   = "${module.ecr.repository_url}:${var.container_image_tag}"
  container_port    = var.container_port
  desired_count     = var.ecs_desired_count
  cpu               = var.ecs_cpu
  memory            = var.ecs_memory
  aws_region        = var.aws_region
  health_check_path = var.health_check_path
}
