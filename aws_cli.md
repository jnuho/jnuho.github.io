
### AWS CLI
 
- aws
  - elbv2
    - describe-target-groups`

```sh
aws elbv2 describe-target-groups \
--profile APPPRD \
--query 'TargetGroups[?LoadBalancerArns[0] != None].{TargetType: TargetType, TG: TargetGroupName, TGARN: TargetGroupArn, ALB: LoadBalancerArns[0], Port: Port}' \
--output text \
| sed -E  's/\s+/,/g' > tg.csv

```


- aws
  - elbv2
    - describe-target-health

```sh
aws elbv2 describe-target-health \
  --profile APPPRD \
  --query 'TargetHealthDesciptions[?TargetHealth.State==`healthy`].{Id:Target.Id, Port: Target.Port}' \
  --target-group-arn \
    arn:aws:elasticloadbalancing:ap-northeast-2:734976340835:targetgroup/AWSDC-TG-COM-PRD-KESOA-7080/68261da5b61413c1


{
    "TargetHealthDescriptions": [
        {
            "Target": {
                "Id": "10.48.46.154",
                "Port": 7080,
                "AvailabilityZone": "ap-northeast-2a"
            },
            "HealthCheckPort": "7080",
            "TargetHealth": {
                "State": "healthy"
            }
        },
        {
            "Target": {
                "Id": "10.48.146.154",
                "Port": 7080,
                "AvailabilityZone": "ap-northeast-2c"
            },
            "HealthCheckPort": "7080",
            "TargetHealth": {
                "State": "unhealthy",
                "Reason": "Target.FailedHealthChecks",
                "Description": "Health checks failed"
            }
        }
    ]
}


aws elbv2 describe-target-health \
  --profile APPPRD \
  --target-group-arn \
    arn:aws:elasticloadbalancing:ap-northeast-2:734976340835:targetgroup/AWSDC-TG-COM-PRD-KESOA-7080/68261da5b61413c1 \
  --query 'TargetHealthDescriptions[*].{ID:Target.Id, PORT:Target.Port}' \
  --output text \
  | sed -E 's/\s+/,/g'

```
