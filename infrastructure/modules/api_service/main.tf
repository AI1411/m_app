# --- Application Load Balancer (ALB) ---
resource "aws_lb" "main" {
  name               = "${var.project_name}-alb"
  internal           = false # 外部向けALB
  load_balancer_type = "application"
  security_groups    = [var.alb_security_group_id]
  subnets            = var.public_subnet_ids # ALBはパブリックサブネットに配置

  enable_deletion_protection = false # 本番ではtrueを推奨

  tags = {
    Name = "${var.project_name}-alb"
  }
}

resource "aws_lb_target_group" "main" {
  name        = "${var.project_name}-tg"
  port        = var.container_port
  protocol    = "HTTP" # コンテナがリッスンするプロトコル
  vpc_id      = var.vpc_id
  target_type = "ip"   # Fargateの場合はip

  health_check {
    path                = var.health_check_path
    protocol            = "HTTP"
    matcher             = "200" # 正常時のHTTPステータスコード
    interval            = 30
    timeout             = 5
    healthy_threshold   = 2
    unhealthy_threshold = 2
  }

  tags = {
    Name = "${var.project_name}-tg"
  }
}

resource "aws_lb_listener" "http" {
  load_balancer_arn = aws_lb.main.arn
  port              = 80
  protocol          = "HTTP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.main.arn
  }
}

# HTTPSリスナーを追加する場合は、ACM証明書なども設定します
# resource "aws_lb_listener" "https" {
#   load_balancer_arn = aws_lb.main.arn
#   port              = 443
#   protocol          = "HTTPS"
#   ssl_policy        = "ELBSecurityPolicy-2016-08" # 推奨ポリシー
#   certificate_arn   = var.acm_certificate_arn # ACMのARN

#   default_action {
#     type             = "forward"
#     target_group_arn = aws_lb_target_group.main.arn
#   }
# }

# --- ECS Cluster ---
resource "aws_ecs_cluster" "main" {
  name = "${var.project_name}-ecs-cluster"

  tags = {
    Name = "${var.project_name}-ecs-cluster"
  }
}

# --- ECS Task Definition ---
resource "aws_ecs_task_definition" "api" {
  family                   = "${var.project_name}-api-task"
  network_mode             = "awsvpc" # Fargateにはawsvpcが必要
  requires_compatibilities = ["FARGATE"]
  cpu                      = var.cpu
  memory                   = var.memory
  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn # ECRからイメージをプルする権限など

  # タスクロール (APIがAWSサービスにアクセスする場合に設定)
  # task_role_arn            = aws_iam_role.ecs_task_role.arn

  container_definitions = jsonencode([
    {
      name      = "${var.project_name}-api-container"
      image     = var.container_image
      cpu       = var.cpu
      memory    = var.memory
      essential = true
      portMappings = [
        {
          containerPort = var.container_port
          hostPort      = var.container_port # awsvpcモードではhostPortとcontainerPortは同じ
        }
      ]
      # 環境変数やシークレットはここで設定
      # environment = [
      #   { name = "DB_HOST", value = var.db_host },
      #   { name = "DB_PORT", value = tostring(var.db_port) },
      #   { name = "DB_USER", value = var.db_user },
      #   { name = "DB_NAME", value = var.db_name },
      # ]
      # secrets = [
      #   { name = "DB_PASSWORD", valueFrom = var.db_password_secret_arn } # Secrets ManagerのARN
      # ]
      logConfiguration = { # CloudWatch Logsへのログ出力設定
        logDriver = "awslogs"
        options = {
          "awslogs-group"         = aws_cloudwatch_log_group.api.name
          "awslogs-region"        = var.aws_region
          "awslogs-stream-prefix" = "ecs"
        }
      }
    }
  ])

  tags = {
    Name = "${var.project_name}-api-task"
  }
}

# --- ECS Service ---
resource "aws_ecs_service" "api" {
  name            = "${var.project_name}-api-service"
  cluster         = aws_ecs_cluster.main.id
  task_definition = aws_ecs_task_definition.api.arn
  desired_count   = var.desired_count
  launch_type     = "FARGATE"

  network_configuration {
    subnets          = var.private_subnet_ids # ECSタスクはプライベートサブネットに配置
    security_groups  = [var.ecs_task_security_group_id]
    assign_public_ip = false # FargateタスクにはパブリックIPを割り当てない (ALB経由アクセス)
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.main.arn
    container_name   = "${var.project_name}-api-container"
    container_port   = var.container_port
  }

  depends_on = [aws_lb_listener.http] # ALBリスナーが作成されてからサービスを開始

  # サービスディスカバリやデプロイ戦略なども設定可能
  # deployment_controller {
  #   type = "ECS"
  # }

  tags = {
    Name = "${var.project_name}-api-service"
  }
}

# --- IAM Role for ECS Task Execution ---
resource "aws_iam_role" "ecs_task_execution_role" {
  name = "${var.project_name}-ecs-task-execution-role"
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "ecs-tasks.amazonaws.com"
        }
      }
    ]
  })

  tags = {
    Name = "${var.project_name}-ecs-task-execution-role"
  }
}

resource "aws_iam_role_policy_attachment" "ecs_task_execution_role_policy" {
  role       = aws_iam_role.ecs_task_execution_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

# --- CloudWatch Log Group ---
resource "aws_cloudwatch_log_group" "api" {
  name = "/ecs/${var.project_name}-api" # ECSタスク定義のログ設定と合わせる

  tags = {
    Name = "${var.project_name}-api-logs"
  }
}