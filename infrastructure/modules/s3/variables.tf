variable "bucket_name" {
  description = "Name of the S3 bucket for terraform state"
  type        = string
}

variable "dynamodb_table_name" {
  description = "Name of the DynamoDB table for terraform state locking"
  type        = string
}

variable "tags" {
  description = "Additional tags for resources"
  type        = map(string)
  default     = {}
}