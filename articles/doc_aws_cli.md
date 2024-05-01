- [aws](#)
  - [elbv2](#)
    - [describe-target-groups](#aws-elbv2-describe-target-groups)
    - [describe-target-health](#aws-elbv2-describe-target-health)
    - [describe-rules](#aws-elbv2-describe-rules)

  - [autoscaling](#)
    - [describe-auto-scaling-groups](#aws-autoscaling-describe-auto-scaling-groups)
  - [ecr](#)
    - [list-images](#aws-ecr-list-images)
  - [ecr-public](#)
    - [describe-image-tags](#aws-ecr-public-describe-image-tags)
  - [ecs](#)
    - [execute-command](#aws-ecs-execute-command)
    - [describe-clusters](#aws-ecs-describe-clusters)
    - [describe-services](#aws-ecs-describe-services)
    - [describe-task-definition](#aws-ecs-describe-task-definition)
    - [describe-container-instances](#aws-ecs-describe-container-instances)
  - [ec2](#)
    - [describe-instances](#aws-ec2-describe-instances)

### aws elbv2 describe-target-groups

```sh
# 타겟그룹 조회
#   로드밸런서 연결된 타겟그룹 여부 설정 가능
# => 조회 결과 csv파일로 저장
aws elbv2 describe-target-groups \
--profile dev \
--query 'TargetGroups[?LoadBalancerArns[0] != None].{TargetType: TargetType, TG: TargetGroupName, TGARN: TargetGroupArn, ALB: LoadBalancerArns[0], Port: Port}' \
--output text \
| sed -E  's/\s+/,/g' > tg-dev.csv
```

```json
{
    "TargetGroups": [
        {
            "TargetGroupName": "my-targets",
            "Protocol": "HTTP",
            "Port": 80,
            "VpcId": "vpc-3ac0fb5f",
            "TargetType": "instance",
            "HealthCheckEnabled": true,
            "UnhealthyThresholdCount": 2,
            "HealthyThresholdCount": 5,
            "HealthCheckPath": "/",
            "Matcher": {
                "HttpCode": "200"
            },
            "HealthCheckProtocol": "HTTP",
            "HealthCheckPort": "traffic-port",
            "HealthCheckIntervalSeconds": 30,
            "HealthCheckTimeoutSeconds": 5,
            "TargetGroupArn": "arn:aws:elasticloadbalancing:us-west-2:123456789012:targetgroup/my-targets/73e2d6bc24d8a067",
            "LoadBalancerArns": [
                "arn:aws:elasticloadbalancing:us-west-2:123456789012:loadbalancer/app/my-load-balancer/50dc6c495c0c9188"
            ]
        }
    ]
}
```

### aws elbv2 describe-target-health

```sh
# 위의 타겟그룹 ARN 리스트를 파라미터로 설정하는 스크립트
TG=arn:aws:elasticloadbalancing:...

aws elbv2 describe-target-health \
  --profile dev \
  --output text \
  --target-group-arn ${TG}

aws elbv2 describe-target-health \
  --profile prd \
  --query 'TargetHealthDescriptions[?TargetHealth.State==`healthy`].{Id:Target.Id, Port: Target.Port}' \
  --output text \
  --target-group-arn ${TG} \
  | /mnt/c/Windows/System32/clip.exe


TG=arn:aws:elasticloadbalancing:us-east-...

aws elbv2 describe-target-health \
  --profile DR \
  --query 'TargetHealthDescriptions[?TargetHealth.State!=`healthy`].{Id:Target.Id, Port: Target.Port}' \
  --output text \
  --target-group-arn ${TG} \
  > tg-dr-1.txt



aws ec2 describe-tags --profile prd \
--filters "Name=resource-id,Values=$AWS_INSTANCE_ID" "Name=key,Values=Name" \
--output text | cut -f5 \
| /mnt/c/Windows/System32/clip.exe
```

```json
{
    "TargetHealthDescriptions": [
        {
            "HealthCheckPort": "80",
            "Target": {
                "Id": "i-xxx",
                "Port": 80
            },
            "TargetHealth": {
                "State": "healthy"
            }
        },
        {
            "HealthCheckPort": "80",
            "Target": {
                "Id": "i-xxx",
                "Port": 80
            },
            "TargetHealth": {
                "State": "healthy"
            }
        }
    ]
}
```


### aws autoscaling describe-auto-scaling-groups

```sh

aws autoscaling describe-auto-scaling-groups \
  --profile prd \
  --query 'AutoScalingGroups[?Instances[0] != None]'
```

- aws elbv2 describe-load-balancers --profile prd

```sh
aws elbv2 describe-load-balancers
  --query 'LoadBalancers[?VpdId==`vpc-0de9d17f8134e676d`]' \
  --profile prd

```

### aws ecs execute-command

```sh
aws ecs execute-command --cluster ccc --task xxx --container containerrname --interactive --command "/bin/sh" --profile stg

aws lightsail get-load-balancer-metric-data \
  --load-balancer-name ALBNAME \
  --metric-name RequestCount \
  --period 3600 \
  --start-time 1652054606 \
  --end-time 1652227188 \
  --unit Milliseconds \
  --statistics Average \
  --profile prd
```

### aws ec2 describe-instances

```sh
aws ec2 describe-instances \
  --instance-ids "i-abcd" "i-efgh" \
  --profile DR
```

### aws ecr list-image

```sh
aws ecr list-images \
  --repository-name reponame \
  --profile stg
```

### aws ecr-public describe-image-tags

```sh
aws ecr-public describe-image-tags \
  --repository-name test\
  --profile dev
```

### aws ecr describe-images

```sh
aws ecr describe-images \
  --repository-name reponame \
  --image-ids imageTag=latest \
  --profile stg
```



### aws iam get-role

```sh
aws iam list-attached-role-policies \
  --role-name ecs-jims \
  --profile stg


```

### aws ecs describe-task-definition

```sh
aws ecs describe-task-definition \
  --task-definition task \
  --query 'taskDefinition.{name:family, taskDefinitionArn: taskDefinitionArn, cpu: cpu, memory: memory, taskRole: taskRoleArn, container: containerDefinitions[*].{ name : name, cpu: cpu, memoryReservation:  memoryReservation, portMappings: portMappings[*].{ container: containerPort, host: hostPort }, environment: environment[*], logConfiguration: logConfiguration } }' \
  --profile dev


ALB 생성
- intmernal 체크
- IPV4
- SG, subnet-a,c 설정

# Task Definition
aws ecs list-tags-for-resource \
  --resource-arn "arn:aws:e..." \
  --profile dev \
  --output l

# Service
aws ecs list-tags-for-resource \
  --resource-arn "arn:aws:ecs:" \
  --profile dev \
  --output yaml

# Cluster
aws ecs list-tags-for-resource \
  --resource-arn "arn:aws:ecs:ap" \
  --profile dev \
  --output yaml

```

### aws ecs describe-servis

```sh
aws ecs describe-services \
  --cluster abc
  --services abcd \
  --query 'services[*].{name:serviceName, serviceArn: serviceArn, loadBalancer: loadBalancers[*].{ targetGroupArn: targetGroupArn, container: containerName, port: containerPort }, taskSets: taskSets[*].{id:id, taskSetArn:taskSetArn, startedBy: startedBy, networkConfiguration: networkConfiguration[*], loadBalancers: loadBalancers[*]}, networkConfiguration : networkConfiguration, healthCheckGracePeriodSeconds: healthCheckGracePeriodSeconds }' \
  --profile prd

aws ec2 describe-subnets \
  --query 'Subnets[?SubnetId==`subnet-07f9cb604134ec4e9`].{Vpc: VpcId, Tags: (Tags[] | [?Key==`Name`].Value)}' \
  --profile stg

aws ec2 describe-security-groups \
  --group-ids "sg-0f5adea7d7a182bac" \
  --query 'SecurityGroups[].GroupName' \
  --output text \
  --profile dev

aws ecs list-tags-for-resource \
  --resource-arn "arn:..." \
  --profile dev
```

### aws ecs describe-clusters

```sh
aws ecs describe-clusters \
  --clusters abc \
  --profile dev
```




### aws elbv2 describe-rules

```sh

PROFILE=dev
ALB=aaa
ALB_ARN=$(aws elbv2 describe-load-balancers \
  --name $ALB \
  --query 'LoadBalancers[].LoadBalancerArn' \
  --profile $PROFILE \
  --output text)


LISTENER_ARN=$(aws elbv2 describe-listeners \
  --load-balancer-arn $ALB_ARN \
  --query 'Listeners[?Port==`7080`].ListenerArn' \
  --output text \
  --profile $PROFILE)

aws elbv2 describe-rules \
  --listener-arn $LISTENER_ARN \
  --profile $PROFILE
  # --query 'Rules[?Conditions[?Values[0]==`jimsdev.koreanair.com`]].{ host: Conditions[?Field==`host-header`].Values[0] ,TargetGroup: Actions[0].TargetGroupArn}' \



```

### aws ecs describe-container-instances

```sh
```


### aws ecs tag-resources
```sh
# The Amazon Resource Name (ARN)of the resource to add tags to.
# Currently, the supported resources are
#   Amazon ECS capacity providers
#   tasks
#   services
#   task definitions
#   clusters
#   container instances.
aws ecs tag-resource \
  --resource-arn arn:... \
  --tags key=key1,value=value1 key=key2,value=value2 key=key3,value=value3
  --tags key=testKey,value=testValue \
  --profile dev
```


### aws elasticache add-tags-to-resource

```sh
aws elasticache add-tags-to-resource \
  --resource-name arn... \
  --tags Key="Department",Value="dept" Key="Owner",Value="ooo" \
  --profile dev
```



aws ecs describe-tasks \
  --cluster ccc \
  --tasks arn:aws... \
  --profile prd


aws ecs list-tasks \
--cluster ccc \
--service-name sss \
--profile prd

aws ecs describe-tasks \
--cluster ccc \
--tasks 'arn:...' \
--profile prd


aws ecr get-lifecycle-policy \
  --repository-name rponame \
  --output yaml \
  --profile prd
