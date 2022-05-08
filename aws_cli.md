

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
{
    "TargetGroups": [
        {
            "TargetGroupName": "my-targets",
            "Protocol": "HTTP",
            "Port": 80,
            "VpcId": "vpc-3ac0fb,
            "TargetType": "instance",
            ...
            "LoadBalancerArns: [
              "", "", ..
            ]
        }
        ,...
   ]
}

```


- aws
  - elbv2
    - describe-target-health

```sh
TG=arn:aws:elasticloadbalancing:ap-northeast-2:734976340835:targetgroup/AWSDC-TG-COM-PRD-KESOA-7080/68261da5b61413c1

aws elbv2 describe-target-health \
  --profile APPPRD \
  --query 'TargetHealthDescriptions[?TargetHealth.State==`healthy`].{Id:Target.Id, Port: Target.Port}' \
  --output text \
  --target-group-arn ${TG} \
  | /mnt/c/Windows/System32/clip.exe



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
        ...
}


```
