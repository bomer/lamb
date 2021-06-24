package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	return fmt.Sprintf("Hello %s!, your sha hash is %s", name.Name, CreateKey(name.Name)), nil
}

func main() {
	lambda.Start(HandleRequest)
}
