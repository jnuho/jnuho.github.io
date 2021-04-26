

## 로드밸런싱 이란?

- https://www.nginx.com/resources/glossary/load-balancing/
- 백엔드 서버 그룹(또는 서버팜, 서버풀)의 인바운드 네트워크 트래픽을 효과적으로 분산 하는 것을 가리킴
- 다량의 트래픽이 발생하는 모던 웹사이트는 수십만건의 concurrent 클라이언트 요청을 처리해야 하며, 동시에 빠르고 안정성 있게 텍스트, 이미지, 비디오, 어플리케이션 데이터 등을 제공해야 함. 이런 다량의 요청을 처리하기 위해 모던 컴퓨팅은 서버의 수를 늘리는 경우가 일반적.

## 역할

- 로드밸런서는 서버에 오는 클라이언트 요청들을 가장 빠르고 공간 효율적으로 통제하는 서버의 traffic cop역할을 하며, 서버 과부하로 인한 성능 저하가 오지 않도록 함.
 - 하나의 서버가 다운 되면, 로드밸런서는 살아있는 다른 서버들로 트래픽을 리다이렉트 시킴.
- 서버그룹에 서버가 새로 추가되면, 로드밸런서는 자동으로 이 새로운 서버에 요청을 보내기 시작함.

- 클라이언트 요청이나 네트워크 로드를 효과적으로 다수의 서버에 걸쳐 분산 시킴
- 사용중인 서버로만 요청을 보내서, 높은 사용성(availability) 과 안정성(reliability)을 보장
- 요청양에 따라 서버를 추가하거나 없앨 수 있는 flexibility 제공
- 스케일업(cpu 업그레이드), 스케일아웃(서버 증설; 로드밸런싱 필요)

## 다양한 로드밸런싱 알고리즘

- Round Robin – Requests are distributed across the group of servers sequentially.
- Least Connections – A new request is sent to the server with the fewest current connections to clients. The relative computing capacity of each server is factored into determining which one has the least connections.
- Least Time – Sends requests to the server selected by a formula that combines the fastest response time and fewest active connections. Exclusive to NGINX Plus.
- Hash – Distributes requests based on a key you define, such as the client IP address or the request URL. NGINX Plus can optionally apply a consistent hash to minimize redistribution of loads if the set of upstream servers changes.
- IP Hash – The IP address of the client is used to determine which server receives the request.
- Random with Two Choices – Picks two servers at random and sends the request to the
one that is selected by then applying the Least Connections algorithm (or for NGINX Plus
the Least Time algorithm, if so configured).

## 로드밸런싱의 장점

- Reduced downtime
- Scalable
- Redundancy
- Flexibility
- Efficiency


## ELB (Elastic Load Balancing)

- Cloudwatch metrics
- RDS/EC2 CPUUtilization
    - rds CPU 사용%
    - was-asg CPU 사용%
- ALB WAS/WEB
    - ActiveConnectionCount The total number of concurrent TCP connections active from clients to the load balancer and from the load balancer to targets.
    - RequestCount The number of requests processed over IPv4 and IPv6.
    - RequestCountPerTarget The average number of requests received by each target in a target group. You must specify the target group using the TargetGroup dimension. This metric does not apply if the target is a Lambda function.
- ALB Response Time    
