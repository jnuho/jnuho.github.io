
```
안녕하세요. 이준호라고 합니다.
저는 3년 정도 개발 경험이 있고, 현재는 미들웨어 운영업무를 하고 있습니다.
	(Why 운영업무: 개발하면서 인프라를 다뤄볼수 있는 기회가 있었고,
	운영 업무에 관심이 생겨서 지원하게 되었습니다.)

로웸이라는 회사에서 2년정도 백엔드 개발을 했구요
- 스프링기반 백엔드 API 개발, 데이터베이스 테이블 인덱스 설계, 배포 업무, AWS 인프라 운영 등 업무를 했습니다.

현재는 씨에스리 운영팀 미들웨어 운영 업무를 하고 있습니다.
- 주로 AWS EC2 내에 미들웨어와 ALB, TG, 그리고 컨테이너 서비스를 ECS, ECR을 통해 생성 및 운영 하고 있습니다.

여기까지 입니다. 감사합니다.

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
개발하면서 이미 작성 되어있는 코드를 읽고 변경하는 과정이 많았는데
이때 코드 리팩토링과 깔끔한 코드를 짜는 것의 중요성을 깨달았습니다.

따라서 이 원칙을 중요시 생각하여 작성하는 편입니다.
팀원들이 모두 보는 코드이기 때문에, 필요하다면 커멘트도 달고 있지만,
커멘트 없이도 이해하기 쉬운 코드를 짜는것이 목표입니다.

개발 API에 대한 테스트 케이스를 작성하고,
maven build 시 spring 설정에 따라 테스트 케이스 검증 과정
그리고 팀원들과의 코드리뷰 및 브랜치 머지과정
```
