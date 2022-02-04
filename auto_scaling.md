
## Auto Scaling

- [auto scaling life cycle](https://docs.aws.amazon.com/autoscaling/ec2/userguide/AutoScalingGroupLifecycle.html)


###  Scale in

다음 scale-in 이벤트들은 Auto Scaling Group 그룹이 EC2 인스턴스를 그룹으로 부터 detach 하고 종료 함:
- 수동으로 그룹 사이즈 축소
- 스케일링 정책 생성하여, 자동으로 그룹 사이즈 요청한 만큼 축소
- 특정 시간에 그룹의 사이즈를 축소 시키는 스케일링을 스케쥴 함

### Scale out

다음 scale-out 이벤트들은 Auto Scaling Group 그룹이 EC2 인스턴스를 실행 하고 그룹에 attach 함
- 수동으로 그룹 사이즈 증가
- 스케일링 정책 생성하여, 자동으로 그룹 사이즈 요청한 만큼 증가
- 특정 시간에 그룹의 사이즈를 증가 시키는 스케일링을 스케쥴 함
- 서버 인스턴스 개수를 늘림


#### Scale Up
- 새로 추가하는 것이 아닌, 기존에 있는 기기의 성능을 up


#### Elastic Load Balancer (ELB)
- point for receiving incoming traffic from users that evenly distributes the traffic
- distribute traffic to other instances in case of failure of one or more instances
- Dynamic Scaling : based on traffic scale, it scales up and and down (number of instances)
- Three types of Load Balancer
  - Application Load Balancer
  - Network Load Balancer
  - Clasic Load Balancer
- Components of the Load Balancer
  - Listeners
    - at least 1 listener is required
    - defines how your inbound connections are routed to your target groups based on ports and protocols set as conditions
  - Target Groups
    - group of your resources that you want your ELB to route request to
    - such as a fleet of EC2 instances
    - can configure ELB with different target groups., associated with different listener configuration
  - Rules
    - rules for listeners configured in your ELB. defines how request gets routed to target group