# Secrets Manager Module

このモジュールはAWS Secrets Managerを使用して、データベースパスワードやアプリケーションシークレットを安全に管理します。

## 使用方法

```hcl
module "secrets_manager" {
  source = "../../modules/secrets_manager"

  project_name = "myproject"
  env          = "dev"
  
  # データベース接続情報
  db_username             = "admin"
  db_password             = var.db_password
  db_host                 = module.rds.db_cluster_endpoint
  db_port                 = module.rds.db_cluster_port
  db_name                 = "myapp"
  db_instance_identifier = "myproject-dev-aurora-cluster"
  
  # アプリケーションシークレット（オプション）
  app_secrets = {
    api_key         = "your-api-key"
    encryption_key  = "your-encryption-key"
    webhook_secret  = "your-webhook-secret"
  }
  
  # JWT シークレット（オプション）
  create_jwt_secret = true
  jwt_secret_value  = "" # 空の場合、ランダムに生成される
  
  # 設定
  recovery_window_in_days = 7
  
  tags = {
    Environment = "dev"
  }
}
```

## 機能

### 1. データベースシークレット
- RDSの接続情報を安全に保存
- パスワード、ユーザー名、エンドポイント等を含む
- ECSタスクから自動的に読み込み可能

### 2. アプリケーションシークレット
- APIキー、暗号化キー等のアプリケーション固有のシークレット
- 任意のキー・バリューペアで設定可能

### 3. JWTシークレット
- JWT認証用のシークレットキー
- 自動生成またはカスタム値の設定が可能

## 入力変数

| 名前 | 説明 | タイプ | デフォルト | 必須 |
|------|-------------|------|---------|:--------:|
| project_name | プロジェクト名 | `string` | n/a | yes |
| env | 環境名 | `string` | n/a | yes |
| db_username | データベースユーザー名 | `string` | n/a | yes |
| db_password | データベースパスワード | `string` | n/a | yes |
| db_host | データベースホスト | `string` | n/a | yes |
| db_port | データベースポート | `number` | `5432` | no |
| db_name | データベース名 | `string` | n/a | yes |
| db_instance_identifier | データベースインスタンス識別子 | `string` | n/a | yes |
| recovery_window_in_days | シークレット削除までの待機日数 | `number` | `7` | no |
| app_secrets | アプリケーションシークレット | `map(string)` | `{}` | no |
| create_jwt_secret | JWTシークレットを作成するか | `bool` | `false` | no |
| jwt_secret_value | JWTシークレット値（空の場合は自動生成） | `string` | `""` | no |
| tags | 追加するタグ | `map(string)` | `{}` | no |

## 出力値

| 名前 | 説明 |
|------|-------------|
| db_password_secret_arn | データベースパスワードシークレットのARN |
| db_password_secret_name | データベースパスワードシークレットの名前 |
| app_secrets_arn | アプリケーションシークレットのARN |
| app_secrets_name | アプリケーションシークレットの名前 |
| jwt_secret_arn | JWTシークレットのARN |
| jwt_secret_name | JWTシークレットの名前 |

## セキュリティ考慮事項

1. **最小権限原則**: ECSタスクには必要なシークレットのみアクセス権限を付与
2. **ローテーション**: 本番環境では定期的なシークレットローテーションを実装
3. **監査**: CloudTrailでシークレットアクセスを監査
4. **暗号化**: すべてのシークレットはKMSで暗号化

## Goアプリケーションでの使用例

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type DBConfig struct {
    Username string `json:"username"`
    Password string `json:"password"`
    Host     string `json:"host"`
    Port     int    `json:"port"`
    DBName   string `json:"dbname"`
}

func getDBConfig() (*DBConfig, error) {
    cfg, err := config.LoadDefaultConfig(context.TODO())
    if err != nil {
        return nil, err
    }
    
    client := secretsmanager.NewFromConfig(cfg)
    
    secretName := "myproject-dev-db-password"
    result, err := client.GetSecretValue(context.TODO(), &secretsmanager.GetSecretValueInput{
        SecretId: &secretName,
    })
    if err != nil {
        return nil, err
    }
    
    var dbConfig DBConfig
    err = json.Unmarshal([]byte(*result.SecretString), &dbConfig)
    if err != nil {
        return nil, err
    }
    
    return &dbConfig, nil
}
```