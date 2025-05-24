# ECR Module

This module creates an Amazon Elastic Container Registry (ECR) repository for storing Docker images of the API service.

## Resources Created

- ECR Repository
- ECR Lifecycle Policy (to limit the number of stored images)

## Usage

```hcl
module "ecr" {
  source = "../../modules/ecr"

  project_name = "my-project"
  
  # Optional parameters
  image_tag_mutability = "MUTABLE"
  scan_on_push         = true
  encryption_type      = "AES256"
  max_image_count      = 10
}
```

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| project_name | Name of the project | `string` | n/a | yes |
| image_tag_mutability | The tag mutability setting for the repository (MUTABLE or IMMUTABLE) | `string` | `"MUTABLE"` | no |
| scan_on_push | Indicates whether images are scanned after being pushed to the repository | `bool` | `true` | no |
| encryption_type | The encryption type to use for the repository (AES256 or KMS) | `string` | `"AES256"` | no |
| max_image_count | Maximum number of images to keep in the repository | `number` | `10` | no |

## Outputs

| Name | Description |
|------|-------------|
| repository_url | The URL of the ECR repository |
| repository_arn | The ARN of the ECR repository |
| repository_name | The name of the ECR repository |