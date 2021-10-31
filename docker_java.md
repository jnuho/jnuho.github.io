
### Network

From the book 'Docker and Kubernetes for Java'

- Docker network types
  - Bridge (default)
  - Host
    - fast network speed
    - for performance
    - vulnerable security
    - accessible through the host's ip address
    - requires port mapping to reach services inside the container
  - None
    - does not configure networking at all
    - networking is disabled

- Networking commands

```sh
docker network ls
docker network create {network_name}
docker network rm {network_name}
docker network connect {network_name}
docker network disconnect {network_name}

# docker network create 커멘드 옵션
# -d, --driver="bridge"
docker network create myNetwork
docker network create -d bridge myNetwork
docker network create --driver=bridge myNetwork

docker network create -d myNetwork
docker network inspect myNetwork
docker run -it --net=myNetwork tomcat

# 'myPostgreSQL' container can use the same network as the other container, 'myTomcat', use
# they can directly communite each other in the same network, 'bridge'
# though the network itself isolates the containers from external networks
# [Docker Host] -> [ISOLATED Network(container1, 2, 3, ...)]
docker run -it --net=bridge myTomcat
docker run -it --net=container:myTomcat myPostgreSQL
```

- Exposing and Mapping ports

```Dockerfile
# Official 'tomcat' image Dockerfile fragment FROM openjdk:8-jre-alpine
ENV CATALINA_HOME /usr/local/tomcat
ENV PATH $CATALINA_HOME/bin:$PATH
RUN mkdir -p "$CATALINA_HOME"
WORKDIR "$CATALINA_HOME"
EXPOSE 8080
CMD ["catalina.sh", "run"]
```

- [Add containers to a network](https://docs.docker.com/engine/tutorials/networkingcontainers/#add-containers-to-a-network)

```sh
docker network create --driver=bridge myNetwork
docker run -it --name myTomcat --net=myNetwork tomcat
docker run -it --net=container:myTomcat busybox
```

- `-p`
  - `{hostPort}:{containerPort}`
  - Specify a port mapping rule, mapping the port on the container with the port on the host machine
  - Makes a port open from the outside of Docker
  - Bind a port(or group of ports) from a host to the container, using `docker run -p` command
  - same as `publish`
  - create a port mapping rule, mapping a port on the container with the port on the host system
  - the mapped port will be available from outside Docker
  - if you do `-p`, but there is no `EXPOSE` in the Dockerfile, docker will do an implicit `EXPOSE`

- `EXPOSE` signals that there is service available on the speicifed port. Used in the Dockerfile and makes exposed ports open for other containers

- `--expose` same as `EXPOSE` but used in the runtime, during the container startup
  - expose ports at runtime, but will not create any mapping to the host
  - exposed ports will be available only to another container running on the same network, on the same Docker host
  - There is no way to create a port mapping in the Dockerfile. Mapping ports `-p` is just a runtime option

- `-P`
  - Map dynamically allocated random host ports to all container ports that have been exposed using `EXPOSE` or `--expose`

```sh
# docker run -p {hostPort}:{containerPort} {image ID or name}
# docker run -expose=7000-8000 {container ID or name}
# docker run -p 7000-8000:7000-8000 {container ID or name}
# TEST by http://localhost:8080
docker run -it --name myTomcat2 --net=myNetwork -p 8080:8080 tomcat
```

```sh
# docker will map a random port on the host to Tomcat's exposed port number 8080
docker run -it --name myTomcat3 --net=myNetwork -P tomcat
# http://localhost:32772
docker ps
docker port myTomcat3
docker inspect myTomcat3
```


### Volumns

- Persistant storage
  - separate the container life cycle and your application from the data
  - Data generated(or being used) by your application is not destroyed or tied to the container life cycle and can thus be reused
