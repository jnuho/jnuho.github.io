
- From the book "도커, 컨테이너 빌드업!"

- 도커엔진 컨테이너 기술은 리눅스 자체기술에 기반
	- chroot, namespace, cgroup을 조합한 LnuX Container, LXC


- nodejs 앱 테스트

```sh
# 이미지 다운로드
docker pull node

# 컨테이너화
docker run -d -it --name=nodejs_test node:latest

# 컨테이너 실행 확인
docker ps

# 소스코드 로컬->컨테이너 내부
docker cp nodejs_test.js nodejs_test:/nodejs_test.js

# 실행중인 npm이 설치된 nodejs_test 컨테이너에 bash shell로 접속
docker exec -it nodejs_test /bin/bash

>	ls
	nodejs_test.js
> node -v
# 샘플코드 테스트 실행
> node nodejs_test.js
```



```sh
docker ps -a

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
journalctl -u docker.service

# 유사한 도커 데몬 디버그
dockerd -D


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

