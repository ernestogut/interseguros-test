#!/bin/bash

set -e

AWS_REGION="us-east-1"
REPO_NAME="fiber-app-repo"

REPO_URL=$(aws ecr describe-repositories --repository-names $REPO_NAME --region $AWS_REGION --query "repositories[0].repositoryUri" --output text)

TAG=$(date +%Y%m%d%H%M%S)

aws ecr get-login-password --region $AWS_REGION | docker login --username AWS --password-stdin $REPO_URL

docker build -t fiber-app:$TAG ../fiber-app
docker tag fiber-app:$TAG $REPO_URL:$TAG
docker push $REPO_URL:$TAG

cd ../infra
terraform apply -auto-approve -var="fiber_app_image_tag=$TAG"