package main

import (
  "fmt"
  "github.com/aws/aws-sdk-go/service/elbv2"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/awserr"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/aws/credentials"
)

func main() {

  sess, err := session.NewSession(&aws.Config{
    Region: aws.String("ap-northeast-2"),
    Credentials: credentials.NewSharedCredentials("", "APPDEV"),
  })

  svc := elbv2.New(sess)
  input := &elbv2.DescribeTargetGroupsInput{
      TargetGroupArns: []*string{
          aws.String("arn:aws:elasticloadbalancing:ap-northeast-2:734976340835:targetgroup/AWSDC-TG-COM-PRD-XENS-7080/6ffa6663e34ecc06"),

      },
  }

  result, err := svc.DescribeTargetGroups(input)
  if err != nil {
      if aerr, ok := err.(awserr.Error); ok {
          switch aerr.Code() {
          case elbv2.ErrCodeLoadBalancerNotFoundException:
              fmt.Println(elbv2.ErrCodeLoadBalancerNotFoundException, aerr.Error())
          case elbv2.ErrCodeTargetGroupNotFoundException:
              fmt.Println(elbv2.ErrCodeTargetGroupNotFoundException, aerr.Error())
          default:
              fmt.Println(aerr.Error())
          }
      } else {
          // Print the error, cast err to awserr.Error to get the Code and
          // Message from an error.
          fmt.Println(err.Error())
      }
      return
  }

  fmt.Println(result)
}
