

- EC2 & key pair
  - 퍼블릭키는 EC2 인스턴스에 설치됨
  - 개인키는 접속주체 개인 저장

```
# Permissions 0644 for '*.pem' are too open!
# 개인키에 최소권한 설정 필요
chmod 600 {my.pem}
ssh -i {my.pem} ubuntu@{public_ip}
```

- AWS CLI
  - access key 필요
  - secret access key 필요
  - 프로필>내보안자격증명>액세스키 생성

- [CLI 설치 (우분투)](https://docs.aws.amazon.com/ko_kr/cli/latest/userguide/install-cliv2-linux.html)
```sh
cat /etc/lsb-release
sudo apt update
sudo apt install unzip build-essential curl

# 최신버전
# curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
# 특정버전
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64-2.2.16.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install
```

- WSL 사용설정 (윈도우)
```shell
# powershell as Admin
dism.exe /online /enable-feature /featurename:Microsoft-Windows-Subsystem-Linux /all /norestart
dism.exe /online /enable-feature /featurename:VirtualMachinePlatform /all /norestart

# 재부팅 후 실행
wsl

# 우분투 설치 (Microsoft Store)

# 버전 확인
wsl -l -v

# 커널 업데이트 파일 다운로드 및 설치
https://wslstorestorage.blob.core.windows.net/wslblob/wsl_update_x64.msi

# 버전 업데이트
wsl --set-version Ubuntu 2
# 또는
wsl --set-version Ubuntu-20.04 2


# 업데이트된 버전 상태확인
wsl -l -v
# stopped된 wsl 재실행
wsl

# 디폴트 버전 2 설정
wsl --set-default-version 2

```


- AWS CLI자격증명 설정 우선순위
  1. CLI명령어 옵션 (실무)
  2. 환경변수 (실무)
  3. CLI 자격증명 파일 - `~/.aws/credentials`
  4. CLI 설정 파일 - `~/.aws/config` (실무)
  5. 컨테이너 자격증명 (ECS의 경우)
  6. 인스턴스 프로파일 자격증명 (EC2의 경우) (실무)

```sh
# 1. CLI 명령어
# --profile

# 2. 환경변수
AWS_ACCESS_KEY_ID
AWS_SECRET_ACCESS_KEY
AWS_PROFILE

# 6. EC2생성시 IAM 역할 부여

# 4. CLI 설정파일
# 액세스키 다운로드 후
cat mycredential.csv
vim ~/.aws/config
[default]
aws_access_key_id=AKI...K
aws_secret_access_key_id=k0r...M8

# 테스트
# 계정정보 확인
aws sts get-caller-identity
aws ec2 describe-key-pairs

# 리전설정 필요
# 매번 명령어에 --region 생략하려면 CLI 설정파일에 region 추가
aws ec2 describe-key-pairs --region ap-northeast-2

[default]
region=ap-northeast-2
aws_access_key_id=AKI...K
aws_secret_access_key_id=k0r...M8


# 리전설정 없이 키페어 조회 가능
aws ec2 describe-key-pairs

aws ec2 describe-key-pairs --output yaml
aws ec2 describe-key-pairs --output table

[default]
output=json
region=ap-northeast-2
aws_access_key_id=AKI...K
aws_secret_access_key_id=k0r...M8

# 아웃풋 옵션없이 키페어 조회 가능
aws ec2 describe-key-pairs


# AWS CLI 여러 사용자 프로필 등록 가능 :
# 프로파일 AWS_PROFILE 환경 변수
# 또는 --profile
# 또는 config파일
[default]
region=ap-northeast-2
output=yaml
aws_access_key_id=AKI...K
aws_secret_access_key_id=k0r...M8

[profile eu-west-1]
region=eu-west-1
aws_access_key_id=AKI...K
aws_secret_access_key_id=k0r...M8

[profile ap-northeast-1]
region=ap-northeast-1
aws_access_key_id=AKI...K
aws_secret_access_key_id=k0r...M8

# A.여러 AWS 계정운영
# B.동일 계정 내 여러 리전 운열
# C.동일 계정 내 여러 IAM 역할 전환 수행
# D.AWS SSO 조직 내 역할 수행
# 디폴트 프로파일의 리전
aws configure get region
ap-northeast-2

# eu-west-1 프로파일의 리전
aws configure get region --profile eu-west-1
eu-west-1

# 환경변수로 디폴트 리전 설정
export AWS_PROFILE=ap-northeast-1
aws configure get region
ap-northeast-1
```


- AWS CLI 사용법

```sh
aws <command> <subcommand> [options and parameters]
aws --version
aws help
aws <command> help
aws <command> <subcommand> help
aws ec2 describe-availability-zones help
```


- 디버그 옵션 활성화

```sh
# 자격 증명정보 확인시 디버그모드 활성화
aws sts get-caller-identity --debug
```


- Auto Scaling
  - increase ec2 machine count
- Elastic Load Balancing
  - distribute network

- IPV4
  - 2^32
  - 서브넷을 통해 추가 세분화
  - A클래스 큰 네트워크 - 국가수준
  - 211.11.124.2

```sh
# IGW -> VPC
aws ec2 attach-internet-gateway --vpc-id "vpc-025bb54e5d273c193" --internet-gateway-id "igw-09fe226851efbc823" --region ap-northeast-2
```



서비스
- EC2 사양과 크기 조절할수 있는 컴퓨팅 서비스
- Auto Scaling 서버의 특정 조건에 따라 서버를 추가/삭제할 수 있게 하는 서비스
  - 서버를 여러대 부하 분산
- Lightsail 가상화 프라이빗 서버
- Workspaces 사내PC가상화
- Route 53 DNS 웹서비스
- VPC 가상네트워크를 클라우드에 구성
- Direct Connect 온프레미스 인프라와 AWS연결 하는 네트워킹 서비스
- ELB 부하분산(로드 밸런싱) 서비스
  - 네트워킹 방식으로 부하 분산
- S3: 여러가지 파일을 형식 구애받지 않고 저장
- Dynamo DB 가상 NoSQL 데이터베이스
- RDS 가상 SQL 데이터베이스
- ElastiCache In-memory 기반의 cache서비스 (빠른속도 필요로 하는 서비스)
- Redshift 데이터 분석에 특화된 스토리지 시스템
- EMR 대량 데이터 효율적 가공
- Sagemaker 머신러닝 & 데이터분석을 위한 클라우드 환경제공


AWS 기초와 VPC
- 네트워킹 기초

- AMI
  - EBS 스냅샷 생성하여 EC2 -> AMI 변환 가능


### Cloud Platform Models

- Iaas (Infrastructure as a Service)
  - EC2, EBS, ELB
  - look and feel you'd get from managing physical resources
  - responsiable for consequences of your configuration

- Paas (Platform as a Service)
  - Elastic Beanstalk, ECS (Elastice Container Service)
  - simplifies the process of building an application by hiding the complexity of the infrastructure
- Saas (Software as a Service)
  - offer services meant to be accessed by end users.
  - Simple Email Service, Amazon WorkSpaces, Gmail, Outlook



- MSA
  - [Django 프로젝트](https://docs.djangoproject.com/en/4.0/intro/tutorial01/)
    - `settings.py`
    - `urls.py`


```shell
python --version

python -m pip install django
python -m django --version
django-admin startproject mysite
cd mysite/
# http://localhost:8000/
# python manage.py runserver 8080
# python manage.py runserver 0:8080
python manage.py runserver

# polls 앱시작
python manage.py startapp polls
# polls/urls.py 추가
# mysite/urls.py 수정

python manage.py runserver
# http://127.0.0.1:8000/polls/



# mysite > settings.py > INSTALLED_APPS PollsConfig 추가 후
# DB 스키마들을 정의 하는 migration 스크립트 0001_initial.py 생성됨
python manage.py makemigrations polls

# 0001_initial.py의 스키마 대로 db.sqllite3에 DB 구성
python manage.py migrate

# DB확인하기
python manage.py shell

```



- Part 5. 도커와 쿠버네티스를 이용한 서비스 운영
  - docker, docker-compose, kubernetes

- 도커
  - 우분투환경 t3.small (쿠버네티스 minikube 사용 위한 최소 사양 - NOT FREE-TIER!)
  - 설치스크립트 install-docker.sh
    - https://github.com/tedilabs/fastcampus-devops/blob/main/3-docker-kubernetes/env/ubuntu/install-docker.sh
    - install-docker-compose.sh


```shell
# Use Docker without root
sudo usermod -aG docker $DOCKER_USER

# github내용 붙여넣고 Ctrl+C (^M캐릭터 등 문제 없애기 위해)
cat > install-docker.sh
chmod u+x install-docker.sh
./install-docker.sh
# 도커가 그룹목록에 없음. 터미널 재실행 후 반영 확인
id
docker ps

cat > install-docker-compose.sh
chmod u+x install-docker-compose.sh
./install-docker-compose.sh

```


- kubectl
  - 쿠버네티스의 API 서버(클러스터)와 통신하여 사용자 명령을 전달하는 CLI 도구
    - https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/#install-using-native-package-management

```shell
cat > install-kubectl.sh
chmod u+x install-kubectl.sh
./install-kubectl.sh

```


- kustomize (또는 Helm사용)
  - https://kubectl.docs.kubernetes.io/installation/kustomize/binaries/


```shell
cat > install-kustomize.sh
chmod u+x install-kustomize.sh
./install-kubectl.sh

```

- minikube (또는 Helm사용)
  - https://minikube.sigs.k8s.io/docs/start/

```shell
cat > install-kustomize.sh
chmod u+x install-kustomize.sh
./install-kubectl.sh

minikube start --driver=docker
cat ~/.kube/config
# 컨텍스트: 클러스터, 인증사용자 정보로 인증,접속 한다는 정보
# minikube 컨텍스트를 통해서 kubectl에 통신
```

```shell
# 노드의이름: minikube, 역할: control-plane 과 master 역할 수행
kubectl get nodes
kubectl cluster-info

# 쿠버네티스 클러스터 커멘드
# status, stop, delete, pause, unpause
minikube status


# 쿠버네티스 클러스터에 addon 설정
# : dashboard, helm-tiller, ingress, istio, metallb, metrics-server
# minikube addons enable ingress
minikube addons list

# 쿠버네티스 클러스터에 있는 노드 컨테이너에 접근
minikube ssh

# 로컬에 설치된 kubectl 버전과 , 쿠버네티스 클러스터 버전일치 하지 않을 수 있음
# 클러스터버전과 동일한 kubectl 버전을 사용하는 커멘드
minikube kubectl version
```

- 테라폼 코드 이용하여 미리 준비된 AWS 실습환경 구성
  - https://github.com/tedilabs/fastcampus-devops/tree/main/3-docker-kubernetes/env/terraform-aws-ubuntu
 
## 도커

- 이미지: build 또는 pull (from 이미지저장소 Registry)
  - 퍼플릭: dockerhub
  - 프라이빗: AWS ECR
- 컨테이너 : 이미지 실행 프로세스
  - 라이프사이클: created, deleted, running, stopped, paused

```shell
docker ps -a

docker run -d nginx
docker run -it ubuntu:focal 
docker run ubuntu:focal id

docker run -p {호스트포트}:{컨테이너포트} -d nginx
docker run -p 9999:80 nginx -d
curl localhost:{호스트포트}

# 컨테이너 생성
docker create [image]

# 컨테이너 시작
docker start [container]

# 컨테이너 생성 및 시작시작
docker run [image]

docker run \
  -i \                          # 호스트의 표준 입력을 컨테이너와 연결 (interactive)
  -t \                          # TTY 할당
  --rm \                        # 컨테이너 실행 종료 후 자동 삭제
  -d \                          # 백그라운드 모드로 실행 (detached)
  --name hello-world \          # 컨테이너 이름 지정
  -p 80:80 \                    # 호스트 - 컨테이너 간 포트 바인딩
  -v /opt/example:/example \    # 호스트 - 컨테이너 간 볼륨 바인딩
  fastcampus/hello-world:latest \   # 실행할 이미지
  my-command                        # 컨테이너 내에서 실행할 명령어
  
  
# 컨테이너 중단 및 빠져나오려면 exit
# 컨테이너 중단X 빠져나오려면 CTRL+P,Q
docker run -p 80:80 -d nginx
curl localhost:80

# 기본 COMMAND = id 세팅
docker run ubuntu:focal id

# 컨테이너가 실행 종료 후 삭제 되도록 하는 옵션 --rm
docker run --rm ubuntu:focal id

# 컨테이너 inspect
docker inspect [container]

# 일시중지 / 재개
docker pause [container]
docker unpause [container]

# SIGTERM 시그널로 종료
# gracefully shutdown 안전하게 종료
docker stop [container]

# SIGKILL 시그널로 종료
# 로그남지 않고 강제종료
docker stop [container]

# 모든 컨테이너 종료
docker stop $(docker ps -a -q)

# 컨테이너 삭제 (실행중 제외)
docker rm [container]

# 컨테이너 강제종료 후 SIGKILL 삭제
docker rm -f [container]

# 컨테이너 실행 종료 후 자동삭제
docker run --rm ...

# 중지된 모든 컨테이너 삭제
docker container prune
```


- 엔트리포인트와 커멘드
  - 엔트리포인트 : 컨테이너 실행 시 고정적으로 실행되는 스크립트/명령어
  - 커멘드 : 

