
- 같은 VPC 공유:
  - ElastiCache 클러스터
  - EC2 starpass-was-00
  - EC2 starpass-bastion

<!-- ![VPC with public and private subnets](https://docs.aws.amazon.com/vpc/latest/userguide/images/nat-gateway-diagram.png) -->


<!-- ![elasticache_and_ec2](https://docs.aws.amazon.com/AmazonElastiCache/latest/mem-ug/images/ElastiCache-inVPC-AccessedByEC2-SameVPC.png) -->

![diagram](./assets/images/ElastiCache.jpg)

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
  - inbound: SSH 오피스ip:22

- EC2-private(was)
  - with a new SG with inbound port 22 for SSH
  - inbound: sg_bastion (SSH 22, TCP 8080), starpass-internal-elb-was(TCP 8080)

- ElastiCache용
  - inbound: sg_was/web/bastion (TCP 6379)

#### EC2
- 퍼블릭 IP자동할당 : 'Enable'
- 다음 커멘드 라인으로 redis클라이언트 설치

```sh
# Linux in general
wget http://download.redis.io/redis-stable.tar.gz
tar -xvzf redis-stable.tar.gz
cd redis-stable
make && make install
redis-cli -h {ElastiCache_Endpoint} -p {port configured in SG above;6379}

# Ubuntu
sudo apt update
sudo apt install redis

> flushall
> keys *
```


#### ElastiCache-redis 생성
- 서브넷그룹은 starpass-was-00, starpass-bastion 과 같은 VPC, 프라이빗 서브넷 선택
```
vpc-0b4163a5f741002f8 (starpass-vpc) 
subnet-01ab087db1ecc6748 (starpass-private-was-a) 
sg-03ceb4c49e904f0aa (starpass-redis)
```

#### 테스트
- 위의 VPC, 퍼블릭서브넷으로 생성된 EC2에서만 ElasticCache접근 확인
- 다른 VPC 또는 프라이빗서브넷으로 생성된 EC2에서는 접근 불가
