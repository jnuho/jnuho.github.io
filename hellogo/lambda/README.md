


```sh
# Linux 환경
goos=linux go build -o main main.go && zip archive.zip main
# Git bash
# main (executable) git에 포함시킨 후 .zip 압축 생성
goos=linux go build -o main main.go && git archive -o archive.zip HEAD main

# Goto AWS lamda console
#  > "Create a function"
#  Handler: 'main' (executable)
#  Memory : 128MB (that function allocates)
#  Timeout : (time functon call lasts)
# Function Code > Action > Upload a zip file
```

```go
package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type InputEvent struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}
func Handler(event InputEvent) (string, error) {
	// fmt.Println("Function invoked!")
	fmt.Println(event.FirstName)
	fmt.Println(event.LastName)
	return "It worked!", nil
}

func main() {
	lambda.Start(Handler)
}
```
- Configure Test event

```json
{
  "firstname": "ABC",
  "lastname": "LAAA"
}
```

- Example
	- Download image from link
  - Upload image to S3 using lambda (serverless)


```go

```

```json
{
	"link": "https://cdn.shopify.com/s/files/1/2570/4610/products/BeefSample_1024x1024@2x.jpg?v=1650451645",
	"key": "BeefSample.jpg"
}
```

- Access Denied
	- S3 Permission required
	- Attach Policy to Role/Profile
		- "S3FullAccess"

- Trigger
	- e.g. API Gateway
	- API endpoint -> AWS Lambda