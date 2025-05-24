# 開発環境 (Dev Environment)

このディレクトリは開発環境のTerraform設定を含んでいます。

## 前提条件

- Terraform 1.0.0以上
- AWS CLIがインストールされ、適切に設定されていること
- 適切なAWS権限を持つIAMユーザーまたはロール

## 使用方法

1. 必要な変数を設定します。`terraform.tfvars`ファイルを作成するか、環境変数を使用します。

```hcl
# terraform.tfvars の例
project_name = "myproject-dev"
vpc_cidr     = "10.0.0.0/16"
public_subnet_cidrs = ["10.0.1.0/24", "10.0.2.0/24"]
private_subnet_cidrs = ["10.0.3.0/24", "10.0.4.0/24"]
availability_zones = ["ap-northeast-1a", "ap-northeast-1c"]

db_instance_class = "db.t3.small"
db_name           = "mydb"
db_username       = "admin"
# db_passwordは環境変数で設定することを推奨: export TF_VAR_db_password="your-secure-password"

container_image = "123456789012.dkr.ecr.ap-northeast-1.amazonaws.com/myproject:latest"
container_port  = 8080
```

2. 初期化とプランの実行

```bash
terraform init
terraform plan
```

3. リソースの作成

```bash
terraform apply
```

4. リソースの削除（必要な場合）

```bash
terraform destroy
```

## 注意事項

- 開発環境用の設定のため、コスト削減のためにいくつかのリソースは最小限の設定になっています
- 本番環境では、高可用性やセキュリティを考慮した設定に変更することを推奨します
- 機密情報（パスワードなど）は環境変数やAWS Secrets Managerなどを使用して管理してください