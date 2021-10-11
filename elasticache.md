
- 인프라 OverView
  - 같은 VPC 내에, 퍼블릭/프라이빗 서브넷으로 구분
  - 캐시 클러스터 미생성 상태. 보안그룹은 생성되어 있음(인바운드 6379)
  - 클러스터 생성 시 VPC, 서브넷그룹 선택
    - VPC: EC2와 같은 VPC 선택
    - [서브넷그룹](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/SubnetGroups.html): 프라이빗 서브넷 선택 

<!-- ![VPC with public and private subnets](https://docs.aws.amazon.com/vpc/latest/userguide/images/nat-gateway-diagram.png) -->

<!-- ![elasticache_and_ec2](https://docs.aws.amazon.com/AmazonElastiCache/latest/mem-ug/images/ElastiCache-inVPC-AccessedByEC2-SameVPC.png) -->

![diagram](https://d2cg24p20j4o18.cloudfront.net/playvote/000/20210819/82331f92-bc8c-403e-a1d1-5d51bc6fec79.jpg)

- 인프라 구축 테스트
  1. [VPC](#vpc)
  2. [서브넷](#서브넷)
  3. [인터넷 게이트웨이](#인터넷-게이트웨이)
  4. [라우트 테이블](#라우트-테이블)
  5. [NACL](#NACL)
  6. [EC2](#ec2)
  7. [보안 그룹](#보안-그룹)
  8. [NAT Gateway](#nat-gateway)
  9. [VPC Endpoint](#vpc-endpoint)
  10. [ElastiCache](#elasticache)

#### VPC
- CIDR 10.0.0.0/16
- e.g. 10.0.0.0 ~ 10.0.255.255
- VPC 생성 후 자동생성 : RT, NACL, SG

#### 서브넷
- 퍼블릭
  - 서브넷 : CIDR 10.0.0.0/24 (현재는 VPC 내 로컬에서만 접근 가능-private)
    - 작업 > 자동 할당 IP 설정 수정정보 > 퍼블릭 IPv4 주소 자동 할당 활성화

- 프라이빗
  - 서브넷 : CIDR 10.0.1.0/24

#### 인터넷 게이트웨이
- 생성 후 VPC에 연결 (Attach)
```sh
aws ec2 attach-internet-gateway --vpc-id "vpc-0e02ef59133d3b5a7" --internet-gateway-id "igw-0a871cc54507807f4" --region ap-northeast-2
```


#### 라우트 테이블
- 퍼블릭 rt 생성 후
  - 작업 > '서브넷연결' > 퍼플릭 서브넷 연결
  - '라우팅' > 편집 > 추가 '0.0.0.0/0','igw-071ed7e5e7b2b8385'
    - 디폴트 10.0.0.0/16 범위 이외 ip 요청을 igw로 연결

- 프라이빗 rt (VPC 생성시 디폴트 생성 됨)
  - '서브넷연결' > 프라이빗 서브넷 연결 (명시연결 필요 X)
    - 명시적 연결이 없는 서브넷: 다음 서브넷은 어떤 라우팅 테이블과도 명시적으로 연결되어 있지 않고 기본 라우팅 테이블에 연결되어 있는 상태

#### NACL

인스턴스 보안 설정 방법
- NACL - stateless
- 보안 그룹 - stateful
- 프라이빗 nacl (VPC 생성시 디폴트 생성 됨) > 서브넷 연결
- 퍼블릭 nacl > 서브넷 연결
  - 인바운드 22번포트는 Allow
  - 규칙 낮은순 우선순위, 만약 99규칙 22번포트 Deny이면 Deny

규칙 번호 | 유형 | 프로토콜 | 포트 범위 | 소스 | 허용/거부
---|---|---|---|---|---
100	|SSH(22)		|TCP(6)		|22		|0.0.0.0/0	| Allow
101	|SSH(22)	|TCP(6)	|22	|0.0.0.0/0	 |Deny
200	|HTTP(80)	|TCP(6)|	80	|0.0.0.0/0	 |Allow
300	| HTTPS(443)	|TCP(6)	|443	|0.0.0.0/0	 |Allow
*	| 모든 트래픽	|모두|	모두|	0.0.0.0/0	| Deny

- 인바운드 22만 Allow 하는 케이스 

규칙 번호 | 유형 | 프로토콜 | 포트 범위 | 소스 | 허용/거부
---|---|---|---|---|---
100	|SSH(22)		|TCP(6)		|22		|0.0.0.0/0	| Allow
101	|모든 TCP	|TCP(6)	|모두(0-65535)	|0.0.0.0/0	 |Deny
200	|HTTP(80)	|TCP(6)|	80	|0.0.0.0/0	 |Allow
300	| HTTPS(443)	|TCP(6)	|443	|0.0.0.0/0	 |Allow
*	| 모든 트래픽	|모두|	모두|	0.0.0.0/0	| Deny

- 아웃바운드 1024-65535 Allow


#### EC2
바스티온 EC2인스턴스 생성 - 퍼블릭 서브넷 지정

- 단계 3: 인스턴스 세부정보 구성
  - 고급세부정보 > 사용자데이터
  ```sh
  #!/bin/bash
  yum install httpd -y
  service httpd start
  ```

#### 보안 그룹
같은 VPC내 보안 그룹 구분

- sg-bastion
  - 인바운드: SSH 오피스ip:22

- sg-starpass-was
  - 인바운드: sg-bastion (SSH 22, TCP 8080)

- sg-starpass-redis
  - 인바운드: sg-was/web/bastion (TCP 7379)

#### NAT Gateway

- 프라이빗 EC2에서 외부 인터넷 막혀있어서, yum으로 소프트웨어 (mysql, apache,...) 설치 안됨
- (VPC) NAT Gateway 만들어 외부와 연결: 퍼블릭 서브넷 내에 생성
- 프라이빗 RT수정 : 라우팅 > 추가 0.0.0.0/0  (대상: nat-...)


#### VPC Endpoint
- 프라이빗 EC2 에서 S3접근하고 싶을때 NAT게이트웨이 대신 VPC Endpoint 사용: 보안상 이유
- IAM 역할 > 추가 > EC2 > 'AmazonS3FullAccess' > 'test-01-s3-fullaccess'
- 프라이빗 EC2 생성 시 IAM역할 선택 또는 [ATTACH IAM역할.](https://aws.amazon.com/blogs/security/easily-replace-or-attach-an-iam-role-to-an-existing-ec2-instance-by-using-the-ec2-console/)

- S3 연결 시도
```sh
# S3버킷 리스트 조회 (VPC Endpoint설정도 안했는데 연결됨)
# 프라이빗 RT에 NAT설정 되어있기때문
aws s3 ls --region ap-northeast-2
# 프라이빗 RT에 NAT설정 삭제후 시도시 조회안됨
aws s3 ls --region ap-northeast-2
```
- VPC Endpoint 설정 이후 S3 연결 시도
  - VPC > 엔드포인트 > 생성
    - 'S3'검색 > com.amazonaws.ap-northeast-2.s3 서비스선택, Gateway유형 선택
    - VPC 및 프라이빗 서브넷 선택
  - 프라이빗 RT > 라우팅 > 대상에 VPC Endpoint 추가됨 확인
    - S3관련된 트래픽을 S3로 보내는 정책
    - VPC Endpoint 밑에 라우팅 추가 > NAT Gateway추가하여 위의 두개 IP이외는 NAT Gateway로 보내도록 함
```sh
# 조회가능 확인
aws s3 ls --region ap-northeast-2
```

#### 운영환경 - ElastiCache
- redis 선택
- 노드유형: r5,m5,r4,m4,r3,m3,t3,t2 메모리 및 네트워크 성능 선택
- 서브넷그룹 '생성' starpass-was-00, starpass-bastion와 같은 VPC, 프라이빗 서브넷 선택
- 보안그룹 sg-starpass-redis 선택 
- 자동 백업 활성화 체크해제

```
vpc-0b4163a5f741002f8 (starpass-vpc) 
subnet-01ab087db1ecc6748 (starpass-private-was-a) 
sg-03ceb4c49e904f0aa (starpass-redis)
```

#### 개발환경 - 레디스서버 (EC2/아이넷호스트)

- redis 서버 설치
- yum install 또는 make install으로 루트 시스템에 설치하지 않고, 다음 방법 선택
- starpass 유저 로그인 후, redis 소스 다운로드, redis.conf 설정 변경
    - 로컬 에서만 redis 서버 접근 가능하도록 bind 127.0.0.1
    - requirepass, port 설정 추가

```sh
# 소스 다운로드
$ su starpass
$ cd ~starpass/
$ wget http://download.redis.io/redis-stable.tar.gz
$ tar -xvzf redis-stable.tar.gz
$ vim redis-stable/redis.conf
# bind 127.0.0.1 로컬(톰캣)에서만 접근
# requirepass [your_password]
# port [your_port]
$ cd redis-stable/src
$ ./redis-server /home/starpass/redis-stable/redis.conf &

# 접속 시도
redis-cli -h {ElastiCache 엔드포인트} -p {보안그룹에 정의된 포트 7379}

```

```sh
> flushall
> keys *
```

#### EC2
- redis서버 설정 수정
  - 서버(EC2또는 아이넷호스트 머신)에 설치된 redis서버를 외부에서 접근할 수 있도록 설정 변경
    - EC2: 보안그룹의 인바운드규칙 TCP 7379 127.0.0.1
    - 아이넷호스트: redis.conf 수정

- 퍼블릭 IP자동할당 : 'Enable'
- yum install 또는 make install으로 루트 시스템에 설치하지 않고, 다음 방법 선택
- starpass 유저로 소스만 받아, redis.conf 설정 추가:
    - 로컬 에서만 redis 접근 가능하도록 bind 제한
    - requirepass, port 설정 추가
- redis-server 스크립트 redis.conf 설정 반영하여 실행

```sh
su starpass
cd ~starpass/
wget http://download.redis.io/redis-stable.tar.gz
tar -xvzf redis-stable.tar.gz
vim redis-stable/redis.conf
# bind 127.0.0.1 로컬(톰캣)에서만 접근
# requirepass [your_password]
# port [your_port]
cd redis-stable/src
./redis-server /home/starpass/redis-stable/redis.conf &

# 접속 시도
redis-cli -h {ElastiCache 엔드포인트} -p {보안그룹에 정의된 포트 7379}

> flushall
> keys *
```

```java
import io.lettuce.core.RedisClient;
import io.lettuce.core.api.StatefulRedisConnection;
import io.lettuce.core.api.async.RedisAsyncCommands;

public class LettuceConnection {

	// String.format("redis://%s:%d/0", hostname, port)
	private static final String REDIS_CON_URL = "redis://13.209.76.95:6379/0"; // 준호EC2 redis
//	private static final String REDIS_CON_URL = "elasticache-junho-0813.eo7tpf.0001.apn2.cache.amazonaws.com:6379";

	public static void main(String[] args) {
		RedisClient redisClient = RedisClient.create(REDIS_CON_URL);
		StatefulRedisConnection<String, String> connection = redisClient.connect();
		RedisAsyncCommands<String, String> async = connection.async();

		final String[] result = new String[1];

		async.set("foo", "bar")
				.thenComposeAsync(ok -> async.get("foo"))
				.thenAccept(s -> result[0] = s)
				.toCompletableFuture()
				.join();

		connection.close();
		redisClient.shutdown();

		System.out.println(result[0]); // "bar"
	}
}
```

- 서버의 redis클라이언트로 데이터 확인
```sh
redis-cli
keys *
```

#### 테스트
- ElastiCache접근은 보안그룹에 인바운드 규칙에 정의된 EC2(private subnet) 이외 ip에서는 접근 불가
- 개발/운영 환경 분리하여 캐시 관리
  - 개발: 개발서버에 redis-server 설치하여 캐시저장
  - 운영: ElastiCache 클러스터 생성하여 EC2->ElastiCache
