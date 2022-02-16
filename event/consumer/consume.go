package main

import (
	"sg-stay-safe.org/event/config"
	"sg-stay-safe.org/event/pkg"
)

// ProcessEvent function Using AWS Lambda computed event
func ProcessEvent() {
	consumer := kafka.New()
	consumer.Topic = config.CheckinEventKafkaTopic
	consumer.ZookeeperConn = config.CheckinEventKafkaZooKeeper

	consumer.InitConsumer()
	consumer.Consume()
}

func main() {
	ProcessEvent()
}
