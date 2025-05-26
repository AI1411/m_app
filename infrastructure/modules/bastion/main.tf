# IAM Role for SSM
resource "aws_iam_role" "bastion_role" {
  name = "${var.project_name}-${var.env}-bastion-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "ec2.amazonaws.com"
        }
      }
    ]
  })

  tags = {
    Name = "${var.project_name}-${var.env}-bastion-role"
  }
}

# SSM ポリシーをアタッチ
resource "aws_iam_role_policy_attachment" "bastion_ssm" {
  role       = aws_iam_role.bastion_role.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore"
}

# インスタンスプロファイル
resource "aws_iam_instance_profile" "bastion_profile" {
  name = "${var.project_name}-${var.env}-bastion-profile"
  role = aws_iam_role.bastion_role.name
}

# Bastion Host用のキーペア
resource "aws_key_pair" "bastion" {
  key_name   = "${var.project_name}-${var.env}-bastion-key"
  public_key = var.public_key
}

# Bastion Host用のセキュリティグループ
resource "aws_security_group" "bastion" {
  name        = "${var.project_name}-${var.env}-bastion-sg"
  description = "Security group for bastion host"
  vpc_id      = var.vpc_id

  ingress {
    description = "SSH from anywhere"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = var.allowed_ssh_cidrs
  }

  egress {
    from_port = 0
    to_port   = 0
    protocol  = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "${var.project_name}-${var.env}-bastion-sg"
  }
}

# RDSセキュリティグループにBastionからのアクセスを許可
resource "aws_security_group_rule" "rds_from_bastion" {
  type                     = "ingress"
  from_port                = 5432
  to_port                  = 5432
  protocol                 = "tcp"
  source_security_group_id = aws_security_group.bastion.id
  security_group_id        = var.rds_security_group_id
}

# 最新のAmazon Linux 2 AMIを取得
data "aws_ami" "amazon_linux" {
  most_recent = true
  owners = ["amazon"]

  filter {
    name = "name"
    values = ["amzn2-ami-hvm-*-x86_64-gp2"]
  }
}

# Bastion Host EC2インスタンス
resource "aws_instance" "bastion" {
  ami                  = data.aws_ami.amazon_linux.id
  instance_type        = var.instance_type
  key_name             = aws_key_pair.bastion.key_name
  vpc_security_group_ids = [aws_security_group.bastion.id]
  subnet_id            = var.public_subnet_id
  iam_instance_profile = aws_iam_instance_profile.bastion_profile.name

  user_data = base64encode(templatefile("${path.module}/user_data.sh", {
    db_endpoint = var.db_endpoint
  }))

  tags = {
    Name = "${var.project_name}-${var.env}-bastion"
  }
}

# Elastic IP for Bastion Host
resource "aws_eip" "bastion" {
  count    = var.enable_eip ? 1 : 0
  instance = aws_instance.bastion.id
  domain   = "vpc"

  tags = {
    Name = "${var.project_name}-${var.env}-bastion-eip"
  }
}