
❒ From the book "도커, 컨테이너 빌드업!"

- 도커엔진 컨테이너 기술은 리눅스 자체기술에 기반
	- chroot, namespace, cgroup을 조합한 LnuX Container, LXC


- nodejs 앱 테스트

```sh
# 이미지 다운로드
docker pull node

# 컨테이너화
# 	--publish [host_port]:[container_port]
docker run -dp 8080:80 -it --name=nodejs_test node:latest

# 컨테이너 실행 확인
docker ps

# 소스코드 로컬->컨테이너 내부
cat > nodejs_test.js
	var http = require('http');
	http.createServer(function (req, res) {
		res.writeHead(200, {'Content-Type': 'text/plain'});
		res.end('Hello World!');
	}).listen(80);

docker cp nodejs_test.js nodejs_test:/nodejs_test.js

# 실행중인 npm이 설치된 nodejs_test 컨테이너에 bash shell로 접속
docker exec -it nodejs_test /bin/bash

>	ls
	nodejs_test.js
> node -v
> node nodejs_test.js &
> exit

curl localhost:8080
	Hello World!
```

```sh
# docker run --publish [host_port]:[container_port]
# docker run -dp 8080:80 docker/getting-started:pwd
# 호스트 PC에서 localhost:8080으로 접근
```

- Kubernetes: Container orchestration tool
	- 다수의 컨테이너 관리
	- 컨테이너 간 네트워킹 문제 해결
	- 컨테이너 인스턴스 확장

- Kubernetes Managed Service
	- EKS
	- GKE
	- AKS


```sh
cat > Dockerfile
	FROM ubuntu:18.04
	MAINTAINER RABOOF <raboof@email.com>
	RUN \
				apt-get update && \
				apt-get install -y apache2

	RUN sed -i 's/80/7080/g' /etc/apache2/ports.conf
	EXPOSE 7080
	CMD ["apachectl", "-D", "FOREGROUND"]
```

```sh
# 호스트가 아닌, 컨테이너 서비스를 통해 명령어 실행
docker run -it busybox sh
docker run busybox echo 'Hello World'

```


```
docker run명령어 [클라이언트]
	-> 서버의 도커데몬
	-> docker.sock의 도커 API로 컨테이너 생성
	-> 수행된 컨테이너에 포함된 서비스 결과를 [클라이언트]에 전달
```
	

```sh
docker -v
docker version

docker info
docker system info
docker info --format '{{json .}}'

# 회수 가능한 공간: RECLAIMABLE
docker system df
docker system df -v

# 회수 가능공간 확보
docker system prune

# 실시간 로그 출력
# 터미널 두개로 확인
docker system events
docker run -itd -p 80:80 --name=webapp nginx
docker ps
docker stop webapp

docker system events
docker system events --filter 'type=image'
docker system events --filter 'event=stop'
docker system events --filter 'container=webapp'
docker system events --filter 'event=stop' --filter 'container=webapp'

docker system events --since 24h

docker system events --format '{{json .}}'


# 도커데몬 디버깅을위한
# OS별 도커데몬 로그 위치
#   우분투, CENTOS /var/lib/docker
#   RHEL /var/log/messages
#   데비안 /var/log/message.log
#   윈도우 ~AppData\Local
#   macOs..

# 로그내용 중 msg 키워드 정보가 상세 로그 내용임
sudo journalctl -u docker.service

# 유사한 도커 데몬 디버그
sudo dockerd -D


# 리눅스 시스템에도 기본적인  로그 수집 데몬이 있음
syslogd
rsyslogd
# [참고] 도커로그를 호스트 운영체제 로그수집 데몬에 연결해서 로그 기록 할 수도 있음
```

- docker login, pull, push

