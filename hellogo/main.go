package main

import (
  "fmt"
	"context"
	"log"
  // "flag"
  "os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
  // Load specific AWS Configuration
  // SDK checks the AWS_PROFILE environment variable to determine which profile to use
  // By default aws choose default profile
  os.Setenv("AWS_PROFILE", "APPPRD")

  //sess, err := session.NewSession(&aws.Config{
  //  Region:      aws.String("ap-northeast-2"),
  //  Credentials: credentials.NewSharedCredentials("", "APPPRD")
  //})

	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
    log.Println("111")
    fmt.Println("ABC junho")
	}

	// Create an Amazon S3 service client
	client := s3.NewFromConfig(cfg)

	// Get the first page of results for ListObjectsV2 for a bucket
	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		//Bucket: aws.String("my-bucket"),
		Bucket: aws.String("awsdc-s3-com-prd-s3accesslog"),
	})
	if err != nil {
		log.Fatal(err)
    log.Println("222")
	}

	log.Println("first page results:")
	for _, object := range output.Contents {
		log.Printf("key=%s size=%d", aws.ToString(object.Key), object.Size)
	}
}

