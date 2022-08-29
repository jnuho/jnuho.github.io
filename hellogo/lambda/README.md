


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
	FirstName string `json:"firstname`
	LastName string `json:"lastname`
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
  - upload image to S3 serverless

```go

```