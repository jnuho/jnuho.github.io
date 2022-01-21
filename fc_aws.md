





















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
```sh
sudo apt update
sudo apt install unzip build-essential curl
```

  - [cli 설치](https://docs.aws.amazon.com/ko_kr/cli/latest/userguide/install-cliv2-linux.html)
```sh
# 특정버전
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64-2.2.16.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install
```

- AWS CLI자격증명 설정 우선순위
  1. CLI명령어 옵션
  2. 환경변수
  3. CLI 자격증명 파일 - `~/.aws/credentials`
  4. CLI 설정 파일 - `~/.aws/config`
  5. 컨테이너 자격증명 (ECS의 경우)
  6. 인스턴스 프로파일 자격증명 (EC2의 경우)

```sh
# 액세스키 다운로드 후
cat mycredential.csv

vim ~/.aws/config
[default]
aws_access_key_id=AKI...K
aws_secret_access_key_id=k0r...M8

# 테스트
aws sts get-caller-identity
aws ec2 describe-key-pairs

# 리전설정 필요
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

# 프로파일 AWS_PROFILE 환경 변수 또는 --profile 또는 config파일
[profile eu-west-1]
region=eu-west-1
aws_access_key_id=AKI...K
aws_secret_access_key_id=k0r...M8

[profile ap-northeast-1]
region=ap-northeast-1
aws_access_key_id=AKI...K
aws_secret_access_key_id=k0r...M8

# 디폴트 프로파일의 리전
aws configure get region
# eu-west-1 프로파일의 리전
aws configure get region --profile eu-west-1

# 환경변수로 디폴트 리전 설정
export AWS_PROFILE=ap-northeast-1
aws configure get region
ap-northeast-1
```


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