## 入力変数

| 名前 | 説明 | 型 | デフォルト | 必須 |
|------|-------------|------|---------|:-----:|
| bucket_name | Terraformステート用S3バケットの名前 | string | - | はい |
| dynamodb_table_name | ステートロック用DynamoDBテーブルの名前 | string | - | はい |
| tags | リソースに追加するタグ | map(string) | {} | いいえ |

## 出力値

| 名前 | 説明 |
|------|-------------|
| bucket_id | 作成されたS3バケットのID |
| bucket_arn | 作成されたS3バケットのARN |
| dynamodb_table_id | 作成されたDynamoDBテーブルのID |
| dynamodb_table_arn | 作成されたDynamoDBテーブルのARN |

## Terraformバックエンド設定例

このモジュールで作成したリソースをTerraformバックエンドとして使用する例：