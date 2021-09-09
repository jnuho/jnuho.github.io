
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

- sg-bastion
  - 인바운드: SSH 오피스ip:22

- sg-starpass-was
  - 인바운드: sg-bastion (SSH 22, TCP 8080)

- sg-starpass-redis
  - 인바운드: sg-was/web/bastion (TCP 7379)


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
