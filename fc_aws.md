

- EC2 & key pair
  - 퍼블릭키는 EC2 인스턴스에 설치됨 - 개인키는 접속주체 개인 저장

```
# Permissions 0644 for '*.pem' are too open!
# 개인키에 최소권한 설정 필요
chmod 600 {my.pem}
ssh -i {my.pem} ubuntu@{public_ip}
```

- AWS CLI
  - Access key, Secret Access key 발급 필요
  - 프로필 > 보안자격증명 > 액세스키 생성

- [CLI 설치 (우분투)](https://docs.aws.amazon.com/ko_kr/cli/latest/userguide/install-cliv2-linux.html)

```shell
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

# 우분투 설치 (Microsoft Store)

# 재부팅 후 실행
wsl


# 버전 확인
wsl -l -v

# 커널 업데이트 파일 다운로드 및 설치
https://wslstorestorage.blob.core.windows.net/wslblob/wsl_update_x64.msi

# 버전 업데이트
# 또는 wsl --set-version Ubuntu-20.04 2
wsl --set-version Ubuntu 2


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

```shell
# 1. CLI 명령어
# --profile

# 2. 환경변수
AWS_ACCESS_KEY_ID
AWS_SECRET_ACCESS_KEY
AWS_PROFILE

# 6. EC2생성시 IAM 역할 부여

# 4. CLI 설정파일
# 액세스키 다운로드 후
cat mycredential.csv >> temp.csv
cat temp.csv

aws configure
  AWS Access Key ID : 
  AWS Secret Access Key : 
  Default region name : ap-northeast-2
  Default output format : json

cat ~/.aws/config
[default]
region = ap-northeast-2
output = json

cat ~/.aws/credentials
[default]
aws_access_key_id=...
aws_secret_access_key_id=...

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

aws ec2 describe-key-pairs --output json
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

```shell
aws <command> <subcommand> [options and parameters]
aws --version
aws help
aws <command> help
aws <command> <subcommand> help
aws ec2 describe-availability-zones help
```


- 디버그 옵션 활성화

```shell
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

```shell
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
    - install-docker.sh
    - install-docker-compose.sh
    - install-kubectl.sh
    - install-kustomize.sh
    - install-minikube.sh
    - update-apt.sh


```shell
# Use Docker without root
sudo usermod -aG docker $DOCKER_USER

# github내용 붙여넣고 Ctrl+C (^M캐릭터 등 문제 없애기 위해)
cat > install-docker.sh
#!/usr/bin/env bash
## INFO: https://docs.docker.com/engine/install/ubuntu/

set -euf -o pipefail

DOCKER_USER=ubuntu

# Install dependencies
sudo apt-get update && sudo apt-get install -y \
  apt-transport-https \
  ca-certificates \
  curl \
  gnupg \
  lsb-release

# Add Docker’s official GPG key
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --yes --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

# Set up the stable repository
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

# Install Docker CE
sudo apt-get update && sudo apt-get install -y docker-ce docker-ce-cli containerd.io

# Use Docker without root
sudo usermod -aG docker $DOCKER_USER

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
# cluster : 클러스터 정보
# context : 클러스터 인증리스트; 어떤 클러스터와 인증을 할지), 유저
#           클러스터, 인증사용자 정보로 인증,접속 한다는 정보
# user : 인증사용자 정보
# minikube 컨텍스트를 통해서 kubectl에 통신
```

```shell
# 노드의이름: minikube, 역할: control-plane 과 master 역할 수행
kubectl get nodes
kubectl cluster-info
# Details on status, CPU, memory, system information, containers
kubectl describe node minikube

# 쿠버네티스 클러스터 커멘드
# status, stop, delete, pause, unpause
minikube status


# 쿠버네티스 클러스터에 addon 설정
# : dashboard, helm-tiller, ingress, istio, metallb, metrics-server
# minikube addons enable ingress
# minikube addons enable dashboard
minikube addons list

# 쿠버네티스 클러스터에 있는 노드 컨테이너에 접근
minikube ssh

# 로컬에 설치된 kubectl 버전과 , 쿠버네티스 클러스터 버전일치 하지 않을 수 있음
# 클러스터버전과 동일한 kubectl 버전을 사용하는 minikube 커멘드
# kubectl get nodes # 쿠버네티스 클러스터 버전
# kubectl version # 로컬에 설치된 kubectl버전
# minikube kubectl 커멘드를 통해 동일 버전 사용
minikube kubectl version
```

- ALIAS 설정

```ssh
# vim ~/.zshrc
vim ~/.bashrc
source ~/.bashrc
alias k=kubectl
```

- TAB Completion 설정
  - https://github.com/zsh-users/zsh-completions#oh-my-zsh

```sh
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
    - 생략가능. 생략 될 경우 커멘드에 지정된 명령어로 수행
  - 커멘드 : 도커가 실행 할 떄 수행할 명령어, 혹은 엔트리포인트에 지정된 명령어에 대한 인자 값
    - 도커파일에 필수로 들어가야함
    - docker run 커멘드에 명시되어 있다면 CMD는 무시됨
    - 필수 파라미터이기 떄문에 명시되지 않아도 디폴트 커멘드 실행 됨

```dockerfile
# 실제실행: [Entrypoint][Command]
ENTRYPOINT ["docker-entrypoint.sh"]
CMD ["node"]
```

```shell
# 우분투 기본 COMMAND  bash를 변경-> sh
docker run --entrypoint sh ubuntu:focal
# 우분투 기본 COMMAND  bash를 변경-> echo hello world
docker run --entrypoint echo ubuntu:focal hello world
```

- 환경변수

```shell
docker run -e MY_HOST=fastcampus.com ubuntu:focal bash
# echo $MY_HOST
# env
# docker inspect [container]
# 리스트 확인 가능 Env: [...]
docker run --env list [ENV]
# 환경 변수파일로 저장하여 참조
docker run --env-file list
cat sample.env
  MY_HOST=helloworld.com
  MY_VAR=123
  MY_VAR2=456

docker run -it --env-file ./sample.env ubuntu:focal env
# dockerhub 에서 nginx 설명페이지에 환경변수 변경하여 서버 실행 방법 있음, grafana도 찾아보기
```

- 컨테이너 명령어 실행
  - `docker exec` 컨테이너의 이슈해결에 사용됨. 컨테이너 내부에 들어가서 컨피그 및 로그 확인

```shell
docker exec [container] [command]

docker run -d --name my-nginx nginx

# my-nginx 컨테이너에 Bash 쉘로 접속하기
docker exec -it my-nginx bash

# my-nginx 컨테이너에 환경변수 확인하기
docker exec -it my-nginx env
# ls -l /etc/nginx
```


- 네트워크
  - 네트워크 구조
    - docker0: 기본생성되는 브릿지 네트워크(veth와 eth간 다리 역할)
    - eth0-docker0-veth
    - eth0: 호스트(EC2 private ip)의 기본 네트워크
    - veth: 컨테이너 생성시 호스트에 해당 컨테이너에 대응되는 가상 veth가 생성 됨
    - 도커컨테이너는 내부에 lo(127.0.0.1), eth0(127.17.0.x) 네트워크 생성됨

  - 컨테이너 포트 노출
  - Expose vs Publish
  - 도커 네트워크 드라이버

```shell
docker run -p [HOST IP:PORT]:[CONTAINER PORT] [container]

# nginx 컨테이너 80포트를 호스트의 '모든'(0.0.0.0) IP의 80포트와 연결하여 컨테이너 실행
# PORTS 0.0.0.0:80->80/tcp, :::80->80/tcp
docker run -d -p 80:80 nginx
curl localhost:80
curl [ec2-private-ip]:80
curl [ec2-public-ip]:80

# nginx 컨테이너 80포트를 호스트 127.0.0.1 IP의 80포트 연결하여 실행
# PORTS 127.0.0.1:80->80/tcp, :::80->80/tcp
docker run -d -p 127.0.0.1:80:80 nginx
curl localhost:80 # 로컬(loopback 127.0.0.1)에서만 접근

# nginx 컨테이너 80포트를 호스트의 사용가능한 임의의 포트와 연결하여 실행
# PORTS 0.0.0.0:49155->80/tcp, :::49155->80/tcp
docker run -d -p 80 nginx
curl localhost:49155
curl [ec2-private-ip]:49155
curl [ec2-public-ip]:49155
```

- Expose vs. Publish
  - expose 옵션은 문서화 용도
    - `docker run -d --expose 80 nginx`
  - publish 옵션은 실제 포트를 바인딩
    - `docker run -d -p 80 nginx`

```shell
# PORTS 80/tcp (-p 퍼블리시와 달리, --expose는 문서화 용도 이기 떄문에 localhost:80 Connection 받을수없음)
docker run -d --expose 80 --name nginx-expose nginx

# Docker network Drivers 공식 문서 참조
docker network ls
  bridge
  host
  none
```

- 도커 네트워크 드라이버
  - Single Host
    - bridge(docker0)/host/none
  - Multi Host Networking
    - overlay (docker swarm)

```shell
# tedilabs/fastcampus-devops/3-docker-kubernetes/1-docker-network
ls
# bridge.sh container.sh host.sh none.sh
docker rm -f $(docker ps -a -q); docker container prune

# cat none.sh
# docker inspect > IPAddress ="", Networks > DriverOpts: null
# apk update시 오류 (network=none)
# 컨테이너의 네트워크가 필요없거나 커스텀 네트워크를 설정해야 하는 경우
#!/usr/bin/env sh
docker run -i -t --net none ubuntu:focal

# cat host.sh
# 도커가 제공하는 가상 네트워크가 아닌, 직접 호스트 네트워크에 붙음(포트바인드 없이)
# curl localhost:3000 (그라파나 디폴트 포트)
#!/usr/bin/env sh
docker run -d --network=host grafana/grafana

# cat bridge.sh
# docker0이 아닌 user-defined 네트워크 사용: fastcampus 네트워크 생성
# 브릿지네트워크에서만 사용옵션 --net-alias
#   브릿지네트워크 안에서 hello라는 도메인이름으로 컨테이너ip 서치가능하도록 내부도메인에 저장
# docker network ls
#!/usr/bin/env sh
docker network create --driver=bridge fastcampus
docker run -d --network=fastcampus --net-alias=hello nginx
docker run -d --network=fastcampus --net-alias=grafana grafana/grafana


docker exec -it [hello_container] bash
# curl이 없으므로 wget
# cd /tmp
# wget hello (도메인의 alias)


docker exec -it [grafana_container] bash
# curl grafana:3000


# veth 볼수 있음
#   ens5: 호스트 eth0
#   docker0: 브릿지
#   br-xxx : fastcampus 브릿지
ip addr
ifconfig
```

- 볼륨

```shell
docker run [app]
docker build -t app .

# 호스트의 디렉토리를 컨테이너의 특정 경로에 '마운트' 합니다
# 호스트 /opt/html 디렉토리를 Nginx의 뤱 루트 디렉토리로 마운트
docker run -d \
  --name nginx \
  -v /opt/html:/usr/share/nginx/html
  nginx



# tedilabs/fastcampus-devops/3-docker-kubernetes/2-docker-volume
docker ps -a
docker run -d -it ubuntu:focal
docker exec -it [container] bash
# cat > hello
# hello world!
# exit
docker rm -f [container]
# 볼륨 사용없이 실행 시 컨테이너 삭제시 삭제됨
docker run -d -it ubuntu:focal
docker exec -it [container] bash


# 호스트 볼륨
# 호스트의 $(pwd)/html 디렉토리에는 index.html이라는 파일이 있음
# 다음 스크립트 실행 시 nginx 컨테이너의 디렉토리에 index.html이 마운트 됨
# curl localhost로 확인
# cat host-volume.sh
#!/usr/bin/env sh
docker run \
  -d \
  -v $(pwd)/html:/usr/share/nginx/html \ # 호스트볼륨:컨테이너볼륨
  -p 80:80 \
  nginx

docker exec -it [container] bash
# ls /usr/share/nginx/html
# cat index.html
# cat > hello
# hello world!
# exit
ls $(pwd)/html
# 컨테이너 안에서 만들었던 hello라는 파일이 호스트에도 생성됨!
```

- 볼륨 컨테이너
  - 특정 컨테이너의 볼륨 마운트를 공유할 수 있음

```shell
docker run -d \
  --name my-volume
  -it \
  -v /opt/html:/usr/share/nginx/html \
  ubuntu:focal

# my-volume 컨테이너의 볼륨을 공유
docker run -d \
  --name nginx \
  --volumes-from my-volume \
  nginx



cat volumne-container.sh
#!/usr/bin/env sh
docker run \
  -d \
  -it \
  -v $(pwd)/html:/usr/share/nginx/html \
  --name web-volume \
  ubuntu:focal

docker run \
  -d \
  --name fastcampus-nginx \
  --volumes-from web-volume \
  -p 80:80 \
  nginx
  
docker run \
  -d \
  --name fastcampus-nginx2 \
  --volumes-from web-volume \
  -p 8080:80 \
  nginx


curl localhost:80
curl localhost:8080
docker inspect fastcampus-nginx
docker inspect fastcampus-nginx2
```

- 도커 볼륨 (호스트볼륨과 대비)
  - 호스트볼륨은 호스트에 마운트할 경로 지정했어야 했음
  - 도커가 제공하는 볼륨 관리기능을 활용하여 데이터 보존 (볼륨 생성, 삭제 등 관리)
  - 이런 도커 볼륨은 도커가 관리하는 특정 호스트 경로에 데이터 저장
  - 기본적으로 /var/lib/docker/volumes/$(volume-name}/_data에 데이터가 저장됨 (호스트에 쌓임)

```shell
# 'db' 라는 도커 볼륨 생성
docker volume create --name db

# 도커의 'db' 볼륨을 mysql 루트 디렉토리로 마운트
docker run -d \
  --name fastcampus-mysql \
  -v db:/var/lib/mysql \
  -p 3306:3306 \
  mysql:5.7


cat docker-volume.sh
#!/usr/bin/env sh
docker volume create --name db
docker volume ls

# 환경변수 지정: 최초생성 데이터베이스 이름
# 환경변수 지정: 루트계정의 비밀번호
# 호스트 볼륨대신 도커볼륨 사용
docker run \
  -d \
  --name fastcampus-mysql \
  -e MYSQL_DATABASE=fastcampus \
  -e MYSQL_ROOT_PASSWORD=fastcampus \
  -v db:/var/lib/mysql \
  -p 3306:3306 \
  mysql:5.7

docker ps
docker volume inspect db
sudo -s
ls -l /var/lib/docker/volumes/db/_data

docker rm -rf $(docker ps -a -q)
docker volume rm db
```

- 읽기전용 볼륨 연결
  - 볼륨 연결 설정에 :ro 옵션을 통해 읽기 전용 마운트 옵션을 설정
  - 변경 되어서는 안되는 디렉토리나 파일 연결 시

```shell
# 도커의 web-volume 볼륨을 nginx의 웹 루트 디렉토리로 읽기 전용 마운트
docker run -d \
  --name nginx \
  -v web-volume:/usr/share/nginx/html:ro \
  nginx

cat host-volume.sh

#!/usr/bin/env sh
docker run \
  -d \
  -v $(pwd)/html:/usr/share/nginx/html:ro \
  -p 80:80 \
  --name ro-nginx \
  nginx

# 에러발생! : read-only file system!!!
docker exec ro-nginx touch /usr/share/nginx/html/test
```

- 로그
  - stdout/stderr
    - 각언어에서 지원하는 로그프레임워크로 어플리케이션 단에서 처리했지만, 도커컨테이너에서는 stdout/stderr 사용
    - APP container (stdou/stderr) -> logging driver
  - 로그확인하기
  - 호스트 운영체제의 로그 저장 경로
  - 로그 용량 제한하기
  - 도커 로그 드라이버

```shell
# 전체 로그 확인
docker logs [container]
# 마지막 10줄
docker logs --tail 10 [container]
# 실시간 로그 스트림
docker logs -f [container]
# 로그마다 타임스탬프 확인
docker logs -f -t [container]
```

- 호스트 운영체제의 로그 저장 경로
  - log driver = json-file 했을때만 유효

```shell
sudo su
cat /var/lib/docker/containers/${CONTAINER_ID}/${CONTAINER_ID}-json.log
less -R /var/lib/docker/containers/${CONTAINER_ID}/${CONTAINER_ID}-json.log
```

- 로그 용량 제한
  - 컨테이너 단위 또는 도커엔진에서 기본설정 진행(운영환경에서는 필수 설정!)

```shell
# 한 로그 파일 당 최대 크기를 3MB제한, 최대 로그 파일 5개로 로테이팅
docker run \
  -d \
  --log-driver=json-file \
  --log-opt max-size=3m \
  --log-opt max-file=5 \
  nginx
```

- 도커 로그 드라이버

- 도커 이미지
  - Layer Architecture
  - 도커 컨테이너 생성시 이미지레이어(Read Only) & 컨테이너 레이어(Read/Write)

```shell
docker images
# RootFS 레이어 배열리스트 확인
docker image inspect nginx:latest
```

- Dockerfile 없이 이미지 생성
  - docker commit
  - 기존 컨테이너를 기반으로 새 이미지를 생성 할 수 있음

```shell
# 컨테이너 -> 이미지
docker commit [OPTIONS] CONTAINER [REPOSITORY[:TAG]]
docker commit -a fastcampus -m "First Commit" ubuntu my_ubuntu:v1


docker run -it --name my_ubuntu ubuntu:focal
  cat > my_file
  Hello Fastcampus!
  # CTRL+P,Q

docker ps

# 컨테이너를 이미지로 저장
# -a 작성자
# -m 커밋 메시지
# my_ubuntu 컨테이너 이름
# my-ubuntu:v1 저장할 이미지 이름
docker commit -a fastcampus -m "add my_file" my_ubuntu my-ubuntu:v1
# my-ubuntu 이미지 생성 확인
docker images
# 레이어 두개 있음(기반으로 만든 레이어 ubuntu:focal도 그중 하나)
docker image inspect my-ubuntu
# 이 이미지를 만드는데 사용 된 ubuntu:focal 이미지 확인하기(-> 해당 레이어가 my-ubuntu 이미지 레이어에 포함 됨)
docker image inspect ubuntu:focal

# 컨테이너 삭제후 다시 신규 이미지를 사용해 컨테이서 생성 및 실행
docker rm -f my_ubuntu
docker run -it my-ubuntu:v1
# cat my_file
# exit
```

- 도커파일 Dockerfile 이용한 이미지 생성
  - Dockerfile로 이미지 생성

```shell
# docker build [OPTIONS] PATH
# ./ 디렉토리를 빌드 컨텍스트로 my_app:v1 이미지 빌드 (Dockerfile 이용)
# 현재 디렉토리 기반으로 도커빌드를 할건데, 빌드 결과물 이미지는 my_app:v1이라는 태그를 생성
# -t 빌드할 이미지에 태그 지정
docker build -t my_app:v1 ./

# 다른 도커파일 MyDockerfile로 빌드
docker build -t my_app:v2 -f example/MyDockerfile ./

# cd fastcampus-devops/3-docker-kubernetes/3-dockerfile/app
# sudo apt install tree & tree -L 2
cat Dockerfile

FROM node:12-alpine
RUN apk add --no-cache python3 g++ make
WORKDIR /app
COPY . .
RUN yarn install --production
CMD ["node", "src/index.js"]


cd app
# Sending build context to Docker daemon xxMB
docker build -t my-app:v1 ./
# my-app:v1 이미지 생성 확인
docker images

# 도커 파일 수정 후 재빌드
# CMD ["node", "src/index.js", "1"]
# Using cache (기존에 빌드된 부분 일부 레이어는 캐시 사용)
docker build -t my-app:v2 ./
docker images
docker run -d my-app:v2
docker ps
docker exec -it CONTAINER_ID sh
```

- 빌드 컨텍스트
  - 도커 빌드 명령 수행 시 현재 디렉토리(Current Working Directory)를 빌드 컨텍스트(Build Context)라고 함
  - Dockerfile로 부터 이미지 빌드에 필요한 정보를 도커 데몬에 전달하기 위한 목적
  - `COPY . .` 현재 디렉토리 내용을  컨테이너 이미지 내부로 복사
  - 현재 디렉토리 정보를 알고 있으려면 도커 데몬이 빌드 컨텍스트 정보를 알고 있어야 함
  - 현재 디렉토리에 파일이 많다면 빌드 오래 걸릴 수 있음-> `.dockerignore` 활용


- Dockefile 작성
  - Dockerfile 공식 문서 참고

```shell
# cd fastcampus-devops/3-docker-kubernetes/3-dockerfile/nodejs-server/
cat Dockerfile
#
# nodejs-server
#
# build:
#   docker build --force-rm -t nodejs-server .
# run:
#   docker run --rm -it --name nodejs-server nodejs-server
#

FROM node:16
# 이미지 메타데이터
LABEL maintainer="FastCampus Park <fastcampus@fastcampus.com>"
LABEL description="Simple server with Node.js"

# 컨테이너안에 디렉토리 설정 (cd 명령어를 통해)
# Create app directory
WORKDIR /app

# COPY [SRC 호스트 운영체제] [DEST 이미지 상에서 경로]
# ./ 현재 디렉토리는 WORKDIR에서 지정한 /app 디렉토리임!
# 호스트 파일시스템에 있는 package*.json 파일들을 도커 이미지 디렉토리 WORKDIR /app에 복사!
# Install app dependencies
# A wildcard is used to ensure both package.json AND package-lock.json are copied
# where available (npm@5+)
COPY package*.json ./

# 도커 이미지 상에서 해당 명령어를 실행하라
# /app에 .json을 복사하고, 여기서 npm install을 실행
RUN npm install
# If you are building your code for production
# RUN npm ci --only=production

# 현재 호스트 디렉토리의 모든 파일과 디렉토리를 도커이미지 /app에 복사
# 위에서 진행된 소스코드의 변경사항을 적용하기 위함
# Bundle app source
COPY . .

# 도커이미지 포트 8080 사용
# 퍼블리싱 -p가 아닌 명시용 목적 (이미지 사용하는 유저에게 알려주는 목적)
EXPOSE 8080

# 해당 이미지로 컨테이너 실행시 수행해야할 명령어 (배열리스트 또는 한개의 문자열 "node server.js" 지정 가능) 
CMD [ "node", "server.js" ]
```


```shell
# build:
docker build --force-rm -t nodejs-server .
docker images

# run:
docker run --rm -it --name nodejs-server nodejs-server
docker ps
# -p 8080 퍼블리싱 하지 않았기 때문에 해당 포트 열려있지 않음
curl localhost:8080 
docker rm -f [container]
docker run -d -p 8080:8080 --name nodejs-server nodejs-server
curl localhost:8080
```

- ENTRYPOINT 엔트리포인트
  - CMD 앞서 시작프로그램 지정
  - ENTRYPOINT 와 CMD가 합쳐져 사용됨

```shell
# Dockerfile
ENTRYPOINT ["executable", "param1", "param2"]
ENTRYPOINT command param1 param2
```

- ADD
  - COPY와 비슷
  - 차이점은 소스디텍토리로 url도 받을 수 있음 (소스가 변경됐는지 확인 어렵기 때문에 ADD 사용 지양)
  - COPY 사용 권장

- USER
  - 컨테이너가 사용하게 될 기본 사용자/ 그룹 지정 가능
  - 보안관련 옵션

- 도커이미지 저장, 불러오기
  - 인터넷이 없는 환경에 유용

```shell
# docker save -o [OUTPUT-FILE] IMAGE
# ubuntu:focal 이미지를 ubuntu_focal.tar 압축파일로 저장
docker save -o ubuntu_focal.tar ubuntu:focal

# docker load -i [INPUT-FILE]
docker load -i ubuntu_focal.tar


docker images
# my-app:v2 이미지를 압축파일로 저장
docker save -o my-app-v2.tar my-app:v2
file my-app-v2.tar

# 이미지 제거
docker rmi my-app:v2
# my-app:v2 이미지 제거 확인
docker images

# 다시 압축파일로부터 이미지를 불러옴
docker load -i my-app-v2.tar
docker images
```


- 도커허브 저장소
  - hub.docker.com 로그인
  - Security > New Access token
  - Repository > Create Repository (Private)
    - `REPO_NAME/my-nginx`
  - 도커이미지 업로드 시 : `docker tag`, `docker push`


```shell
docker ps
docker login -u [ID]
# Password: 액세스토큰
# ~/.docker/config.json에 저장됨 -> 안전하게 따로 보관

docker images

# 기존 이미지의, 리포지토리와 태그명만 바꾼 이미지 생성하여
# 도커허브 리포지토리에 업로드
# USERNAME/REPO_NAME:TAG
docker tag nginx:latest jnuho/my-nginx:v1.0.0
docker images
docker push jnuho/my-nginx:v1.0.0


# 도커 이미지 pull 받을 수 있는지 확인
docker rmi jnuho/my-nginx:v1.0.0
docker pull jnuho/my-nginx:v1.0.0
docker images
```


- AWS ECR 저장소 이용
  - ECR(Elastic Container Registry): 컨테이너 이미지 저장소
    - Create Repository 'Private' > 'my-nginx'
    - Image scan settings, KMS encryption : 'Disabled'

```shell
# Nginx 도커 이미지 다운로드
docker pull nginx

# aws-cli 사용
aws sts get-caller-identity
export AWS_PROFILE=fastcampus
aws sts get-caller-identity

# AWS ECR > View Push Commands
# 로그인
aws ecr get-login-password --region ap-northeast-2 | \
  docker login \
  --username AWS \
  --password-stdin [AWS_ID].dkr.ecr.ap-northeast-2.amazonaws.com

# ECR로 푸시
docker tag nginx:latest [AWS_ID].dkr.ecr.ap-northeast-2.amazonaws.com/my-nginx:v1.0.0
docker images
docker push [AWS_ID].dkr.ecr.ap-northeast-2.amazonaws.com/my-nginx:v1.0.0

# 삭제후 다시 pull 테스트
docker rmi [AWS_ID].dkr.ecr.ap-northeast-2.amazonaws.com/my-nginx:v1.0.0

# credential 있는 머신에만 다운로드 가능
docker pull [AWS_ID].dkr.ecr.ap-northeast-2.amazonaws.com/my-nginx:v1.0.0
```

- 컨테이너 경량화
  - 꼭 필요한 패키지 및 파일만 추가
  - 컨테이너 레이어 수 줄이기 (도커파일 RUN 지시어 한개의 레이어 생성)
  - 경량 베이스 이미지 선택 (debain slim, alpine, stretch)
    - FROM node:16-slim
  - 멀티 스테이지 빌드 사용 (멀티 스테이지 파이프라인)
    - `FROM node:16-alpine AS base`
    - `FROM base AS build`
    - `FROM base AS release`


```shell
cd fastcampus-devops/3-docker-kubernetes/3-dockerfile/
cat slacktree/Dockerfile

FROM alpine:3.14
LABEL  maintainer="FastCampus Park <fastcampus@fastcampus.com>"
LABEL description="Simple utility to send slack message easily."

# RUN 한개에 여러개 커멘드 구성 (레이어 수 1개로)
# --no-cache 옵션으로 캐시 삭제
# install needed package
RUN \
  apk add --no-cache bash curl git && \
  git clone https://github.com/course-nero/slacktee /slacktee && \
  apk dell --no-cache git

RUN chmod 755 /slacktee/slacktee.sh
```

```shell
# --force-rm=true 빌드 성공,실패 상관없이 intermediate 컨테이너 삭제
docker build --force-rm -t nodejs-server:slim -f Dockerfile.alpine ./
```

- 멀티스테이지 파이프라인 : 블록에 임시 이미지 이름 'base'

``` shell
FROM node:16-alpine AS base
LABEL maintainer=""
LABEL desctiption=""
WORKDIR /app
COPY package*.json ./
```


- 도커 데몬

```shell
docker system info

# 새롭게 발생하는 도커 이벤트 로그 (ex: docker run nginx)
# ALIAS: docker events
docker system events

docker run --name

# 우분투 OS
journalctl -u docker
df -h
# RECLAIMABLE 회수 가능 용량
docker system df
# 중지된 컨테이너, 이미지, 네트워크 삭제
docker system prune
# 컨테이너별로 CPU, IO, MEMORY 등
docker stats
docker run -it -d mysql:5.7
docker run -it -d nginx
```

- 도커 컴포즈
  - 명시적으로 여러 컨테이너 관리하기 (docker-compose.yml 파일을 통해)
    - 네트워크, 볼륨, 서비스 의존성, 디스커러리 자동화 및 컨테이너 수평 확장
  - 단일 서버에서 여러 컨테이너를 프로젝트 단위로 묶어서 관리
    - 도커 컴포즈 프로젝트 안에, 네트워크, 볼륨, 서비스 등 의존성 관리
    - 서비스들 s1, s2, s3 간의 의존성 관리
      - 도커엔진에서는, 서비스의 내부 도메인 명으로 호출시 브릿지네트워크 net alias
      - 도커컴포즈에서는 서비스명으로 호출 가능 (서비스 디스커버리), 컨테이너 수평확장 용이
  - 프로젝트/서비스/컨테이너
    - 프로젝트 : 워크 스페이스; 서비스 컨테이너 묶음
    - 서비스 : 컨테이너를 관리하기 위한 단위
    - 컨테이너 : 서비스를 통해 컨테이너 관리
  - docker-compose.yml 각 서비스 정의
    - 최상위 4가지 옵션: version, services, networks, volumes
    - EX) scale=3 : 서비스 3개 실행 가능
  - docker-compose 명령어
    - 버전3는 도커 스왐과 호환
    - `docker-compose deploy`
    - `docker stack`
    - 도커 스왐 (쿠버네티스와 동일 목적, 시장 점유율 낮음)


- 도커 컴포즈 기본 명령어
- 
```shell
cd fastcampus-devops/3-docker-kubernetes/4-docker-compose
cd build

# redis, flask 사용
cat app.py

cat docker-compose.yml

docker-compose -h
docker-compose version

# docker run과 비슷, 이미지 받아서, 컨테이너 생성
# 디폴트 네트워크 생성(브릿지 네트워크, build_default: 프로젝트디렉토리명_default), 컨테이너 2개 (build_web_1) 생성
# CTRL+C: 컨테이너 중단
docker-compose up
# -p: 프로젝트 명 명시
docker-compose -p my-project up -d

# 도커 컴포즈 v2 이상부터 사용가능
docker-compose ls -a

# 프로젝트 내 컨테이너 및 네트워크 종료 및 제거
docker-compose down

# 프로젝트 내 컨테이너 및 네트워크 및 볼륨 종료 및 제거
docker-compose down -v

cd wordpress
docker-compose up -d
docker-compose down -v
```

- 서비스 확장

```shell
cd build
docker-compose ls
docker-compose ps
# 해당 프로젝트 내에 컨테이너 목록
docker-compose -p my-project ps

# 스케일링
docker-compose -p my-project up --scale web=3 -d
# web-1, web-2, web-3 확인
docker-compose -p my-project ps

```

```yaml
# docker-compose.yml
# web을 스케일링 할때 주의할 점:
#  1.중복된 호스트포트는 사용할 수 없기때문에, [호스트포트]:[컨테이너포트]로 지정하면 안됨
#  2. web: container_name: "web" 옵션 사용하면 안됨!
version: "3.9"
services:
  web:
    container_name: "web"
    build: .
    ports:
      - "5000"
  redis:
    image: "redis:alpine"
```

- 프로젝트 내 추가 명령어
```shell
docker-compose logs
docker-compose -p my-project logs -f

docker-compose events
docker-compose -p my-project events

docker-compose images
docker-compose -p my-project images

docker-compose ps
docker-compose -p my-project ps

docker-composet top
docker-compose -p my-project top
```

- 도커 컴포즈 yaml: serivces의 depends_on으로 컨테이너 실행순서 설정
  - 이전 컨테이너의 실행은 보장하지만 준비된 것은 보장X
  - 엔트리 포인트를 통해 특정 shell 스크립트 실행하여 이전 컨테이너의 준비여부 판단하여 커멘드 실행
  - 서비스 ports는 docker run의 -p 옵션과 같음

- 도커 컴포즈 사용목적
  - 로컬 개발환경 구성
  - 자동화된 테스트 환경 구성
    - CI/CD 파이프라인 중 쉽게 격리된 테스트 환경을 구성하여 테스트 수행 가능
  - 단일 호스트 내 컨테이너를 선언적 관리


- 도커 컴포즈 실습 - Grafana, MySql
  - Grafana 도커 [가이드](https://grafana.com/docs/grafana/latest/installation/docker/)
  - MySQL 도커 [가이드](https://hub.docker.com/_/mysql)


```shell

# 그라파나-only
cd grafana-only
cat docker-compose.yml

docker-compose up -d
docker-compose ps
curl localhost:3000
  그파라나 로그인 admin/admin > Server Admin > Database : sqlite

cat grafana-only/files/grafana.ini

docker-compose down -v


# 그라파나-Mysql
# 환경변수 : mysql 도커 문서 > Where to Store Data
# 디비컨테이너 생성 된 후, 그라파나 서비스 생성
#  grafana depends on: - db
cd grafana-mysql
cat docker-compose.yml
docker-compose up -d
docker-compose ps

# DB영구적 저장확인을 위해 서비스 내리기 (볼륨은 그대로 두기)
# 볼륨이 호스트에 영구적으로 보관 됨
docker-compose down
docker volume ls
docker-compose up -d
docker-compose ps
```




- 쿠버네티스 아키텍쳐
  - master nodes (Control Plane) / worker nodes (Workerload Plane)
  - Control Plane
    - Kubernetes API
    - etcd
    - Scheduler
    - Controllers
  - Workerload Plane
    - Kublet
    - Container Runtime
    - Kubernetes Service Proxy
    

- 쿠버네티스가 앱을 실행하는 과정
  - DEFINE your application
    - yaml, json 파일을 통해 Object들을 정의
    - e.g. Deployment 2개, Services 2개로 manifest를 만들어 어플리케이션 배포
    - 1. Kubernetes API를 통해 manifest의 Object를 생성하여 etcd 저장
    - 2. 컨트롤러가 각 application 인스턴스에 해당하는 Object 생성
    - 3. 스케쥴러가 각각의 인스턴스에 node 배정
    - 4. Kublet이 instance가 node에 배정된 것을 notice받고, Container Runtime을 통해 application instace를 실행
    - 5. Kube Proxy는 application instance가 클라이언트로 부터 connection을 받을 준비가 되면 load balancer 설정함
    - 6. Kublet과 Controller는 system을 모니터링하고 application이 실행 상태를 유지하도록 함
  - SUBMIT application to the API
    - `kubectl` 커맨드라인으로 yaml, json파일을 오브젝트 단위로 나눠서 API에 전달

- CONTROLLER
  - API서버로 부터 Object 생성알림을 받으면, Kubernetes API  통해 Object 생성

- SCHEDULER
  - 컨트롤러 타입으로, application 인스턴스를 워크노드에 스케쥴하는 역할
  - 각 어플리케이션 인스턴스 오브젝트의 최적의 워커노드를 찾아 인스턴스에 할당 (API를 통해 object를 수정하여)

- KUBELET AND THE CONTAINER RUNTIME
  - 각 워커노드에서 실행되는 kublet은 컨트롤러 타입 중 하나.
  - 해당 워커 노드에 어플리케이션 인스턴스가 할당 될때까지 기다리다가, 어플리케이션을 실행
  - 컨테이너 runtime이 해당 어플리케이션 컨테이너를 실행하도록 함

- KUBE PROXY
  - 어플리케이션 deployment는 여러 어플리케이션 인스턴스로 구성될 수 있기 때문에, 로드밸런서가 이들을 하나의 IP로 expose 해야함
  - Kube Proxy도 컨트롤러 타입으로서 로드밸런서를 만드는 역할을 함


- KEEPING THE APPLICATION HEALTHY
  - 어플리케이션이 실행되면, kubelet은 어플리케이션이 종료되는 경우, 재시작하여 healthy한 상태를 유지함
  - 어플리케이션 인스턴스를 represent하는 object를 업데이트하여, 어플리케이션의 status를 report함





- Docker Container

```sh
ls kaida-0.1
  Dockerfile  Makefile  app.js  html/
cat Dockerfile
  FROM node:16
  COPY app.js /app.js
  COPY html /html
  ENTRYPOINT ["node", "app.js"]
docker build -t kiada:latest .
docker run --namee kiada-container -p 1234:8080 -d kiada
docker tag kiada jnuho/kiada:0.1

docker login -u jnuho docker.io
docker push jnuho/kiada:0.1

# Run the image on other Hosts
docker run --name kiada-container -p 1234:8080 -d jnuho/kiada:0.1

```

- API Resources

```sh
kubectl api-resources
kubectl explain pod
kubectl explain configmap
kubectl explain service

# 파드: pods, pod, po
kubectl get pod --all-namespaces
kubectl get po --all-namespaces
```

- 쿠버네티스는 오브젝트를 yaml기반 매니페스트 파일로 관리
  - apiVersion : 오브젝트가 어떤 API 그룹에 속하고 api버전은 몇인가
  - kind : 오브젝트가 어떤 API 리소스인가
  - metadata : 오브젝트 식별하기 위한 정보 (이름, 네임스페이스, 레이블)
    - Label (optional) : 정확한 리소스 이름 정의, 소유자, 종류 명시하여 식별하기 위한 목적
    - Annotations  : 쿠버네티스 애드온이 해당 오브젝트를 어떻게 처리할지 결정하기위한 설정 용도
      - e.g. LogAgent 애드온의 로그 수집설정
  - spec : 오브젝트가 가지고자 하는 데이터
    - API 리소스에 따라 spec 대신 data,rules,subjects 등 다른 속성 사용



- `kubectl` : 쿠버네티스 클라이언트
  - Kubernetes API Server와 통신 할 수 있는 툴

- 클러스터 생성하는 방법
  1. minikube 싱글 노드 클러스터
  2. kind 멀티 노드 클러스터
  3. GKE, EKS, kind 멀티 노드 클러스터


```sh
### METHOD 1. minikube로 싱글 node 클러스터 생성

# Create a single-node Kubernetes Cluster to test Kubernetes functionalities
# testing certain features related to managing apps on multiple nodes are limited

minikube start --driver=docker

kubectl cluster-info

### METHOD 2. kind
# 문서: https://kind.sigs.k8s.io/docs/user/quick-start/
# 다른 minikube와 같은 툴에비해 아직 unstable

curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.12.0/kind-linux-amd64
chmod +x ./kind
mv ./kind /usr/local/kind
ls -al /usr/local/kind
  drwxr-xr-x  2 root root 4096 Apr  5 15:20 kind
  alias k=kubectl
  complete -F __start_kubectl k
  source <(kubectl completion bash)
  export PATH=/usr/local/kind:$PATH

vim ~ubuntu/.bashrc
  export PATH=/usr/local/kind:$PATH
source .bashrc

cat kind-multi-node.yaml
  kind: Cluster
  apiVersion: kind.sigs.k8s.io/v1alpha3
  nodes:
  - role: control-plane
  - role: worker
  - role: worker

# 3-node cluster
kind create cluster --config kind-multi-node.yaml

kind get nodes
  kind-worker
  kind-control-plane
  kind-worker2

# Instead of using Docker to run containers, nodes created by kind use the CRI-O container runtime
crictl ps
docker ps


### METHOD 3. GKE로 멀티 node 클러스터 생성 (Google 가입 및 신용카드 등록 필요)

# Signing up for a Google account, in the unlikely case you don’t have one already.
# Creating a project in the Google Cloud Platform Console.
# Enabling billing. This does require your credit card info, but Google provides a
#   12-month free trial. And they’re nice enough to not start charging automatically after the free trial is over.)
# Enabling the Kubernetes Engine API.
# Downloading and installing Google Cloud SDK. (This includes the gcloud
#   command-line tool, which you’ll need to create a Kubernetes cluster.)
# Installing the kubectl command-line tool with `gcloud components install kubectl`

# After completing the installation, you can create a Kubernetes cluster with three
#   worker nodes using the command shown in the following listing.
# Then interact with cluster using `kubectl` which issues REST request
# to the Kubernetes API Server running on the master node

# set default zone
# region europe-west3 has three different zones
# but here we set all nodes to use the same zone 'europe-west3-c'
gcloud config set compute/zone europe-west3-c

# create kubernetes cluster
# three worker nodes in the same zone
# thus --num-nodes indicates the number of nodes per zone
# if the regions contains three zones and you only want three nodes, set --num-nodes=1
gcloud container clusters create kiada --num-nodes 3

gcloud compute instances list

# scale the number of nodes
gcloud container clusters resize kiada --size 0

### 클러스터 정보 확인

# explore what's running inside a node; you can check logs inside it
gcloud compute ssh <node-name>

# check if cluster is up by listing cluster nodes
#   minikube : 1 node in the cluster
#   GKE : 3 nodes in the cluster



### METHOD 4. EKS
# Install `eksctl` command-line tool
# https://docs.aws.amazon.com/eks/latest/userguide/getting-started-eksctl.html

# creates a three-node cluster in the eu-central-1 region
eksctl create cluster --name kiada --region eu-central-1 --nodes 3 --ssh-access
```

- 쿠버네티스 툴 `kubectl`

```
kubectl get nodes
kubectl describe nodes
kubectl describe node <node-name>
```

- kubectl 이 특정 Kubernetes cluster를 사용 하도록 설정
  - point kubectl to it by setting the KUBECONFIG environment variable as follows:

```sh
cat ~/.kube/config
export KUBECONFIG=/path/to/custom/kubeconfig
```

- Kubernetes Dashboard

```sh
# https://github.com/kubernetes/dashboard
# add a configuration to a resource by file name or stdin
# k apply [options]

kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.5.1/aio/deploy/recommended.yaml

# This command runs a local proxy to the API server, allowing you to access the services through it. Let the proxy process run and use the browser to open the dashboard at the following URL:

kubectl proxy

# 대시보드 URL:
# http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/

# Token 생성 해야함:
# https://github.com/kubernetes/dashboard/blob/master/docs/user/access-control/creating-sample-user.md#getting-a-bearer-token

kubectl -n kubernetes-dashboard get secret $(kubectl -n kubernetes-dashboard get sa/admin-user -o jsonpath="{.secrets[0].name}") -o go-template="{{.data.token | base64decode}}"


kubectl -n kubernetes-dashboard describe secret $(kubectl -n kubernetes-dashboard get secret | sls admin-user | ForEach-Object { $_ -Split '\s+' } | Select -First 1)


# Using Helm to install dashboard?  Install helm :
# https://medium.com/@munza/local-kubernetes-with-kind-helm-dashboard-41152e4b3b3d
curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
chmod 700 get_helm.sh
./get_helm.sh

helm repo add kubernetes-dashboard https://kubernetes.github.io/dashboard/
helm install dashboard kubernetes-dashboard/kubernetes-dashboard -n kubernetes-dashboard --create-namespace

```


- 클러스터에 app 배포하기 (Deployment)
  - 보통은 yaml, json 파일에 Object 정의 하여 배포
    - (처음에는 일단 manifest 파일없이 Deployment 만들어보기)
  - `kubectl create` -> Kubernetes API에 전달되어 -> Deployment 생성 -> Pod 오브젝트 생성
    - 워커 노드에 assigned/scheduled 됨
    - 워커 노드의 Kublet (쿠버네티스 Agent)는 Pod가 새로 해당 node에 스케쥴된것을 확인
    - 도커 registry에서 이미지를 받아, 컨테이너로 실행 
    - Deployment 오브젝트는 클러스터에 배포된 어플리케이션을 의미
    - 기존에 이미지로 만들어 컨테이너에 배포했던 kiada를 쿠버네티스 클러스터에 적용해보기
  - pod 생성
    - 쿠버네티스는 컨테이너 단위가 아닌 Pods 단위 관리
    - Pod: multiple co-located 컨테이너; 컨테이너 그룹
    - 같은 worker노드와 리눅스 namespace위에서 실행되는 밀접하게 연결된 컨테이너 그룹
    - 각각의 Pod은 separate logical machine with same IP, hostname, process, and so on.
    - 하나의 어플리케이션을 실행 함
  - service 생성
    - 각 pods들은 IP가 할당되지만 클러스터로의 internal ip라서 외부에서 접근 불가능
    - 서비스 오브젝트를 생성하여 external ip 할당하여 외부에 노출
    - LoadBalancer 타입의 서비스를 생성-> external load balancer를 생성되고
    - 이 로드밸런서의 퍼블릭 IP를 통해 pod에 접근 가능
    - pod는 ephemeral함. pod가 실패하거나 다운되고 다시 생성되면 ip변동-> service를 통해 static ip 할당

```sh
### POD 생성

# OLD VERSION : 
# --generator=run/v1 옵션은 Deployment대신 ReplicationController를 create
# kubectl run kubia --image=jnuho/kubia --port=8080

# kubectl로 dokcerhub 이미지를 쿠버네티스 클러스터에서 실행
# create all the necessary components without having to deal with JSON or YAML
# Create a deployment object, called 'kiada'
# deployment to use the container image jnuho/kiada:0.1
# the Deployment object is stored in Kubernetes API
# now the jnuho/kiada container is run in the Cluster

kubectl create deployment kiada --image=jnuho/kiada:0.1
  deployment.apps/kiada created

# 생성된 Deployment 리스트 : NOT READY (컨테이너가 준비되지 않음)
# 다만 컨테이너 조회 하는 커멘드는 없음 e.g. k get containers 는 없음
# 컨테이너가 아닌 PODS가 쿠버네티스에서 가장 작은 단위임!
# POD 내의 컨테이너는 네트워크와 UTS 네임스페이스를 공유
# 각 POD은 논리적인 단위의 컴퓨터로서 어플리케이션을 실행
# 하나의 노드에 여러개의 POD가 있더라도, 각 POD들은 서로 다른 POD들의 프로세스를 볼 수 없음
k get deployments


# 생성된 pods 정보 (Pending/ContainerCreating -> Running)
kubectl get pods
kubectl describe pods

# 생성된 pod 상세 정보에 worker노드에서 실행 중 확인
kubectl describe pod <pod-name>
  Name:         kiada-7fc9cfcf4b-t8fcg
  Namespace:    default
  Priority:     0
  Node:         kind-worker/172.18.0.2


### SERVICE 생성
# Exposing your application to the world

# Create Service Object
#   expose all pods that belong to the kiada Deployment as a new service.
#   pods to be accessed from outside the cluster via a load balancer.
#   application listens on port 8080, so you want to access it via that port.
#   로드밸런서 타입 서비스는 cluster내에만 expose 시키거나 퍼블릭IP로 외부로 expose 가능

kubectl expose deployment kiada --type=LoadBalancer --name kiada-http --port 8080

# load balancer가 생성되고 EXTERNAL-IP 할당됨
kubectl get services
kubectl get svc

# List all available resources - Kubernetes Objects
kubectl api-resources


# 32232 is the node port on the worker node
# that forwards connections to your service.
# 이 포트로 모든 워커 노드에서 서비스에서 접근가능
# (minikube, kind, 등등 클러스터 종류 상관 없이)

regardless of whether you’re using Minikube or any other Kubernetes cluster.
kubectl get svc kiada-http
  NAME         TYPE           CLUSTER-IP     EXTERNAL-IP   PORT(S)          AGE
  kiada-http   LoadBalancer   10.96.59.164   <pending>     8080:32232/TCP   83m
  kubernetes   ClusterIP      10.96.0.1      <none>        443/TCP          6h50m


# Kubernetes 클러스터가 클라우드 (AWS, Gcloud)에 deploy 된다면
# 해당 클라우드 인프라에 로드밸런서를 통해
# 클러스터로 오는 트래픽을 실행 중인 컨테이너로 forward 함
# 해당 인프라는 Kubernetes에게 로드밸런서의 Ip를 알려주고 서비스의 외부주소가 됨


# Accessing your application through Load Balancer
curl <EXTERNAL-IP>:8080
curl 104.155.74.57:8080


# Accessing your application without Load Balancer
#   not all kubernetes clusters provide Load Balancer
#   minikube shows where to access the services
#     prints the url of the service
minikube service kiada --url
curl <ip:port>


# you can access this service locally using the Kubectl proxy tool.
kubectl port-forward kiada-http 8080:3002

# Deploying your application REQUIRES only two steps
# 'Deployment' is a representation of an application
# 'Service' exposes that deployment

kubectl create deployment [options]
kubectl expose deployment [options]

```


- Horizontally Scaling the application
  - increase (scale-out) / decrease the number of running application instances
  - scale the deployment object



```sh
# the result of 'Scale Out'
k get deploy
k get deployments.apps
k get deployment

  NAME    READY   UP-TO-DATE   AVAILABLE   AGE
  kiada   3/3     3            3           31h

k get pods

  NAME                     READY   STATUS    RESTARTS       AGE
  kiada-7fc9cfcf4b-64qvl   1/1     Running   0              3m3s
  kiada-7fc9cfcf4b-mf8ls   1/1     Running   0              3m3s
  kiada-7fc9cfcf4b-t8fcg   1/1     Running   1 (148m ago)   31h
```


- Displaying the pods' host node when listing pods
  - To see which nodes the pods were scheduled to,
  - you can use the -o wide option to display a more detailed pod list

- Object created
  - Deployment object you created,
  - Pod objects that were automatically created based on the Deployment
  - Service object you created manually
    - Instead of connecting directly to the pod,
    - clients should connect to the IP of the service.
    - This ensures that their connections are always routed to a healthy pod,
    - Each service provides internal load balancing in the cluster,
    - but if you set the type of service to LoadBalancer,
    - Kubernetes will ask the cloud infrastructure it runs in for an additional
    - load balancer to make your application available at a publicly accessible address


![image](https://drek4537l1klr.cloudfront.net/luksa3/v-12/Figures/03image014.png)


```sh
k get pods -o wide

# observing requests hitting all three pods when using the service
k get svc
  NAME         TYPE           CLUSTER-IP     EXTERNAL-IP   PORT(S)          AGE
  kiada-http   LoadBalancer   10.96.59.164   <pending>     8080:32232/TCP   43h
  kubernetes   ClusterIP      10.96.0.1      <none>        443/TCP          2d

k describe pods | grep Node:
Node:         kind-worker2/172.18.0.3
Node:         kind-worker/172.18.0.4
Node:         kind-worker/172.18.0.4

# Each request arrives at a different pod in random order.
# This is what services in Kubernetes do when more than one pod instance is behind them.
# They act as load balancers in front of the pods.

curl 172.18.0.3:32232
curl 172.18.0.4:32232

```




