- Go AWS Sdk
  - https://aws.amazon.com/sdk-for-go/
  - 주요기능 : 
    - ALB/NLB -> TargetGroup -> ECS/EC2C 자원 조회
    - ECS 자원 상세조회 : 각 Task내의 Container 및 이미지, IP 정보 상세조회
    - ECR 이미지 : Tag 및 이미지 uri 조회 최신 순 조회
  - 구현 시 고려사항 :
    - AWS 자원간에 의존성이 있는데, API에서 조회 결과로 복잡한 의존관계를 한번의 호출로 조회할 수 없음
    - 의존 관계에 있는 각 자원 API 호출하여 Go언어로 데이터를 가공하여 결과 얻음
  - 사용언어 :
    - 백엔드: Go, 프론트엔드: Javascript, HTML, CSS

![go sdk app](Animation.gif)
