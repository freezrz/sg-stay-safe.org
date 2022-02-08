package kafka

import (
	"fmt"
	"log"
	"time"

	"github.com/Shopify/sarama"
	"github.com/wvanbergen/kafka/consumergroup"
)

type Consumer struct {
	ZookeeperConn string
	Topic         string
	ConsumerGroup *consumergroup.ConsumerGroup
}

func New() *Consumer {
	return &Consumer{}
}

func (k *Consumer) WithZookeeperConn(s string) *Consumer {
	k.ZookeeperConn = s
	return k
}

func (k *Consumer) WithTopic(s string) *Consumer {
	k.Topic = s
	return k
}

func (k *Consumer) InitConsumer() {
	if k.Topic == "" {
		k.Topic = "checkin-msk-topic"
	}
	if k.ZookeeperConn == "" {
		k.ZookeeperConn = "z-3.checkin-msk-clust.0srrl5.c2.kafka.ap-southeast-1.amazonaws.com:2181"
	}
	cgroup := "zgroup"
	// consumer config
	config := consumergroup.NewConfig()
	config.Offsets.Initial = sarama.OffsetNewest
	config.Offsets.ProcessingTimeout = 3 * time.Second

	// join to consumer group
	cg, err := consumergroup.JoinConsumerGroup(cgroup, []string{k.Topic}, []string{k.ZookeeperConn}, config)
	if err != nil {
		log.Panic(err.Error())
	}
	k.ConsumerGroup = cg
}

func (k *Consumer) Consume() {
	for {
		select {
		case msg := <-k.ConsumerGroup.Messages():
			// messages coming through chanel
			// only take messages from subscribed Topic
			if msg == nil || msg.Topic != k.Topic {
				continue
			}
			fmt.Println("Topic: ", msg.Topic)
			fmt.Println("Value: ", string(msg.Value))
			// commit to zookeeper that message is read
			// this prevent read message multiple times after restart
			err := k.ConsumerGroup.CommitUpto(msg)
			if err != nil {
				fmt.Println("Error commit zookeeper: ", err.Error())
			}
		}
	}
}
