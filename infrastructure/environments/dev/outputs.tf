output "alb_dns_name" {
  description = "DNS name of the Application Load Balancer"
  value       = module.api_service.alb_dns_name # api_service モジュールからの出力
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