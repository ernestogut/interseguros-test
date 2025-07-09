output "fiber_ecr_repo_url" {
  value = aws_ecr_repository.fiber.repository_url
}

output "express_ecr_repo_url" {
  value = aws_ecr_repository.express.repository_url
}
