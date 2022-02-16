package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"sg-stay-safe.org/event/config"
	"sg-stay-safe.org/event/pkg"
	"sg-stay-safe.org/event/protocol"
)

// ProcessEvent function Using AWS Lambda computed event
func ProcessEvent(ctx context.Context, event protocol.CheckInEvent) (protocol.GeneralResponse, error) {
	log.Println("producing event...")
	producer, err := kafka.InitProducer(config.CheckinEventKafkaBootstrap)
	if err != nil {
		log.Println(err.Error())
	}
	bs, err := json.Marshal(event)
	if err != nil {
		// TODO: notify monitoring system
		log.Println(err.Error())
		return protocol.GeneralResponse{Code: config.CodeMarshalError, Msg: err.Error()}, err
	}
	if err := kafka.Send(producer, config.CheckinEventKafkaTopic, string(bs)); err != nil {
		return protocol.GeneralResponse{Code: config.CodeProduceEventError, Msg: err.Error()}, err
	}
	return protocol.GeneralResponse{Code: config.CodeOK, Msg: "produce event ok..."}, nil
}

func main() {
	lambda.Start(ProcessEvent)
}
