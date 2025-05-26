#!/bin/bash
yum update -y

# PostgreSQLクライアントのインストール
yum install -y postgresql

# AWS CLIのインストール（既にインストール済みだが念のため）
yum install -y aws-cli

# psqlコマンドのテスト用スクリプトを作成
cat > /home/ec2-user/connect_to_rds.sh << 'EOF'
#!/bin/bash
echo "RDS Aurora PostgreSQL Connection Helper"
echo "Database Endpoint: ${db_endpoint}"
echo ""
echo "To connect to RDS, use the following command:"
echo "psql -h ${db_endpoint} -U m_app_admin -d m_app_dev"
echo ""
echo "Or use this script with password:"
echo "PGPASSWORD=your_password psql -h ${db_endpoint} -U m_app_admin -d m_app_dev"
EOF

chmod +x /home/ec2-user/connect_to_rds.sh
chown ec2-user:ec2-user /home/ec2-user/connect_to_rds.sh

# Secrets Managerから認証情報を取得するスクリプト
cat > /home/ec2-user/get_db_credentials.sh << 'EOF'
#!/bin/bash
echo "Getting database credentials from Secrets Manager..."
aws secretsmanager get-secret-value \
    --secret-id "m-app-dev-db-password" \
    --region "ap-northeast-1" \
    --query 'SecretString' \
    --output text | jq '.'
EOF

chmod +x /home/ec2-user/get_db_credentials.sh
chown ec2-user:ec2-user /home/ec2-user/get_db_credentials.sh

# jqのインストール（JSON parsing用）
yum install -y jq

echo "Bastion host setup completed" > /var/log/user-data.log