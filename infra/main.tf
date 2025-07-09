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

resource "aws_ecs_task_definition" "fiber_task" {
  family                   = "fiber-task"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = "256"
  memory                   = "512"
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
    ],
    secrets = [
      {
        name      = "NODE_APP_URL"
        valueFrom = data.aws_secretsmanager_secret.jwt_secret.arn
      },
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
  name            = "fiber-service"
  cluster         = aws_ecs_cluster.main.id
  task_definition = aws_ecs_task_definition.fiber_task.arn
  desired_count   = 1
  launch_type     = "FARGATE"

  network_configuration {
    subnets          = [aws_subnet.public_a.id]
    security_groups  = [aws_security_group.ecs_sg.id]
    assign_public_ip = true
  }

  force_new_deployment = true
}

resource "aws_ecs_service" "express_service" {
  name            = "express-service"
  cluster         = aws_ecs_cluster.main.id
  task_definition = aws_ecs_task_definition.express_task.arn
  desired_count   = 1
  launch_type     = "FARGATE"

  network_configuration {
    subnets          = [aws_subnet.public_a.id]
    security_groups  = [aws_security_group.ecs_sg.id]
    assign_public_ip = true
  }

  force_new_deployment = true
}
