#!/bin/bash

set -e

AWS_REGION="us-east-1"
REPO_NAME="fiber-app-repo"

REPO_URL=$(aws ecr describe-repositories --repository-names $REPO_NAME --region $AWS_REGION --query "repositories[0].repositoryUri" --output text)

aws ecr get-login-password --region $AWS_REGION | docker login --username AWS --password-stdin $REPO_URL

docker build -t fiber-app:latest ../fiber-app
docker tag fiber-app:latest $REPO_URL:latest
docker push $REPO_URL:latest

cd ../infra
terraform apply -auto-approve
