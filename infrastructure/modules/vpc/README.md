# VPC Module

このモジュールはAWS VPCとそれに関連するリソース（サブネット、ルートテーブル、インターネットゲートウェイ、NATゲートウェイなど）を作成します。

## 使用方法

```hcl
module "vpc" {
  source = "../../modules/vpc"

  project_name       = "myproject"
  vpc_cidr           = "10.0.0.0/16"
  public_subnet_cidrs = ["10.0.1.0/24", "10.0.2.0/24"]
  private_subnet_cidrs = ["10.0.3.0/24", "10.0.4.0/24"]
  availability_zones = ["ap-northeast-1a", "ap-northeast-1c"]
  enable_nat_gateway = true
  single_nat_gateway = true
  tags = {
    Environment = "dev"
  }
}
```

## 入力変数

| 名前                   | 説明                          | タイプ            | デフォルト   | 必須  |
|----------------------|-----------------------------|----------------|---------|-----|
| project_name         | プロジェクト名（リソース名のプレフィックスとして使用） | `string`       | n/a     | はい  |
| vpc_cidr             | VPCのCIDRブロック                | `string`       | n/a     | はい  |
| public_subnet_cidrs  | パブリックサブネットのCIDRブロックのリスト     | `list(string)` | n/a     | はい  |
| private_subnet_cidrs | プライベートサブネットのCIDRブロックのリスト    | `list(string)` | n/a     | はい  |
| availability_zones   | 使用するアベイラビリティゾーンのリスト         | `list(string)` | n/a     | はい  |
| enable_nat_gateway   | NATゲートウェイを有効にするかどうか         | `bool`         | `true`  | いいえ |
| single_nat_gateway   | 単一のNATゲートウェイを使用するかどうか       | `bool`         | `false` | いいえ |
| tags                 | リソースに追加するタグのマップ             | `map(string)`  | `{}`    | いいえ |

## 出力値

| 名前                 | 説明                      |
|--------------------|-------------------------|
| vpc_id             | 作成されたVPCのID             |
| public_subnet_ids  | 作成されたパブリックサブネットのIDのリスト  |
| private_subnet_ids | 作成されたプライベートサブネットのIDのリスト |
