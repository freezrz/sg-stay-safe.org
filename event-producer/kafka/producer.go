package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
)

func InitProducer(kafkaConn string) (sarama.SyncProducer, error) {
	// setup sarama log to stdout
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	// producer config
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.NoResponse
	config.Producer.Return.Successes = true

	// async producer
	//prd, err := sarama.NewAsyncProducer([]string{kafkaConn}, config)

	// sync producer
	prd, err := sarama.NewSyncProducer([]string{kafkaConn}, config)
	log.Printf("Connecting to kafka server: %s", kafkaConn)
	return prd, err
}

// SendWithKey with partition key
func Send(producer sarama.SyncProducer, topic string, data string) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		// Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(data),
	}
	log.Printf("Sending msg to kafka Topic: %s", topic)
	_, _, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("Error publish: ", err.Error())
	}
	//fmt.Println("Partition: ", p)
	//fmt.Println("Offset: ", o)
	fmt.Println("Published msg to kafka for consuming")
}
