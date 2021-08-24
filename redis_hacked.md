
- 레디스 기본포트 오픈 후 Unknown Traffic
  - /var/tmp/.system에 Monero마이닝 스크립트 및 실행 프로세스 탐지
  - 아웃바운드 트래픽 발생

![1](./assets/images/hacked1.jpg)
![2](./assets/images/hacked2.jpg)
![3](./assets/images/hacked3.jpg)

- 처리
  - 해당 프로세스 kill 및 /var/tmp/.system 디렉토리 삭제
  - 실행 중 레디스 프로세스 kill
  - 방화벽 레디스 기본포트 6379 closed처리

- 추후 방안
  - 기본포트 6379 사용자제
  - 방화벽 포트 오픈시 특정 IP로 Source IP 제한 (e.g. 스타패스 서버 IP)
