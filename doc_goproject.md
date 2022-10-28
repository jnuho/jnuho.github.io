❚ Go Application Using AWS Go Sdk
- [repository](https://github.com/jnuho/goproject)
- [aws sdk for go](https://aws.amazon.com/sdk-for-go/)

<div>
<img src="assets/images/Animation.gif" height="40%" width="40%" alt="go sdk app">
</div>

❚ Skills
```
Backend : Go
Frontend : Javascript, HTML, CSS
```

❚ Features
```
GET operation on EC2
	Target Group
	Target's instance id (or container ip), port, target health
	Load Balancer
GET operation on ECS
	Cluster
	Service
	Task Definition
	Container
GET operation on ECR
	Repository
	Image
	Lifecycle Policy
GET operation on Tags


Typing 'Enter Key' or clicking 'Search Button'
	appends a list of resources on the web page below the text input.

Onclick event on each list element would trigger
	ajax GET API call for detailed info of the clicked resource element.
	The list element will unfold or show a modal windows for detailed info.

In order to sort a list of resources (ECR image uris),
	"sort" library of golang was used with overridden methods 'Len', 'Less' and 'Swap'
```

<br>

❚ Aws Profile
```
# IAM Role with permission policies :
#		Creatre role
#			Add Permissions

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


```go
import "github.com/aws/aws-sdk-go/aws/session"

// AWS profile names (~/.aws/config)
type AwsProfile struct{
	dev string
	stg string
	prd string
}
var awsProfile AwsProfile

// Init session object
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

❚ Define struct Repo<br>
❚ Define interface that returns Repo<br>
❚ Define methods that implement Repo struct<br>

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
// Session repository
type Repo struct {
  sess *session.Session
}

// Session repository interface
func RepoInterface(param *session.Session) *Repo {
  return &Repo{sess: param}
}

/**
 * Session repository Repo implementation
 */

// Get TG -> store to tgMap
func (repo *Repo) getAWSTargetGroups() {

  // Initialize object for storing data
  tgMap = make(map[string]*elbv2.TargetGroup)

  // Create elbv2 Service
  svc := elbv2.New(repo.sess)
  input := &elbv2.DescribeTargetGroupsInput{ // request parameter
    Names: []*string {
      // aws.String("awsdc-tg-erp-dev-tdms-7080"),
    },
  }
  // ELBV2 service calls api DescribeTargetGroups
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



