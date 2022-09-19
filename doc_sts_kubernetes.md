







IAM > EKE Cluster role


fargate-bastion
	EC2 > Amazon Linux > Key pair
	VPC-Shared-network
	subnet-
	public ip 자동할당
	SG-ssh로 내 ip에서만 접속



	aws configure
	ssh -i ...



EKS>cluster , iam role
cluster endpoint(private)



wget https://eks-fargate-test.s3.ap-northeast-2.amazonaws.com/Dockerfile &&
wget https://eks-fargate-test.s3.ap-northeast-2.amazonaws.com/pod.yaml &&
wget https://eks-fargate-test.s3.ap-northeast-2.amazonaws.com/web-ingress.yaml

aws ecr create-repository --repository-name repository-httpd-junho --region ap-northeast-2

docker build -t image-httpd-junho .
docker tag image-httpd-junho 170698194833.dkr.ecr.ap-northeast-2.amazonaws.com/repository-httpd-junho

aws ecr get-login-password --region ap-northeast-2 | docker login --username AWS --password-stdin 170698194833.dkr.ecr.ap-northeast-2.amazonaws.com


docker push 170698194833.dkr.ecr.ap-northeast-2.amazonaws.com/repository-httpd-junho:latest

fargate profile 생성
	junho



kubectl create namespace junho
kubectl get ns

vi pod.yaml
kubectl apply -f pod.yaml
kubectl get pod -n junho

k get pod -A
k expose pod web --name=web-svc -n junho

k get service -n junho

	NAME      TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
	web-svc   ClusterIP   172.20.111.54   <none>        7080/TCP   7s


sts-eks-cluster 클러스터
Load Balancer 생성

kubectl get pods -A | grep load



vi web-ingress.yaml

vpc 로 서브넷검색 public 아이디 -> io/subnets 2개
LB가 퍼블리서브넷에 있어야 접근가능


kubectl apply -f web-ingress.yaml
kubectl get ingress -n junho
