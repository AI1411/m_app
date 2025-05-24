terraform {
  backend "s3" {
    bucket         = "mapp-terraform-state-bucket"
    key            = "environments/dev/terraform.tfstate"
    region         = "ap-northeast-1"
    encrypt        = true
  }
}