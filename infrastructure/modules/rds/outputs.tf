output "db_cluster_endpoint" {
  description = "The connection endpoint for the RDS cluster writer instance"
  value       = aws_rds_cluster.default.endpoint
}

output "db_cluster_reader_endpoint" {
  description = "The connection endpoint for the RDS cluster reader instances"
  value       = aws_rds_cluster.default.reader_endpoint
}

output "db_cluster_port" {
  description = "The port for the RDS cluster"
  value       = aws_rds_cluster.default.port
}

output "db_cluster_id" {
  description = "The ID of the RDS cluster"
  value       = aws_rds_cluster.default.id
}