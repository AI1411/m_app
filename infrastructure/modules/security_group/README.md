# Security Group Module

このモジュールはAWSセキュリティグループを作成し、アプリケーションの各コンポーネント（ALB、ECSタスク、RDS）に必要なセキュリティグループを設定します。

## 使用方法

```hcl
module "security_group" {
  source = "../../modules/security_group"

  project_name = "myproject"
  vpc_id       = module.vpc.vpc_id
}
```

## 入力変数

| 名前         | 説明                                  | タイプ      | デフォルト | 必須 |
|-------------|--------------------------------------|-----------|----------|-----|
| project_name | プロジェクト名（リソース名のプレフィックスとして使用） | `string`  | n/a      | はい |
| vpc_id      | セキュリティグループを作成するVPCのID        | `string`  | n/a      | はい |

## 出力値

| 名前           | 説明                                |
|---------------|-----------------------------------|
| alb_sg_id     | ALB用セキュリティグループのID            |
| ecs_task_sg_id | ECSタスク用セキュリティグループのID       |
| rds_sg_id     | RDS用セキュリティグループのID            |