
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
docker start container-test1
docker ps

# 컨테이너에 접속
docker attach container-test1
	exit
docker rm container-test1


# docker run
docker run -it --name container-test1 ubuntu:14.04 bash
	exit
docker rm container-test1	

```




