
❒ From the book "도커, 컨테이너 빌드업!"

- 도커엔진 컨테이너 기술은 리눅스 자체기술에 기반
	- chroot, namespace, cgroup을 조합한 LnuX Container, LXC


- nodejs 앱 테스트

```sh
# 이미지 다운로드
docker pull node

# 컨테이너화
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
#   우분투,CENTOS /var/lib/docker
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
```

- docker pull
	- docker hub에서 로컬로 다운하면서 생성된 레이어 distribution id 출력
	- 이미지 변경 되면, 로컬 저장된 이미지를 제외한 새로운 레이어 계층들만 로컬에 다운로드



```sh
docker tag 원본이미지:[태그] 참조이미지:[태그]

docker images
	REPO		TAG				IMAGE ID
	httpd		latest		b2c2ec
docker tag b2c2ec debian-httpd:1.0

docker images
	REPO					TAG				IMAGE ID
	debin-httpd		1.0				b2c2ec
	httpd					latest		b2c2ec

docker login
	username:
	password:

docker push [본인아이디]/httpd:3.0

docker pull [본인아이디]/httpd:3.0

```



- 도커 이미지를 파일로 관리

```sh
```











