


















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

