


## 도커

- docker run
  - pull > create > start > attach(-it 옵션 사용 시)
- docker create
  - pull > create

```shell
docker -v
docker images
docker run 

# 컨테이너 centos 생성 후 시작
docker create -it --name mycentos centos:7
docker start mycentos

# 도커 내부로 들어감
# CTRL + P 누른상태로 Q : 컨테이너 중단없이 외부로 나옴
# exit 커멘드는 컨테이너를 중단시킴
docker attach mycentos

# 현재 실행 중인 mycentos 컨테이너의 Id 추출
docker inspect mycentos | grep Id

# CONTAINER_ID :
# IMAGE: 생성된 원본 이미지 이름
# COMMAND : 컨테이너가 시작될 때 실행될 명령어
#  centos의 경우 /bin/bash 이미 내장되어 있어서 별도 지정없음
# PORTS: 컨테이너가 개방한 포트와, 호스트에 연결한 포트
#  외부에 노출하지 않을 떄는 항목내용 없음
# NAMES
docker ps
docker run -it ubuntu:14:04 echo hello world!
docker rename mycentos mynewcentos

docker stop NAME
docker rm NAME

docker rm -f NAME

# remove all stopped containers
docker container prune

# -a : 중지여부 상관없이 모든 컨테이너
# -q : ID만출력
docker ps -a -q

docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)
```

```shell
docker start myubuntu
docker exec -it myubuntu bash

# ehh0: 도커의 NAT IP 할당받은 인터페이스
# lo: 인터페이스
ifconfig

# 
```