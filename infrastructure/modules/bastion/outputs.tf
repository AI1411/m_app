output "bastion_public_ip" {
  description = "Public IP address of the bastion host"
  value       = var.enable_eip ? aws_eip.bastion[0].public_ip : aws_instance.bastion.public_ip
}

output "bastion_private_ip" {
  description = "Private IP address of the bastion host"
  value       = aws_instance.bastion.private_ip
}

output "bastion_instance_id" {
  description = "Instance ID of the bastion host"
  value       = aws_instance.bastion.id
}

output "ssh_command" {
  description = "SSH command to connect to bastion host"
  value       = "ssh -i ~/.ssh/${aws_key_pair.bastion.key_name}.pem ec2-user@${var.enable_eip ? aws_eip.bastion[0].public_ip : aws_instance.bastion.public_ip}"
}

output "rds_connection_command" {
  description = "Command to connect to RDS from bastion"
  value       = "psql -h ${var.db_endpoint} -U m_app_admin -d m_app_dev"
}