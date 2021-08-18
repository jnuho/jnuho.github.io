
- 인프라 OverView
  - 현재 같은 VPC 내에, 퍼블릭/프라이빗 서브넷으로 구분되어 있음
  - ElastiCache 클러스터는 미생성 상태이지만, 보안그룹은 생성되어 있음 (인바운드 포트6379)
  - 클러스터 생성 시 VPC, 서브넷그룹 선택필요
    - VPC: EC2와 같은 VPC 선택
    - [서브넷그룹](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/SubnetGroups.html): 프라이빗 서브넷 선택 

<!-- ![VPC with public and private subnets](https://docs.aws.amazon.com/vpc/latest/userguide/images/nat-gateway-diagram.png) -->


<!-- ![elasticache_and_ec2](https://docs.aws.amazon.com/AmazonElastiCache/latest/mem-ug/images/ElastiCache-inVPC-AccessedByEC2-SameVPC.png) -->

<!-- ![diagram](./assets/images/ElastiCache.jpg) -->
![diagram](https://d2cg24p20j4o18.cloudfront.net/playvote/000/20210819/82331f92-bc8c-403e-a1d1-5d51bc6fec79.jpg)

- 인프라 구축 테스트 과정 - [참고링크](https://github.com/ROWEM-Development/dev-infra-info/blob/main/jenkins/01-jenkins.md)
  1. [VPC](#vpc)
  2. [인터넷 게이트웨이](#인터넷-게이트웨이)
  3. [서브넷, 라우팅 테이블](#서브넷-라우팅-테이블)
  4. [보안 그룹](#보안-그룹)
  5. [EC2](#ec2)
  5. [ElastiCache](#elasticache)

#### VPC
- CIDR 172.20.0.0/16
- e.g. 172.20.0.0 ~ 172.20.255.255


#### 인터넷 게이트웨이
- VPC에 Attach

#### 서브넷, 라우팅 테이블
- 퍼블릭
  - 서브넷 : CIDR 172.20.0.0/24 (현재는 VPC 내 로컬에서만 접근 가능-private)
  - 라우트테이블
  - '라우트' 탭 > 추가 '0.0.0.0/0', 'igw-071ed7e5e7b2b8385'
  - '서브넷연결' 탭 > 위의 서브넷 연결하여 퍼블릭으로 만듦

- 프라이빗
  - 서브넷 : CIDR 172.20.10.0/23
  - 라우트테이블
    - '서브넷연결' 탭 > 위의 서브넷 연결하여 프라이빗으로 만듦

#### 보안 그룹
같은 VPC내 보안 그룹 구분

- EC2-public(bastion)
  - 인바운드: SSH 오피스ip:22

- EC2-private(was)
  - with a new SG with inbound port 22 for SSH
  - 인바운드: sg_bastion (SSH 22, TCP 8080), starpass-internal-elb-was(TCP 8080)

- ElastiCache용
  - 인바운드: sg_was/web/bastion (TCP 6379)

#### EC2
- 퍼블릭 IP자동할당 : 'Enable'
- redis클라이언트 설치

#### ElastiCache
- redis 선택
- 서브넷그룹은 starpass-was-00, starpass-bastion 과 같은 VPC, 프라이빗 서브넷 선택
```
vpc-0b4163a5f741002f8 (starpass-vpc) 
subnet-01ab087db1ecc6748 (starpass-private-was-a) 
sg-03ceb4c49e904f0aa (starpass-redis)
```

#### 테스트
- 위의 VPC, 프라이빗 서브넷으로 생성된 EC2(starpass-was)에서만 ElasticCache접근 확인
- 보안그룹 인바운드 규칙에 정의된 호스트 이외 ip에서는 접근 불가 확인
- 스프링 환경에서 jedis를 접근가능 한지 여부 확인하기 위해, 해당 EC2에서 redis-cli로 접속 테스트

```sh
# Linux in general
wget http://download.redis.io/redis-stable.tar.gz
tar -xvzf redis-stable.tar.gz
cd redis-stable
make && make install

# Ubuntu
sudo apt update
sudo apt install redis

# ElastiCache 접속 시도
redis-cli -h {ElastiCache 엔드포인트} -p {보안그룹에 정의된 포트 6379}

> flushall
> keys *
```
