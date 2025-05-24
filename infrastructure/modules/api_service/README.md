# API Service Module

このモジュールはAWS ECS Fargateを使用したAPIサービスとApplication Load Balancerを作成します。

## 使用方法

```hcl
module "api_service" {
  source = "../../modules/api_service"

  project_name               = "myproject"
  vpc_id                     = module.vpc.vpc_id
  public_subnet_ids          = module.vpc.public_subnet_ids
  private_subnet_ids         = module.vpc.private_subnet_ids
  alb_security_group_id      = module.security_group.alb_sg_id
  ecs_task_security_group_id = module.security_group.ecs_task_sg_id

  container_image = "123456789012.dkr.ecr.ap-northeast-1.amazonaws.com/myproject:latest"
  container_port  = 8080
  desired_count   = 2
  cpu             = 512
  memory          = 1024
  aws_region      = "ap-northeast-1"
}
```

## 入力変数

| 名前                      | 説明                                  | タイプ            | デフォルト | 必須 |
|--------------------------|--------------------------------------|----------------|----------|-----|
| project_name             | プロジェクト名（リソース名のプレフィックスとして使用） | `string`       | n/a      | はい |
| vpc_id                   | VPCのID                              | `string`       | n/a      | はい |
| public_subnet_ids        | パブリックサブネットのIDのリスト            | `list(string)` | n/a      | はい |
| private_subnet_ids       | プライベートサブネットのIDのリスト          | `list(string)` | n/a      | はい |
| alb_security_group_id    | ALBのセキュリティグループID              | `string`       | n/a      | はい |
| ecs_task_security_group_id | ECSタスクのセキュリティグループID        | `string`       | n/a      | はい |
| container_image          | コンテナイメージのURI                   | `string`       | n/a      | はい |
| container_port           | コンテナがリッスンするポート              | `number`       | `8080`   | いいえ |
| desired_count            | 実行するタスクの数                      | `number`       | `1`      | いいえ |
| cpu                      | タスクに割り当てるCPUユニット             | `number`       | `256`    | いいえ |
| memory                   | タスクに割り当てるメモリ (MB)            | `number`       | `512`    | いいえ |
| aws_region               | AWSリージョン                          | `string`       | n/a      | はい |

## 出力値

| 名前                | 説明                                |
|--------------------|-----------------------------------|
| alb_dns_name       | Application Load BalancerのDNS名    |
| alb_arn            | Application Load BalancerのARN     |
| ecs_cluster_name   | ECSクラスターの名前                   |
| ecs_service_name   | ECSサービスの名前                    |