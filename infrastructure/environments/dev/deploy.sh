#!/bin/bash

# デプロイスクリプト for m-app

set -e

# 設定
AWS_REGION="ap-northeast-1"
ECR_REPOSITORY_NAME="m-app-dev-api"
ECS_CLUSTER_NAME="dev-m-app-ecs-cluster"
ECS_SERVICE_NAME="m-app-api-service"

echo "🚀 Starting deployment process..."

# AWS Account IDを取得
AWS_ACCOUNT_ID=$(aws sts get-caller-identity --query Account --output text)
ECR_REGISTRY="${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com"
IMAGE_URI="${ECR_REGISTRY}/${ECR_REPOSITORY_NAME}:latest"

echo "📦 Building Docker image..."

# Dockerイメージをビルド
docker build -t ${ECR_REPOSITORY_NAME}:latest .

# ECRにログイン
echo "🔐 Logging in to ECR..."
aws ecr get-login-password --region ${AWS_REGION} | docker login --username AWS --password-stdin ${ECR_REGISTRY}

# イメージにタグ付け
echo "🏷️  Tagging image..."
docker tag ${ECR_REPOSITORY_NAME}:latest ${IMAGE_URI}

# ECRにプッシュ
echo "📤 Pushing image to ECR..."
docker push ${IMAGE_URI}

# ECSサービスを更新
echo "🔄 Updating ECS service..."
aws ecs update-service \
    --cluster ${ECS_CLUSTER_NAME} \
    --service ${ECS_SERVICE_NAME} \
    --force-new-deployment \
    --region ${AWS_REGION}

echo "⏳ Waiting for deployment to complete..."
aws ecs wait services-stable \
    --cluster ${ECS_CLUSTER_NAME} \
    --services ${ECS_SERVICE_NAME} \
    --region ${AWS_REGION}

# ALBのDNS名を取得
ALB_DNS=$(aws elbv2 describe-load-balancers \
    --names "m-app-alb" \
    --query 'LoadBalancers[0].DNSName' \
    --output text \
    --region ${AWS_REGION})

echo "✅ Deployment completed successfully!"
echo "🌐 Application URL: http://${ALB_DNS}"
echo "🏥 Health check: http://${ALB_DNS}/health"

# ヘルスチェック実行
echo "🔍 Performing health check..."
sleep 30  # サービス起動待ち
curl -f "http://${ALB_DNS}/health" || echo "⚠️  Health check failed - service may still be starting"

echo "🎉 Deployment process completed!"