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
