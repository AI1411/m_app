output "bucket_id" {
  description = "ID of the S3 bucket"
  value       = aws_s3_bucket.terraform_state.id
}

output "bucket_arn" {
  description = "ARN of the S3 bucket"
  value       = aws_s3_bucket.terraform_state.arn
}

output "dynamodb_table_id" {
  description = "ID of the DynamoDB table"
  value       = aws_dynamodb_table.terraform_locks.id
}

output "dynamodb_table_arn" {
  description = "ARN of the DynamoDB table"
  value       = aws_dynamodb_table.terraform_locks.arn
}