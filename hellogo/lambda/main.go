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