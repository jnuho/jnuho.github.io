
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