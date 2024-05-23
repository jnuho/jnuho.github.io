❚ AWS 자원 조회
- [Eng Ver.](doc_goproject)
- [repository](https://github.com/jnuho/goproject)
- [aws sdk for go](https://aws.amazon.com/sdk-for-go/)

<div>
<img src="../assets/images/goproject.gif" height="60%" width="60%" alt="go sdk app">
</div>

❚ 진행 배경
- EC2 -> ECS 이관 및 생성 작업에 필요한, ECS, TG, ALB 등 자원조회 기능 필요
- 자원 생성 시 stage 별 데이터 확인 필요하여 AWS Console 대신, SDK 활용

❚ 사용 기술

```
Backend : Go, Go AWS SDK
Frontend : Javascript, HTML, CSS
```

❚ 주요 기능
- 조회 : ALB > TargetGroup > Target Health >  Container_Ip, Instance_Id
- 조회 : ECS Task의 Container, Image, IP 등
- 조회 : ECR tag, image uri -> 최신순 정렬

❚ 구현 시 고려사항 :
- CLI (go cobra library) 툴 대신, 웹페이지 url (ip:port)로 간편하게 사용할 수 있는 web 어플리케이션 생성
	- url:  프라이빗 클라우드 내부 vpn으로만 접근 가능
- AWS 자원간 의존성이 있는데, AWS가 제공하는 단일 API로 원하는 자원상태 조회가 힘듦
- AWS API 호출 결과들로 object list 만들어, 정렬 처리 (Override Len, Less, Swap functions)

❚ 진행 상태
- Web application (go, javascript, html, css) 완성
- 컨테이너화
	- Dockerfile 빌드, Image 생성, 컨테이너 Run
	- EC2 role based access로, aws credential 관리 -> 도커 컨테이너로 credential propagation 테스트 중
	- 고루틴 및 채널 적용: API 호출시 다수의 AWS api 통신하는 경우
- Cloudwatch 로그, EC2, Task event 조회(down, up) 추가 예정

❚ 테스트
- Cloud접속 -> Vpn > http://{ip_addr}:port

![go sdk app](../assets/images/goproject_kor.png)


❚ Aws Profile 관리
- `credential_source=Ec2InstanceMetadata`
	- AWS CLI / SDK가 EC2 인스턴스에 attach된 IAM Role 사용하여 Source credential 가져옴

- EC2인스턴스에 IAM Role Attach하기

```
# IAM Role with permission policies :
#		Create role > Add Permissions

AmazonEC2FullAccess
	Provides full access to Amazon EC2 via the AWS Management Console.

IAMFullAccess
	Provides full access to IAM via the AWS Management Console.

AmazonEC2ContainerRegistryFullAccess
	Provides administrative access to Amazon ECR resources

AmazonS3FullAccess
	Provides full access to all buckets via the AWS Management Console.

ReadOnlyAccess
	Provides read-only access to AWS services and resources.

AmazonSESFullAccess
	Provides full access to Amazon SES via the AWS Management Console.

AmazonAPIGatewayAdministrator
	Provides full access to create/edit/delete APIs in Amazon API Gateway via the AWS Management Console.

AmazonECS_FullAccess
	Provides administrative access to Amazon ECS resources and enables ECS features through access to other AWS service resources, including VPCs, Auto Scaling groups, and CloudFormation stacks.

AWSCloudFormationFullAccess
	Provides full access to AWS CloudFormation.

AWSLambda_FullAccess
	Grants full access to AWS Lambda service, AWS
```

❚ Credential 접근: go aws sdk

```go
import "github.com/aws/aws-sdk-go/aws/session"

// AWS 프로파일 명 (~/.aws/config)
type AwsProfile struct{
	dev string
	stg string
	prd string
}
var awsProfile AwsProfile

// 세션객체 초기화
func InitSession(profile string) *session.Session {
	if profile == "dev" {
		profile = awsProfile.dev
	} else if profile == "stg" {
		profile = awsProfile.stg
	} else if profile == "prd" {
		profile = awsProfile.prd
	}
	sess, err := session.NewSessionWithOptions(session.Options{
		// Specify profile to load for the session's config
		Profile: profile,
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		panic(err)
	}
	return sess
}
```

❚ 구조체 Repo 정의<br>
❚ Repo 반환 인터페이스 정의<br>
❚ Repo 구조체 구현 메소드 정의
- `func (repo *Repo) getAWSTargetGroups()`
- `func (repo *Repo) getAWSTargetHealths(tgarn string)`
- `func (repo *Repo) getAWSEcsClusterDetails(clusterArn string)`
- `func (repo *Repo) getAWSEcsSvcList(clusterArns []*string)`
- `func (repo *Repo) getAWSListClusters()`
- `func (repo *Repo) getAWSEcsClusters(clusters []*string)`
- `func (repo *Repo) getAWSEcsDescribeTaskDefinition(services []*ecs.Service)`
- `func (repo *Repo) getAWSEcsListAndDescribeTasks(clusterName, serviceName string)`
- `func (repo *Repo) getAWSEcsDescribeService(clusterName, serviceName string)`
- `func (repo *Repo) getAWSEcrRepos(repoName string)`
- `func (repo *Repo) getAWSEcrDescribeImages(repoUri, repoName string)`

```go
// 세션 리포지토리
type Repo struct {
	sess *session.Session
}

// 세션 리포지토리 인터페이스
func RepoInterface(param *session.Session) *Repo {
	return &Repo{sess: param}
}

/**
 * 세션 리포지토리 Repo 구현체
 */

// TG 조회 -> tgMap 저장
func (repo *Repo) getAWSTargetGroups() {

	// 데이터 초기화
	tgMap = make(map[string]*elbv2.TargetGroup)

	// ELBV2 서비스 생성
	svc := elbv2.New(repo.sess)
	input := &elbv2.DescribeTargetGroupsInput{ // 요청 파라미터
		Names: []*string {
			// aws.String("awsdc-tg-erp-dev-tdms-7080"),
		},
	}
	// ELBV2 서비스 api DescribeTargetGroups 호출
	pageNum := 0
	// result, err := svc.DescribeTargetGroups(input)
	err := svc.DescribeTargetGroupsPages(input, func(page *elbv2.DescribeTargetGroupsOutput, lastPage bool) bool {
			pageNum++
			log.Println("PAGE result data size: ", len(page.TargetGroups))
			for _, tg := range page.TargetGroups {
				tgMap[*tg.TargetGroupName] = tg
			}
			return !lastPage
	})
	if err != nil {
		handleError(err)
		return
	}
}
```


❚ 컨테이너화


```sh
```

