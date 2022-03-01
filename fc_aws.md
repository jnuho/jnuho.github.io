

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
    - 생략가능. 생략 될 경우 커멘드에 지정된 명령어로 수행
  - 커멘드 : 도커가 실행 할 떄 수행할 명령어, 혹은 엔트리포인트에 지정된 명령어에 대한 인자 값
    - 도커파일에 필수로 들어가야함

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
docker run \
  -d \
  --name fastcampus-mysql \
  -e MYSQL_DATABASE=fastcampus \ # 환경변수 지정: 최초생성 데이터베이스 이름
  -e MYSQL_ROOT_PASSWORD=fastcampus \ # 환경변수 지정: 루트계정의 비밀번호
  -v db:/var/lib/mysql \  # 호스트 볼륨대신 도커볼륨 사용
  -p 3306:3306 \
  mysql:5.7

docker ps
docker volume inspect db
sudo ls -l /var/lib/docker/volumes/db/_data

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


- 이미지 빌드