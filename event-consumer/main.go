package main

import (
	"sg-stay-safe.org/event-consumer/kafka"
)

// ProcessEvent function Using AWS Lambda computed event
func ProcessEvent() {
	consumer := kafka.New()
	consumer.InitConsumer()
	consumer.Consume()
}

func main() {
	ProcessEvent()
}
