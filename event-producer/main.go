package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"sg-stay-safe.org/event-producer/kafka"
)

// ProcessEvent function Using AWS Lambda computed event
func ProcessEvent(event events.KafkaEvent) {

	//jsonPayload := Payload{}
	log.Println("producing event... 1")
	producer, err := kafka.InitProducer("b-1.checkin-msk-clust.0srrl5.c2.kafka.ap-southeast-1.amazonaws.com:9092")
	if err != nil {
		log.Println(err.Error())
	}
	kafka.Send(producer, "checkin-msk-topic", `{"hello":"world"}`)
}

func main() {
	lambda.Start(ProcessEvent)
}
