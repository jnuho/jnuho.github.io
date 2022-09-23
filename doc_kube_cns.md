### 2022-09-15

IAM > EKE Cluster role

- fargate-bastion 생성
	- EC2 > Amazon Linux > Key pair
	- VPC-Shared-network
	- subnet-
	- public ip 자동할당
	- SG-ssh로 내 ip에서만 접속 `ssh -i .. `
	- kubeconfig 설정하기 -> Control Plane 접근

- EKS > Cluster with created IAM role
	- endpoint : private

- ECR 생성 & image push

```sh
aws ecr create-repository --repository-name repository-httpd-junho --region ap-northeast-2

wget https://eks-fargate-test.s3.ap-northeast-2.amazonaws.com/Dockerfile &&
wget https://eks-fargate-test.s3.ap-northeast-2.amazonaws.com/pod.yaml &&
wget https://eks-fargate-test.s3.ap-northeast-2.amazonaws.com/web-ingress.yaml

docker build -t image-httpd-junho .
docker tag image-httpd-junho 170698194833.dkr.ecr.ap-northeast-2.amazonaws.com/repository-httpd-junho

aws ecr get-login-password --region ap-northeast-2 | docker login --username AWS --password-stdin 170698194833.dkr.ecr.ap-northeast-2.amazonaws.com

docker push 170698194833.dkr.ecr.ap-northeast-2.amazonaws.com/repository-httpd-junho:latest
```

- Fargate profile 생성
	- EKS > Cluster > Compute > 'fargate profile' > 'jnuho'

```sh
# namespace 생성
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


# sts-eks-cluster 클러스터 > Load Balancer 생성
# ...생략

kubectl get pods -A | grep load

vi web-ingress.yaml

vpc 로 서브넷검색 public 아이디 -> io/subnets 2개
LB가 퍼블리서브넷에 있어야 접근가능

kubectl apply -f web-ingress.yaml
kubectl get ingress -n junho
```



### 2022-09-22

```sh
# pod down, efs에 로그 등 저장가능


# EFS 생성하기
# Amazon EFS > 파일시스템 > 생성

wget https://eks-fargate-test.s3.ap-northeast-2.amazonaws.com/pv-pvc.yaml &&
wget https://eks-fargate-test.s3.ap-northeast-2.amazonaws.com/deployment.yaml &&
wget https://eks-fargate-test.s3.ap-northeast-2.amazonaws.com/ingress-deployment.yaml


sudo mount -t nfs4 -o nfsvers=4.1,rsize=1048576,wsize=1048576,hard,timeo=600,retrans=2,noresvport fs-0875850a606547cfe.efs.ap-northeast-2.amazonaws.com:/ /efs-mount-point

df -h
	fs-0875850a606547cfe.efs.ap-northeast-2.amazonaws.com:/  8.0E     0  8.0E   0% /efs-mount-point

k apply -f pv-pvc.yaml

# 볼륨 샘성
k get pv

# 볼륨 Claim 샘성

k get pvc -A

# pvc를 Application에서 사용하도록 deployment생성

vim deployment.yaml

# container의 
k apply deployment.yaml

k get pods -A
	NAMESPACE     NAME                                            READY   STATUS              RESTARTS   AGE
	junho         deploy-web-786fbf8dfd-ggpc8                     1/1     Running             0          77s


k describe pod deploy-web-786fbf8dfd-ggpc8 -n junho

# 기능동작 efs들어가서 확인
cd ~/efs-mount-point
ls -lrth


# deployment 정의 pod수만큼 유지
# pod delete 해도 개수 유지
k get deployments.app -A

# pod down되었을떄 pvc에 로그 저장기록 남음
# k delete pod deploy-web-786fbf8dfd-ggpc8

k delete pvc sts-efs-pvc-junho -n junho


# pod에 접속에서 pod에도 쌓이는지 확인
k exec -it  deploy-web-786fbf8dfd-ggpc8 -n junho -- /bin/bash


ps -ef
	UID        PID  PPID  C STIME TTY          TIME CMD
	root         1     0  0 01:30 ?        00:00:00 /bin/sh -c while true; do echo $HOSTNAME $(date -u) >> /deploy-app/deploy-out-junho.txt; sleep 5; done
	root       337     0  0 01:44 pts/0    00:00:00 /bin/bash
	root       391     1  0 01:46 ?        00:00:00 sleep 5
	root       392   337  0 01:46 pts/0    00:00:00 ps -ef


# apache 충돌로 기동 안되어 수동기동
/usr/sbin/apache2ctl start

# Replica set
# Deployment > Replicaset > pod
k get rs -A


# pod 개수 늘리기
k scale deployment deploy-web --replicas=3 -n junho


k describe pod deploy-web-786fbf8dfd-ggpc8 -n junho

# Deployment vs ReplicaSet
#	- Deployment: 추가설정가능
#		- RollingUpdateStrategy e.g. 배포설정 max 25%
#		- ReplicaSet

# LoadBalancing 설정

# Service 및 Ingerss 배포
k expose deployment deploy-web --port=7080 --target-port=7080 --name=service-deploy-web -n junho

k get svc -A
	NAMESPACE     NAME                                TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)         AGE
	default       kubernetes                          ClusterIP   172.20.0.1       <none>        443/TCP         7d2h
	hjpark        servicedeploy-web                   ClusterIP   172.20.156.136   <none>        7080/TCP        31s
	junho         service-deploy-web                  ClusterIP   172.20.223.236   <none>        7080/TCP        5s

# k 내부에서만 접근가능
k get ep -n junho

# 외부에서 접근하도록 설정 : ingress
# 퍼블리서브넷2개(외부접근허용)
k apply -f ingress-deployment.yaml

k get ingress -A
	NAMESPACE   NAME                 CLASS    HOSTS   ADDRESS                                                                      PORTS   AGE
	junho       deploy-web-ingress   <none>   *       k8s-junho-deploywe-e8535b747e-134376457.ap-northeast-2.elb.amazonaws.com     80      12s
	junho       web-ingress          <none>   *       k8s-junho-webingre-566f1e1e28-1596269189.ap-northeast-2.elb.amazonaws.com    80      6d20h


# pod 3개 중 1만 직접기동
# ALB가 health check ='healthy'되는 타겟에만 direct

# EKS Strategy 30 min
```



### 2022-09-29


