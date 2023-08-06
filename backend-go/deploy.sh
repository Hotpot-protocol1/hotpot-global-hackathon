# Get token for ECR
aws ecr --profile hotpot get-login-password --region ap-southeast-1 | docker login --username AWS --password-stdin 226144105235.dkr.ecr.ap-southeast-1.amazonaws.com

# Build docker image for backend, tag and push it on ECR
docker build -t hotpot-global .
docker tag hotpot-global:latest 226144105235.dkr.ecr.ap-southeast-1.amazonaws.com/hotpot-global:latest
docker push 226144105235.dkr.ecr.ap-southeast-1.amazonaws.com/hotpot-global:latest

ecs-cli compose --project-name hotpot-global service up --aws-profile hotpot --ecs-profile hotpot --cluster Hotpot --region ap-southeast-1