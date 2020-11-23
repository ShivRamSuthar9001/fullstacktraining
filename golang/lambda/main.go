package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type event struct {
	UserName string
}

func handler(e event) (string, error) {
	return fmt.Sprintf("<h1>Hello %v from lambda Go</h1>", e.UserName), nil
}

func main() {
	lambda.Start(handler)
}
