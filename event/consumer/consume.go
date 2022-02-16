package main

import (
	"sg-stay-safe.org/config"
	"sg-stay-safe.org/pkg/kafka"
)

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
