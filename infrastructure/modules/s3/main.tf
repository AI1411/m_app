resource "aws_s3_bucket" "terraform_state" {
  bucket = var.bucket_name

  # バケットの削除防止（本番環境では有効にすることを推奨）
  # force_destroy = false

  tags = merge(
    {
      Name = var.bucket_name
    },
    var.tags
  )
}

# バージョニングの有効化
resource "aws_s3_bucket_versioning" "versioning" {
  bucket = aws_s3_bucket.terraform_state.id
  versioning_configuration {
    status = "Enabled"
  }
}

# サーバーサイド暗号化の有効化
resource "aws_s3_bucket_server_side_encryption_configuration" "encryption" {
  bucket = aws_s3_bucket.terraform_state.id

  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm = "AES256"
    }
  }
}

# バケットの公開アクセスをブロック
resource "aws_s3_bucket_public_access_block" "public_access" {
  bucket                  = aws_s3_bucket.terraform_state.id
  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

# DynamoDBテーブル（ステートロック用）
resource "aws_dynamodb_table" "terraform_locks" {
  name         = var.dynamodb_table_name
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "LockID"

  attribute {
    name = "LockID"
    type = "S"
  }

  tags = merge(
    {
      Name = var.dynamodb_table_name
    },
    var.tags
  )
}