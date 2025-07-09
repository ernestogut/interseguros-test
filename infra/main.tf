resource "aws_ecr_repository" "fiber" {
  name = "fiber-app-repo"
}

resource "aws_ecr_repository" "express" {
  name = "express-app-repo"
}

resource "aws_ecs_cluster" "main" {
  name = "interseguros-cluster"
}

data "aws_secretsmanager_secret" "jwt_secret" {
  name = "jwt-secret"
}

# ALB
resource "aws_lb" "app" {
  name               = "interseguros-alb"
  load_balancer_type = "application"
  subnets = [
    aws_subnet.public_a.id,
    aws_subnet.public_b.id,
  ]
  security_groups = [aws_security_group.ecs_sg.id]
}

# Target Group para Fiber
resource "aws_lb_target_group" "fiber_tg" {
  name        = "fiber-tg"
  port        = 8080
  protocol    = "HTTP"
  vpc_id      = aws_vpc.main.id
  target_type = "ip"
  health_check {
    path                = "/fiber/health"
    interval            = 30
    timeout             = 5
    healthy_threshold   = 2
    unhealthy_threshold = 2
    matcher             = "200-399"
  }
}

# Target Group para Express
resource "aws_lb_target_group" "express_tg" {
  name        = "express-tg"
  port        = 3000
  protocol    = "HTTP"
  vpc_id      = aws_vpc.main.id
  target_type = "ip"
  health_check {
    path                = "/express/health"
    interval            = 30
    timeout             = 5
    healthy_threshold   = 2
    unhealthy_threshold = 2
    matcher             = "200-399"
  }
}

# Listener HTTP (puerto 80)
resource "aws_lb_listener" "http" {
  load_balancer_arn = aws_lb.app.arn
  port              = "80"
  protocol          = "HTTP"

  default_action {
    type = "fixed-response"
    fixed_response {
      content_type = "text/plain"
      message_body = "Default response"
      status_code  = "404"
    }
  }
}

# Listener Rule para Fiber (ej: /fiber/*)
resource "aws_lb_listener_rule" "fiber_rule" {
  listener_arn = aws_lb_listener.http.arn
  priority     = 10

  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.fiber_tg.arn
  }

  condition {
    path_pattern {
      values = ["/fiber/*"]
    }
  }
}

# Listener Rule para Express (ej: /express/*)
resource "aws_lb_listener_rule" "express_rule" {
  listener_arn = aws_lb_listener.http.arn
  priority     = 20

  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.express_tg.arn
  }

  condition {
    path_pattern {
      values = ["/express/*"]
    }
  }
}

resource "aws_ecs_task_definition" "fiber_task" {
  family                   = "fiber-task"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = "256"
  memory                   = "512"
  task_role_arn            = aws_iam_role.ecs_task_role.arn
  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn

  container_definitions = jsonencode([{
    name      = "fiber-container"
    image     = "${aws_ecr_repository.fiber.repository_url}:${var.fiber_app_image_tag}"
    essential = true
    portMappings = [{
      containerPort = 8080
      hostPort      = 8080
    }]
    environment = [
      {
        name  = "PORT"
        value = "8080"
      },
      {
        name  = "NODE_APP_URL"
        value = "http://${aws_lb.app.dns_name}/express"
      }
    ],
    secrets = [
      {
        name      = "JWT_SECRET"
        valueFrom = data.aws_secretsmanager_secret.jwt_secret.arn
      },
    ]
  }])
}

resource "aws_ecs_task_definition" "express_task" {
  family                   = "express-task"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = "256"
  memory                   = "512"
  task_role_arn            = aws_iam_role.ecs_task_role.arn
  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn

  container_definitions = jsonencode([{
    name      = "express-container"
    image     = "${aws_ecr_repository.express.repository_url}:${var.express_app_image_tag}"
    essential = true
    portMappings = [{
      containerPort = 3000
      hostPort      = 3000
    }]
    environment = [
      {
        name  = "PORT"
        value = "3000"
      }
    ],
    secrets = [
      {
        name      = "JWT_SECRET"
        valueFrom = data.aws_secretsmanager_secret.jwt_secret.arn
      },
    ]
  }])
}

resource "aws_ecs_service" "fiber_service" {
  name                   = "fiber-service"
  cluster                = aws_ecs_cluster.main.id
  task_definition        = aws_ecs_task_definition.fiber_task.arn
  desired_count          = 1
  launch_type            = "FARGATE"
  enable_execute_command = true

  network_configuration {
    subnets = [
      aws_subnet.public_a.id,
      aws_subnet.public_b.id
    ]
    security_groups  = [aws_security_group.ecs_sg.id]
    assign_public_ip = true
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.fiber_tg.arn
    container_name   = "fiber-container"
    container_port   = 8080
  }

  depends_on = [aws_lb_listener_rule.fiber_rule]

  force_new_deployment = true
}

resource "aws_ecs_service" "express_service" {
  name                   = "express-service"
  cluster                = aws_ecs_cluster.main.id
  task_definition        = aws_ecs_task_definition.express_task.arn
  desired_count          = 1
  launch_type            = "FARGATE"
  enable_execute_command = true

  network_configuration {
    subnets = [
      aws_subnet.public_a.id,
      aws_subnet.public_b.id
    ]
    security_groups  = [aws_security_group.ecs_sg.id]
    assign_public_ip = true
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.express_tg.arn
    container_name   = "express-container"
    container_port   = 3000
  }

  depends_on = [aws_lb_listener_rule.express_rule]

  force_new_deployment = true
}
