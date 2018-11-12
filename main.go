package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/plivo/ant-service/work"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// set these
const (
	qURL = ""
	callURL = ""
)

func main() {
	s, err := session.NewSession(&aws.Config{
		Region: aws.String(endpoints.UsEast1RegionID),
	})
	if err != nil {
		log.Fatalf("session.NewSession() failed: %s", err.Error())
	}

	svc := sqs.New(s)

	payload := &work.MsgPayload{
		URI:         callURL,
		Method:      "GET",
		SuccessCode: "200",
	}

	msg := &work.Message{
		Payload:       payload,
		RetryCount:    2,
		ID:            "asdfgjkl",
		SourceService: "source",
		MaxRetry:      5,
	}

	b, err := json.Marshal(msg)
	if err != nil {
		log.Fatal("json.Matshal() failed: %s", err.Error())
	}

	result, err := svc.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(string(b)),
		QueueUrl:    aws.String(qURL)})
	if err != nil {
		log.Fatalf("svc.SendMessage() failed: %s", err.Error())
	}

	fmt.Printf("%+v\n", *result)
}
