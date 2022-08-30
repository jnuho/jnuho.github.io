
- From the book "Docker Container Buildup"


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