```sh
# 도커 허브
# tag생략 시 latest
docker pull imagename:tag
docker pull imagename:digest
docker pull gcr.io/google-samples/hello-app:1.0
docker push
docker login
docker logout

docker image ls
docker image inspect imagename --format="{{ .RepoTags}}" httpd
docker image inspect imagename --format="{{ .RepoTags}}" httpd
docker image inspect imagename --format="{{ .RepoTags}}" httpd


# docker login 이후
docker search httpd

docker image history [OPTIONS] IMAGE
docker image history httpd

# docker pull
# docker hub에서 로컬로 다운하면서 생성된 레이어 distribution id 출력
# 이미지 변경 되면, 로컬 저장된 이미지를 제외한 새로운 레이어 계층들만 로컬에 다운로드
docker pull httpd:latest


docker tag 원본이미지:[태그] 참조이미지:[태그]

docker images
	REPO		TAG				IMAGE ID
	httpd		latest		b2c2ec
docker tag b2c2ec debian-httpd:1.0

docker images
	REPO					TAG				IMAGE ID
	debin-httpd		1.0				b2c2ec
	httpd					latest		b2c2ec

docker push [본인아이디]/httpd:3.0
docker pull [본인아이디]/httpd:3.0

# 베이스 이미지에 특정 애플리케이션을 서비스와 코드를
# 포함해 컨테이너로 포함해 실행하는 경우 docker commit으로
# 컨테이너를 이미지로 저장가능
```


- docker login 방법
	1. username/password
	2. access token

```sh
# 1.
docker login
	username:
	password:

cat ~/.docker/config.json
	{
		"auths": {
			"https://index.docker.io/v1/": {
				"auth": "...."
			}
		}
	}
docker info | grep username

# config.json 파일의 auth 값이 삭제되고, docker info에서도 사용자명이 제거됨
docker logout

# 2.
# 	hub.docker.com > Account settings > Security > New Access Token
# 	복사된 액세스 토큰 사용하여 로그인
vi .access_token

cat .access_token | docker login --username jnuho --password-stdin
```


- 도커 이미지를 파일로 관리
	- hub.docker.com 에서 이미지 내려받아 내부망으로 이전하는 경우
	- 신규 애플리케이션 서비스위해 Dockerfile로 새롭게 생성한 이미지를 저장 및 배포해야 하는 경우
	- Container를 commit하여 생성한 이미지를 저장 및 배포해야 하는 경우
	- 개발 및 수정한 이미지 등

```sh
docker image save [옵션] <파일명> [image명]
docker image load [옵션]

# 이미지-> .tar 저장
docker pull mysql:5.7
docker images
docker image save mysql:5.7 > test-mysql57.tar
ls -lh test-mysql57.tar

# 묶인 파일내용 확인
tar tvf test-mysql57.tar

docker image rm mysql:5.7
docker images

# 이미지<- .tar 로드: 방법1
docker image load < test-mysql57.tar
docker images

# 이미지<- .tar 로드: 방법2
cat test-mysql57.tar | docker import - mysql57:1.0
docker images


# 이미지 용량 줄이기 : gzip
docker image save mysql:5.7 | gzip > test-mysql57zip.tar.gz
ls -lh test-mysql57zip.tar.gz
docker image rm mysql:5.7
docker images
docker image load < test-mysql57zip.tar.gz
docker images
```

- 도커 이미지 삭제

```sh
docker image rm [option] {이미지이름[:태그] | 이미지ID}
docker rmi {옵션} {이미지이름[:태그] | 이미지 ID}

docker pull ubuntu:14.04
docker image rm ubuntu:14.04

docker image rm b2c2a
docker image rm -f b2c2a

# show images ids
docker images -q

# remove all images
docker rmi $(docker images -q)

docker rmi $(docker images | grep debian)
docker rmi $(docker images | grep -v centos)

# 상태가 exited 인 container를 찾아서 모두 삭제
vi .bashrc
	alias cexrm='docker rm $(docker ps --filter 'status=exited' -a -q)'
source .bashrc
alias

# -a옵션으로 사용 중이 아닌 모든 이미지 제거
docker image prune -a
docker image prune -a -f --filter "until=48h"


# 도커 컨테이너의 PID 1번 프로세스도 init 프로세스?
# 호스트의 셸 프로세스 ID
echo $$
	32370

docker run -it centos:8 bash
> echo $$
	1

# 다른 터미널에서 실행 중인 PID 조회
ps -ef | grep 32370
```

- 컨테이너 실행

