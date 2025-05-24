output "alb_sg_id" {
  description = "ID of the ALB security group"
  value       = aws_security_group.alb.id
}

output "ecs_task_sg_id" {
  description = "ID of the API Service (ECS Task) security group"
  value       = aws_security_group.api_service.id
}

output "rds_sg_id" {
  description = "ID of the RDS security group"
  value       = aws_security_group.rds.id
}