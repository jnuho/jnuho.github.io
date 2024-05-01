
```
안녕하세요. 이준호라고 합니다.
저는 3년 정도 개발 경험이 있고, 현재는 미들웨어 운영업무를 하고 있습니다.
	(Why 운영업무: 개발하면서 인프라를 다뤄볼수 있는 기회가 있었고,
	운영 업무에 관심이 생겨서 지원하게 되었습니다.)

로웸이라는 회사에서 2년정도 백엔드 개발을 하면서
- 스프링기반 백엔드 API 개발,
- 데이터베이스 테이블 인덱스 설계,
- 배포 업무 및 AWS 인프라 업무를 했습니다.

현재는 씨에스리 운영팀 미들웨어 운영 업무를 하고 있습니다.
- 주로 AWS EC2 내에 미들웨어와 ALB, TG,
- 그리고 컨테이너 서비스를 ECS, ECR을 통해 생성 및 운영 하고 있습니다.

```

- Network
	- [network](https://github.com/WeareSoft/tech-interview/blob/master/contents/network.md)
	- [tcpip](https://madplay.github.io/post/network-tcp-udp-tcpip)
	- [http https](https://github.com/WeareSoft/tech-interview/blob/master/contents/network.md#http%EC%99%80-https)
	- [http https 동작](https://github.com/WeareSoft/tech-interview/blob/master/contents/network.md#http%EC%99%80-https-%EB%8F%99%EC%9E%91-%EA%B3%BC%EC%A0%95)
	- [cookie & session](https://github.com/WeareSoft/tech-interview/blob/master/contents/network.md#%EC%BF%A0%ED%82%A4%EC%99%80-%EC%84%B8%EC%85%98)
	- [rest](https://github.com/WeareSoft/tech-interview/blob/master/contents/network.md#rest%EC%99%80-restful%EC%9D%98-%EA%B0%9C%EB%85%90)
	- [socket](https://github.com/WeareSoft/tech-interview/blob/master/contents/network.md#socketio%EC%99%80-websocket%EC%9D%98-%EC%B0%A8%EC%9D%B4)
	- [async sync](https://github.com/WeareSoft/tech-interview/blob/master/contents/etc.md#blocking-non-blocking-vs-synchronous-asynchronous)
	- [load balancer vs reverse proxy](https://www.nginx.com/resources/glossary/reverse-proxy-vs-load-balancer/)
	- [authorization - JWT](https://inpa.tistory.com/entry/WEB-%F0%9F%93%9A-JWTjson-web-token-%EB%9E%80-%F0%9F%92%AF-%EC%A0%95%EB%A6%AC)


```
Class Struct 차이

부하테스트 성능이슈해결

http가 무엇인지
https

NLB ALB

REST
	GET POST

mybatis
JPA
ORM

slice
go channel

interceptor, aop, filter

GC

java out of memory - oome
https://www.nextree.co.kr/p3878/
```

- [HTTP](https://medium.com/@rlatla626/http-%EC%A0%95%EB%A6%AC-3958d2a82312)
	- Hypertext Text Transfer Protocol
		- How web clients and servers communicate
		- Where resources (web content) come from
		- How web transactions work
		- The format of the messages used for HTTP communication
		- The underlying TCP network transport
		- The different variations of the HTTP protocol
		- Some of the many HTTP architectural components installed around the Internet
			- We’ve got a lot of ground to cover, so let’s get started on our tour of HTTP.


- 서버와 클라이언트가 HTTP 프로토콜을 통해 통신
	- HTTP 메시지 = Header + body
		- Header
			- General 헤더: 요청과 응답에 공통으로 들어가는 header (Date, Connection, Via, Transfer-Encoding)
			- Request 헤더: Host, User-Agent, Accept-Encoding, Authorization
			- Response 헤더: Server, Age, Set-Cookie, Allow
			- Entity 헤더:  HTTP 메시지의 body에 관한 정보
				- Content-Encoding, Content-length, Content-Type, Content-Location, Expires, Last-Modified, ETag
	- HTTP 연결과정
		- HTTP 프로토콜 전송과정 (TCP/IP)
			- 4계층 Application Layer
			- 3계층 Transport Layer
			- 2계층 Internet Layer
			- 1계층 Network Interface Layer
	- HTTP 요청
		- 브라우저 URL입력
		- DNS서버로 전달 : URL을 통해 서버 IP를 알려줌
		- DNS서버에 해당 URL에 대한 IP가 없으면 다른 DNS서버를 추가로 거쳐서 IP정보 받음
		- 해당 IP주소로 HTTP 요청 시도
			- TCP/IP 기반 이라서 IP:port (기본 80)로 요청
		- 신뢰성 보장을 위해 3번의 패킷 교환 과정 - 3 way handshake 
		- 클라이언트의 요청 시도
		- 서버에서 응답 제공
		- 서버와 연결을 끊기 위해 4번의 패킷 교환과정 - 4 way handshake 
		- 브라우저에 결과 출력
	- HTTP 요청의 특징: stateless

```
```