```sh
# docker run과 달리, container 내부접근 없이 생성(스냅샷)만 수행
docker crerate -it --name container-test1 ubuntu:14.04
docker ps -a
# status : created -> up
docker start container-test1
docker ps

# 컨테이너에 접속 - 실행중인 어플리케이션 컨테이너에 단순한 조회 작업 수행
docker attach container-test1
	exit
docker rm container-test1


# docker run = [pull] + create + start + [command]
docker run -it --name container-test1 ubuntu:14.04 bash
	exit
docker rm container-test1	

docker run -it --name container-test1 ubuntu:14.04 hostname

```

- docker run 옵션
	- https://docs.docker.com/engine/reference/run/


- 실습1. Mysql 컨테이너 실행 후 database생성

```sh
docker pull mysql:5.7
docker images | grep mysql

docker run -it mysql:5.7 /bin/bash
> cat /etc/os-release
> /etc/init.d/mysql start
> mysql -uroot
mysql> show databases;
mysql> create database dockerdb;
mysql> show databases;
> cd /var/lib/mysql
> ls
... dockerdb...

docker ps -a
docker start ba5a
docker exec -it ba5a bash

# 컨테이너 종료시키지 않고 나가려면 ctrl+p+q
docker inspect flamboyant_nash | grep IPAddress
docker exec -it flamboyant_nash bash

# 다른이름으로 새로운 컨테이너 실행
docker run -it mysql:5.7 bash
```


- 컨테이너 모니터링 도구 cAdvisor (google) 컨테이너 실행

```sh
# http://hostip:9559 에서 볼수 있음
docker run \
	--volume=/:/rootfs:ro \
	--volume=/var/run:/var/run:rw \
	--volume=/sys:/sys:ro \
	--volume=/var/lib/docker/:/var/lib/docker:ro \
	--publish=9559:8080 \
	--detach=true \
	--name=cadvisor \
	google/cadvisor:latest

docker ps
	CONTAINER ID   IMAGE                    COMMAND                  CREATED         STATUS         PORTS
	NAMES
	1c653c540628   google/cadvisor:latest   "/usr/bin/cadvisor -…"   8 seconds ago   Up 7 seconds   0.0.0.0:9559->8080/tcp, :::9559->8080/tcp   cadvisor

```

- 실습2. 웹서비스 실행 nginx 컨테이너

```sh
docker pull nginx:1.8
docker images
docker run --name webserver1 -d -p 8001:80 nginx:1.18
docker ps
sudo netstat -nlp | grep 8001
curl localhost:8001
docker stats webserver1
docker top webserver1

# -f: 실시간, -t 마지막로그
docker logs -f webserver1
docker stop webserver1
docker start webserver1

vi index.html
	<h1> Hello Jpub Docker. </h1>
docker cp index.html webserver1:/usr/share/nginx/html/index.html
curl localhost:8001

# 프로세스 일시중단
docker pause webserver1
docker ps
docker unpause webserver1
docker ps
docker restart webserver1
```





- 실습3. 파이썬 프로그래밍 환경 구축 - 컨테이너 내부에서 python 실행

```sh
cat > py_lotto.py

from random import shuffle
from time import sleep
gamenum = input('로또 게임 횟수 입력:')
for i in range(int(gamenum)):
	balls = [x+1 for x in range(45)]
	ret = []
	for j in range(6):
		shuffle(balls)
		number = balls.pop()
		ret.append(number)
	ret.sort()
	print('로또번호[%d]: ' %(i+1), end='')
	print(ret)
	sleep(1)

# 파이썬 컨테이너 실행 후 py_lotto.py 코드복사
docker run -it -d --name=python_test -p 8900:8900 python
docker cp py_lotto.py python_test:/

# 파이썬 컨테이너에 설치된 모듈 확인
docker exec -it python_test bash
> python --version
> pip list
> python -c 'help("modules")'

docker exec -it python_test python /py_lotto.py
```


- 실습4. Nginx 환경 conf 로컬에서 수정후 다시 컨테이너에 복사


```sh
docker run -d -p 8010:80 --name=webserver10 nginx:latest
# 컨테이너 -> 로컬로 복사
docker cp webserver10:/etc/nginx/nginx.conf ./nginx.conf
# 변경후 다시 로컬->컨테이너로 복사
docker cp nginx.conf webserver10:/etc/nginx/nginx.conf
docker restart webserver10
```

