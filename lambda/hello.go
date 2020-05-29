package main

import (
  "encoding/json"
  "fmt"
  "context"
  "github.com/aws/aws-lambda-go/lambda"
  "github.com/pusher/pusher-http-go"
)

type HttpEvent struct {
  Body string `json:"body"`
}

type RtcConfirm struct {
  Id string `json:"id"`
  Room string `json:"room"`
  Answer string `json:"answer"`
}

func HandleRequest(ctx context.Context, event HttpEvent) (string, error) {
  pusherClient := pusher.Client{
    AppID: "1009454",
    Key: "151c09f47c37e803cb69",
    Secret: "e2364b03a43a70ed918e",
    Cluster: "ap3",
    Secure: true,
  }

  response := RtcConfirm{}
  if err := json.Unmarshal([]byte(event.Body), &response); err != nil {
    return fmt.Sprintf("ERROR! %s", event.Body), nil
  }

  responseData := map[string]string{"answer": response.Answer, "id": response.Id}
  pusherClient.Trigger(response.Room, "join", responseData)

  return fmt.Sprintf("DONE %s!", event.Body), nil
}

func main() {
  lambda.Start(HandleRequest)
}
