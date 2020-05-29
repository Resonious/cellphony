package main

import (
  "fmt"
  "context"
  "github.com/aws/aws-lambda-go/lambda"
)

type RtcEvent struct {
  Session string `json:"session"`
}

func HandleRequest(ctx context.Context, event RtcEvent) (string, error) {
  return fmt.Sprintf("Hello %s!", event.Session ), nil
}

func main() {
  lambda.Start(HandleRequest)
}
