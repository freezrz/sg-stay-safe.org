package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	lambdaService "github.com/aws/aws-sdk-go/service/lambda"
	"log"
	"os"
	"sg-stay-safe.org/config"
	"sg-stay-safe.org/pkg/kafka"
	lambdaPkg "sg-stay-safe.org/pkg/lambda"
	"sg-stay-safe.org/protocol"
)

func ProcessEvent() {
	Consume(kafka.New(os.Getenv("Topic"), os.Getenv("ZooKeeper")))
}

func main() {
	ProcessEvent()
}

func Consume(k *kafka.Consumer) {
	for {
		select {
		case msg := <-k.ConsumerGroup.Messages():
			// messages coming through chanel
			// only take messages from subscribed Topic
			if msg == nil || msg.Topic != k.Topic {
				continue
			}
			log.Println("topic: ", msg.Topic)
			log.Println("value: ", string(msg.Value))
			// commit to zookeeper that message is read
			// this prevent read message multiple times after restart
			var event protocol.ViolationEvent
			err := json.Unmarshal(msg.Value, &event)
			if err != nil {
				log.Println(err.Error())
			}
			log.Println(fmt.Sprintf("debug trying to send email for site: %s", event.CheckInEvent.SiteId))
			SendEmail(event)
			err = k.ConsumerGroup.CommitUpto(msg)
			if err != nil {
				log.Println("error commit zookeeper: ", err.Error())
			}
		}
	}
}

func SendEmail(event protocol.ViolationEvent) {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	client := lambdaService.New(sess, &aws.Config{Region: aws.String(config.AWSRegion)})

	payload, err := json.Marshal(event)
	if err != nil {
		log.Println(err.Error())
		return
	}

	var cmd string
	cmd = config.RetrieveRegionEmailBySiteIdLambda
	_, bs, err := lambdaPkg.CallLambdaFunc(client, cmd, payload)
	if err != nil {
		log.Println(err.Error())
		return
	}
	email := string(bs)
	log.Println(email)
}
