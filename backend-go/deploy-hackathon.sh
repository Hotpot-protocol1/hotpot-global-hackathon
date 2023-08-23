aws ecr --profile hotpot get-login-password --region ap-southeast-1 | docker login --username AWS --password-stdin 226144105235.dkr.ecr.ap-southeast-1.amazonaws.com

docker build -t hotpot-hackathon .
docker tag hotpot-hackathon:latest 226144105235.dkr.ecr.ap-southeast-1.amazonaws.com/hotpot-hackathon:latest
docker push 226144105235.dkr.ecr.ap-southeast-1.amazonaws.com/hotpot-hackathon:latest

ecs-cli compose -f docker-compose-hackathon.yml --project-name hotpot-global service up --aws-profile hotpot --ecs-profile hotpot --cluster hotpot-hackathon --region ap-southeast-1