- 실습5. nodjs 테스트 환경을 위한 컨테이너 실행

```sh
cat > nodejs_test.js
var http = require('http');
var content = function(req, res) {
	resp.end("Good morning Korea~!" + "\n");
	resp.writeHead(200);
}
var web = http.createServer(content);
web.listen(8002);

# Host OS에서 실행시 nodejs 미설치로 에러발생!
node nodejs_test.js

# 컨테이너에서 실행 가능
docker pull node
docker run -d -it -p 8002:8002 --name=nodejs_test node
docker ps

docker cp nodejs_test.js nodejs_test:/nodejs_test.js

# 컨테이너 내부에서 nodejs_test.js를 실행!
docker exec -it nodejs_test node /nodejs_test.js

curl localhost:8002

# 노드 서버  실행 중인 컨테이너 이름변경하고 싶다면 docker rename
# 동적으로 변경됨!
# docker rename <기존 Container명> <새로운 Container명>
docker renmae nodejs_test nodeapp
docker ps


# nodejs_test.js를 실행하는 컨테이너를 다시 이미지로 저장 가능
```


- 실습6. base 이미지로 실행 후 변경한  컨테이너를 이미지로 저장


```sh
docker run -it --name=webserver8 -d -p 8008:80 nginx:latest
docker ps
curl localhost:8008

cat > index.html
<h1> Hello World! </h1>
docker cp index.html webserver8:/usr/share/nginx/html/index.html
# 추가된 웹소스 변경 정보 확인
# 	A 추가, D 삭제, C 변경
docker diff webserver8
	C /usr/share/nginx/html/index.html

# 컨테이너 -> 이미지
docker commit -a "jpub" webserver8 webfront:1.0

docker images | grep webfront
# 이미지 -> 도커허브
docker login

# 본인ID/이미지:태그
docker tag webfront:1.0 jnuho/webfront:1.0
docker push jnuho/webfront:1.0

docker run -it --name webserver9 -d -p 8009:80 jnuho/webfront:1.0
curl localhost:8009
```



- 도커 볼륨 활용
	- 도커가 유니언 파일시스템: 이미지-> 여러 컨테이너 프로세스
	- 서비스의 데이터와 로직 분리 해야 함
	- 도커 볼륨은 컨테이너에서 생성, 재사용 가능, 호스트 OS에서 접근 가능
	- 일반적으로 컨테이너 Lifecycle에 따라 컨테이너 중지시 삭제 됨
	- 도커 볼륨은 독립적으로 운영 되므로 컨테이너 중지시에도 유지 됨


- 도커 볼륨 타입

1. volume
	- 볼륨 생성:  `docker volume create 볼륨명`
	- 도커 명령어로 볼륨 관리
	- 여러 컨테이너 간에 안전하게 공유 가능
	- 볼륨 드라이버를 통해 원격 호스트 및 클라우드 환경에 볼륨 내용 저장 및 암호화 가능
	- 새 볼륨으로 지정될 영역에 데이터 미리 채우고 컨테이너 연결 하면, 컨테이너에서 바로 데이터 사용 가능

```sh
# 볼륨 생성
docker volume create my-appvol-1
docker volume ls

docker volume inspect my-appvol-1

# --mount 옵션으로 볼륨 지정
docker run -d --name vol-test1 \
	--mount source=my-appvol-1,target=/app \
	ubuntu:20.04

# -v 옵션으로 볼륨 지정
docker run -d --name vol-test2 \
	-v my-appvol-1:/var/log \
	ubuntu:20.04

# 볼륨 미리 만들지 않아도 자동 생성
docker run -d --name vol-test3 \
	-v my-appvol-2:/var/log \
	ubuntu:20.04

docker volume ls
DRIVER	VOLUME NAME
local		my-appvol-1
local		my-appvol-2

docker inspect vol-test1
	"Mounts": [{}]...
docker inspect --format="{{ .Mounts}}" vol-test1
	[{}]

# 자원 삭제: 컨테이너 제거-> 볼륨 삭제
docker volume rm my-appvol-1
docker stop vol-test1 vol-test2
docker rm vol-test1 vol-test2
docker volume rm my-appvol-1
```


