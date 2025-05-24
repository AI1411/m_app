variable "vpc_id" {
  description = "ID of the VPC where to create security groups"
  type        = string
}

variable "project_name" {
  description = "Name of the project"
  type        = string
}

variable "alb_sg_ingress_cidrs" {
  description = "CIDR blocks allowed to access the ALB (e.g., ['0.0.0.0/0'] for public)"
  type = list(string)
  default = ["0.0.0.0/0"]
}

# --- ALB Security Group ---
resource "aws_security_group" "alb" {
  name        = "${var.project_name}-alb-sg"
  description = "Security group for Application Load Balancer"
  vpc_id      = var.vpc_id

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = var.alb_sg_ingress_cidrs
  }

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = var.alb_sg_ingress_cidrs
  }

  egress {
    from_port = 0
    to_port   = 0
    protocol = "-1" # Allow all outbound traffic
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "${var.project_name}-alb-sg"
  }
}

# --- API Service Security Group (e.g., for ECS Tasks) ---
resource "aws_security_group" "api_service" {
  name        = "${var.project_name}-api-service-sg"
  description = "Security group for API service (e.g., ECS tasks)"
  vpc_id      = var.vpc_id

  # Ingress: Allow traffic from ALB SG on the container port
  # The port will be passed as a variable or defined in the api_service module
  # For simplicity, we define a variable here, but it's better to get the port from the api_service module's variables
  # ingress {
  #   from_port       = var.api_service_container_port # This should ideally come from api_service module's input
  #   to_port         = var.api_service_container_port
  #   protocol        = "tcp"
  #   security_groups = [aws_security_group.alb.id] # Allow from ALB
  # }

  egress {
    from_port = 0
    to_port   = 0
    protocol  = "-1"
    cidr_blocks = ["0.0.0.0/0"] # Allow outbound to NAT Gateway (for external APIs, ECR, etc.)
  }

  tags = {
    Name = "${var.project_name}-api-service-sg"
  }
}

# --- RDS Security Group ---
resource "aws_security_group" "rds" {
  name        = "${var.project_name}-rds-sg"
  description = "Security group for RDS Aurora PostgreSQL"
  vpc_id      = var.vpc_id

  # Ingress: Allow traffic from API Service SG on PostgreSQL port
  # ingress {
  #   from_port       = 5432 # PostgreSQL port
  #   to_port         = 5432
  #   protocol        = "tcp"
  #   security_groups = [aws_security_group.api_service.id] # Allow from API service
  # }
  # Egress: Not strictly necessary for RDS if all outbound is allowed by default or to specific targets

  tags = {
    Name = "${var.project_name}-rds-sg"
  }
}