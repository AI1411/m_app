# RDS Module

このモジュールはAmazon RDS（Aurora PostgreSQL）クラスターとそれに関連するリソースを作成します。

## 使用方法

``` hcl
module "rds" {
  source = "../../modules/rds"

  project_name         = "myproject"
  db_instance_class    = "db.t3.medium"
  db_name              = "appdb"
  db_username          = "admin"
  db_password          = var.db_password
  db_security_group_id = module.security_group.db_security_group_id
  private_subnet_ids   = module.vpc.private_subnet_ids
  vpc_id               = module.vpc.vpc_id
}
```

## 入力変数

| 名前                   | 説明                          | タイプ            | デフォルト | 必須 |
|----------------------|-----------------------------|----------------|-------|----|
| project_name         | プロジェクト名（リソース名のプレフィックスとして使用） | `string`       | n/a   | はい |
| db_instance_class    | DBインスタンスクラス                 | `string`       | n/a   | はい |
| db_name              | データベース名                     | `string`       | n/a   | はい |
| db_username          | データベースのマスターユーザー名            | `string`       | n/a   | はい |
| db_password          | データベースのマスターパスワード            | `string`       | n/a   | はい |
| db_security_group_id | データベース用セキュリティグループID         | `string`       | n/a   | はい |
| private_subnet_ids   | DBを配置するプライベートサブネットIDのリスト    | `list(string)` | n/a   | はい |
| vpc_id               | VPC ID                      | `string`       | n/a   | はい |

## 出力値

| 名前                  | 説明                    |
|---------------------|-----------------------|
| db_cluster_endpoint | RDS クラスターのエンドポイント     |
| db_cluster_port     | RDS クラスターのポート         |
| db_instance_ids     | 作成されたRDSインスタンスのIDのリスト |