2. bind mount
	- 도커 볼륨 기법에 비해 사용 제한적
	- 호스트 파일 시스템 절대경로: 컨테이너 내부 경로를 직접 마운트 하여 사용
	- 사용자가 파일,디렉토리 생성시 해당 호스트 파일시스템의 소유자 권한으로 연결.
		- 존재하지 않는경우 자동생성됨. 이 자동생성 디렉토리는 루트 사용자 소유가 됨
	- 컨테이너 실행시 지정하여 사용하고, 컨테이너 제거 시,  바인드 마운트는 해제되지만, 호스트 디렉터리는 유지


```sh
# --mount 옵션으로 사전에 생성한 경로와 바인드 마운트 지정.
mkdir /home/foo/target
docker run -d -it --name bind-test1 \
	--mount type=bind,source=$(pwd)"/target,target=/var/log \
	centos:8


# -v 옵션으로 사전에 생성한 경로와 바인드 마운트 지정.
docker run -d -it --name bind-test2 \
	-v "$(pwd)"/target:/var/log \
	centos:8
...

```

3. tmpfs mount
	- bind mount 방법은 컨테이너 중지후 데이터 유지하지만
		- tmpfs 마운트 방법은 임시적이며, 호스트 메모리에서만 지속되므로
		- 컨테이너가 중지되면 tmpfs 마운트가 제거되고 내부에 기록된 파일은 유지되지 않는다.


```sh
```

- 도커 볼륨 활용1: 데이터베이스의 데이터 지속성 유지
	- 컨테이너 장애로 중단되어도, 새로운 컨테이너에 동일볼륨 연결시 DB, Table, Data 모두 동일하게 지속가능

```sh
# 볼륨 생성
docker volume create mysql-data-vol
docker volume ls

# mysql 컨테이너 실행 시 데이터베이스 설정도 추가하여 실행
# 	hub.docker.com의 mysql페이지에 환경변수 설명 참조
docker run -it --name=mysql-vtest \
	-e MYSQL_ROOT_PASSWORD=jhlee \
	-e MYSQL_DATABASE=dockertest \
	-v mysql-data-vol:/var/lib/mysql -d \
	mysql:5.7

docker exec -it mysql-vtest bash

/etc/init.d/mysql start
mysql -uroot -p
show databases;
use dockertest;
create table mytab (c1 int, c2 char);
insert into mytab values (1, 'a');
select * from mytab;
exit
ls /var/lib/mysql/dockertest/
exit

docker inspect --format="{{ .Mounts }}" mysql-vtest
sudo ls -l /var/lib/docker/volumes/mysql-data-vol/_data/dockertest

# 컨테이너 장애 테스트 케이스: 컨테이너 정지,제거 후 동일 볼륨지정 컨테이너 실행 시
# 	기존 데이터 그대로 유지 확인
docker stop mysql-vtest
docker rm mysql-vtest
docker run -it --name=mysql-vtest \
	-e MYSQL_ROOT_PASSWORD=jhlee \
	-e MYSQL_DATABASE=dockertest \
	-v mysql-data-vol:/var/lib/mysql -d \
	mysql:5.7

```


- 도커 볼륨 활용2: 웹서비스의 로그 정보 보호 및 분석을 위한 바인드 마운트 설정
	- access.log통해 장애 상황 정보 및 실시간 접근 로그 확인
	- awk를 이용하여 로그 분석 가능


```sh
mkdir nginx-log
docker run -d -p 8011:80 \
	-v /home/jhlee/nginx-log:/var/log/nginx \
	nginx:1.19


http://localhost:8011
tail -f nginx-log/access.log

# 지정된 날짜/시간 내에 기록된 IP 정보를 내림차순으로 카운트하여 횟수와 IP 주소 출력.(예, 01/Mar/2021:12:32:37)
awk '$4>"[로그에 기록된 날짜, 시간]" && $4<"[로그에 기록된 날짜, 시간]"' access.log | awk '{ print $1}' | sort | uniq -c | sort -r | more
```


- 도커 볼륨 활용3: 컨테이너 간 데이터 공유를 위한 데이터 컨테이너 만들기
